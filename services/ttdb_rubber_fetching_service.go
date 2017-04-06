package services

import (
	"github.com/YaroslavDev/table-tennis-db-crawler/model"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"log"
)

type TTDBRubberFetchingService struct {}

func (service TTDBRubberFetchingService) FetchRubbers() ([]model.Rubber, error) {
	doc, err := goquery.NewDocument("http://www.tabletennisdb.com/rubber")
	if err != nil {
		return nil, err
	}

	var rubbers []model.Rubber
	var numRubbersFound int
	doc.Find("div [id^=brand-]").Each(func(i int, s *goquery.Selection) {
		s.Find("tr:not([class])").Each(func(i int, trSelection *goquery.Selection) {
			rubber := model.Rubber{}
			trSelection.Find("td").Each(func(i int, tdSelection *goquery.Selection) {
				numRubbersFound++
				tdText := tdSelection.Text()
				tdText = strings.TrimSpace(tdText)
				var parameter float32
				if i > 0 {
					parameter64, err := strconv.ParseFloat(tdText, 32)
					if err != nil {
						return
					}
					parameter = float32(parameter64)
				}
				switch i {
				case 0:
					rubber.Name = tdText
				case 1:
					rubber.Speed = parameter
				case 2:
					rubber.Spin = parameter
				case 3:
					rubber.Tackiness = parameter
				}
			})
			rubbers = append(rubbers, rubber)
		})
	})
	log.Printf("Found in total %d rubbers!", numRubbersFound)
	return rubbers, nil
}
