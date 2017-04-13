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

func (a RubberController) RubbersPage(w http.ResponseWriter, req *http.Request) {
	rubbers, err := a.fetcher.FetchRubbers()
	if err != nil {
		log.Fatalln(err)
	}
	err = a.template.ExecuteTemplate(w, "rubbers.html", rubbers)
	if err != nil {
		log.Fatalln(err)
	}
}

func (a RubberController) SynchronizeRubbers(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/rubber", http.StatusSeeOther)
}