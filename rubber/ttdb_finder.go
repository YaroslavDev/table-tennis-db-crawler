package rubber

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"log"
	"sync"
	"strings"
)

type ttDBRubberFinder struct {
	newDocument func(url string) (*goquery.Document, error)
}

const MAX_NUM_WORKERS = 30

func NewTTDBRubberFinder() *ttDBRubberFinder {
	return &ttDBRubberFinder{newDocument: goquery.NewDocument}
}

// FindRubbers finds list of available rubbers from http://www.tabletennisdb.com/rubber
// and then spawns several goroutine workers to concurrently find detailed information
// about each rubber from its own page. E.g. http://www.tabletennisdb.com/rubber/andro-rasant.html
func (service ttDBRubberFinder) FindRubbers() ([]*Rubber, error) {
	doc, err := service.newDocument("http://www.tabletennisdb.com/rubber")
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

	// TODO: Put sleep/reduce amount of workers as ttdb returns 503.
	numWorkers := MAX_NUM_WORKERS
	numFoundRubbers := len(rubberUrls)
	log.Printf("Found %d rubbers", numFoundRubbers)
	if numWorkers > numFoundRubbers {
		numWorkers = numFoundRubbers
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(numWorkers)
	rubberChannel := make(chan *Rubber, numFoundRubbers)
	urlChannel := make(chan string, numFoundRubbers)
	var url string
	for worker := 0; worker < numWorkers; worker++ {
		go service.findRubbersWorker(urlChannel, rubberChannel, &waitGroup)
		url, rubberUrls = rubberUrls[0], rubberUrls[1:]
		urlChannel <- url
	}

	rubbers := make([]*Rubber, 0, 2000)
	go func() {
		waitGroup.Wait()
		log.Println("Found all rubbers. Closing rubber channel...")
		close(rubberChannel)
	}()

	var finished bool = false
	for rubber := range rubberChannel {
		rubbers = append(rubbers, rubber)
		if !finished {
			if len(rubberUrls) > 0 {
				url, rubberUrls = rubberUrls[0], rubberUrls[1:]
				urlChannel <- url
			}
			if len(rubberUrls) == 0 {
				finished = true
				log.Println("No more rubbers left. Closing rubber URL channel...")
				close(urlChannel)
			}
		}
	}

	return rubbers, nil
}

// findRubbersWorker finds detailed information about rubber while there are URLs in urlChannel
func (service ttDBRubberFinder) findRubbersWorker(urlChannel <-chan string, rubberChannel chan<- *Rubber, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urlChannel {
		log.Println("Finding rubbers at " + url)
		service.findRubberFromSingleURL(url, rubberChannel)
	}
}

// findRubberFromSingleURL finds detailed information about rubber, parses it and sends Rubber instance to rubberChannel
func (service ttDBRubberFinder) findRubberFromSingleURL(url string, rubberChannel chan<- *Rubber) {
	doc, err := service.newDocument("http://www.tabletennisdb.com/" + url)
	if err != nil {
		log.Fatal(err)
	}

	rubber := Rubber{}
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
