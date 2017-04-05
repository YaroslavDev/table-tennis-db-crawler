package services

import (
	"github.com/YaroslavDev/table-tennis-db-crawler/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type TTDBRubberFetchingService struct {}

func (service TTDBRubberFetchingService) FetchRubbers() ([]model.Rubber, error) {
	doc, err := goquery.NewDocument("http://www.tabletennisdb.com/rubber")
	if err != nil {
		return nil, err
	}

	doc.Find("div #brand-andro").Each(func(i int, s *goquery.Selection) {
		s.Find("tr:not([class])").Each(func(i int, trSelection *goquery.Selection) {
			fmt.Println(trSelection.Text() + "JORA")
		})
	})
	return nil, nil
}
