package rubber

import (
	"html/template"
	"net/http"
	"log"
	"path/filepath"
)

type RubberController struct {
	template              *template.Template
	RubberFetchingService RubberFetchingService
}

func NewAppController(templateBasePath string) *RubberController {
	templatesPath := filepath.Join(templateBasePath, "*.html")
	tpl := template.Must(template.ParseGlob(templatesPath))
	return &RubberController{template: tpl}
}

func (a RubberController) IndexPage(w http.ResponseWriter, req *http.Request) {
	rubbers, err := a.rubberFetchingService().FetchRubbers()
	if err != nil {
		log.Fatalln(err)
	}
	err = a.template.ExecuteTemplate(w, "index.html", rubbers)
	if err != nil {
		log.Fatalln(err)
	}
}

func (a RubberController) rubberFetchingService() RubberFetchingService {
	if a.RubberFetchingService == nil {
		a.RubberFetchingService = TTDBRubberFetchingService{}
	}
	return a.RubberFetchingService
}