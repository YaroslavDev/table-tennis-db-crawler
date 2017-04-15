package rubber

import (
	"html/template"
	"net/http"
	"log"
)

type rubberController struct {
	template *template.Template
	fetcher  RubberFetcher
}

func NewRubberController(template *template.Template) *rubberController {
	return &rubberController{template: template, fetcher: NewTTDBRubberFetcher()}
}

func (a rubberController) RubbersPage(w http.ResponseWriter, req *http.Request) {
	rubbers, err := a.fetcher.FetchRubbers()
	if err != nil {
		log.Fatalln(err)
	}
	err = a.template.ExecuteTemplate(w, "rubbers.html", rubbers)
	if err != nil {
		log.Fatalln(err)
	}
}

func (a rubberController) SynchronizeRubbers(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/rubber", http.StatusSeeOther)
}