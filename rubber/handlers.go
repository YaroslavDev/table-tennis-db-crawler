package rubber

import (
	"html/template"
	"net/http"
	"log"
)

type RubberController struct {
	template *template.Template
	fetcher  RubberFetcher
}

func NewRubberController(template *template.Template) *RubberController {
	return &RubberController{template: template, fetcher: NewTTDBRubberFetcher()}
}

func (a RubberController) IndexPage(w http.ResponseWriter, req *http.Request) {
	rubbers, err := a.fetcher.FetchRubbers()
	if err != nil {
		log.Fatalln(err)
	}
	err = a.template.ExecuteTemplate(w, "index.html", rubbers)
	if err != nil {
		log.Fatalln(err)
	}
}