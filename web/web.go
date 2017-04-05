package web

import (
	"html/template"
	"net/http"
	"log"
	"path/filepath"
	"github.com/YaroslavDev/table-tennis-db-crawler/services"
)

type AppController struct {
	template              *template.Template
	RubberFetchingService services.RubberFetchingService
}

func NewAppController(templateBasePath string) *AppController {
	templatesPath := filepath.Join(templateBasePath, "*.html")
	tpl := template.Must(template.ParseGlob(templatesPath))
	return &AppController{template: tpl}
}

func (a AppController) IndexPage(w http.ResponseWriter, req *http.Request) {
	rubbers, err := a.rubberFetchingService().FetchRubbers()
	if err != nil {
		log.Fatalln(err)
	}
	err = a.template.ExecuteTemplate(w, "index.html", rubbers)
	if err != nil {
		log.Fatalln(err)
	}
}

func (a AppController) rubberFetchingService() services.RubberFetchingService {
	if a.RubberFetchingService == nil {
		a.RubberFetchingService = services.TTDBRubberFetchingService{}
	}
	return a.RubberFetchingService
}