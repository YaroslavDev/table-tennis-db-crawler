package services

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YaroslavDev/table-tennis-db-crawler/model"
	"strconv"
	"log"
	"sync"
	"strings"
)

type TTDBRubberFetchingService struct{}

func (service TTDBRubberFetchingService) FetchRubbers() ([]*model.Rubber, error) {
	doc, err := goquery.NewDocument("http://www.tabletennisdb.com/rubber")
	if err != nil {
		return nil, err
	}

	rubberUrls := make([]string, 0, 2000)
	doc.Find("div [id^=brand-]").Each(func(i int, s *goquery.Selection) {
		s.Find("tr:not([class])").Each(func(i int, trSelection *goquery.Selection) {
			rubberUrl, exists := trSelection.Find("td").First().Find("a").First().Attr("href")
			if exists {
				rubberUrls = append(rubberUrls, rubberUrl)
			}
		})
	})

	// TODO: Use 30 goroutines to fetch all rubbers
	rubberUrls = rubberUrls[:30]
	numRubbersFound := len(rubberUrls)
	var waitGroup sync.WaitGroup
	waitGroup.Add(numRubbersFound)
	rubberChannel := make(chan *model.Rubber)

	for _, rubberUrl := range rubberUrls {
		go func(url string) {
			fetchRubber(url, rubberChannel)
			waitGroup.Done()
		}(rubberUrl)
	}

	rubbers := make([]*model.Rubber, 0, 2000)

	go func() {
		waitGroup.Wait()
		close(rubberChannel)
	}()

	for rubber := range rubberChannel {
		rubbers = append(rubbers, rubber)
	}

	return rubbers, nil
}

func fetchRubber(url string, rubberChannel chan *model.Rubber) {
	doc, err := goquery.NewDocument("http://www.tabletennisdb.com/" + url)
	if err != nil {
		log.Println(err)
	}

	rubber := model.Rubber{}
	rubberName := doc.Find("h1.ul.fn").First().Text()
	rubber.Name = rubberName
	doc.Find("table.ProductRatingTable.ratingtable").First().
		Find("tr").Each(func(parameterIndex int, parameterSelection *goquery.Selection) {
			parameterSelection.Find("td").Each(func(columnIndex int, tdSelection *goquery.Selection) {
				if columnIndex == 1 {
					tdText := tdSelection.Text()
					tdText = strings.TrimSpace(tdText)
					tdText = tdText[:3]
					parameterValue64, _ := strconv.ParseFloat(tdText, 32)
					parameterValue := float32(parameterValue64)
					switch parameterIndex {
					case 0:
						rubber.Speed = parameterValue
					case 1:
						rubber.Spin = parameterValue
					case 2:
						rubber.Control = parameterValue
					case 3:
						rubber.Tackiness = parameterValue
					case 4:
						rubber.Weight = parameterValue
					case 5:
						rubber.SpongeHardness = parameterValue
					case 6:
						rubber.Gears = parameterValue
					case 7:
						rubber.ThrowAngle = parameterValue
					case 8:
						rubber.Consistency = parameterValue
					case 9:
						rubber.Durability = parameterValue
					}
				}
			})
		})
	rubberChannel <- &rubber
}
