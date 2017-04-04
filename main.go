package main

import (
	"net/http"
	"html/template"
	"log"
)

type App struct {}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	app := App{}
	http.HandleFunc("/", app.IndexPage)
	log.Println("Table tennis DB crawler started...")
	http.ListenAndServe(":8080", nil)
}

func (a App) IndexPage(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
