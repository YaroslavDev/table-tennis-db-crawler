package web

import (
	"html/template"
	"net/http"
	"log"
	"path/filepath"
)

type AppController struct {
	template *template.Template
}

func NewAppController(templateBasePath string) *AppController {
	templatesPath := filepath.Join(templateBasePath, "*.html")
	tpl := template.Must(template.ParseGlob(templatesPath))
	return &AppController{template: tpl}
}

func (a AppController) IndexPage(w http.ResponseWriter, req *http.Request) {
	err := a.template.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}