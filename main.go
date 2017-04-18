package main

import (
	"net/http"
	"log"
	"github.com/YaroslavDev/table-tennis-db-crawler/rubber"
	"github.com/YaroslavDev/table-tennis-db-crawler/config"
)

func routes() {
	rubberController := rubber.NewRubberController(config.TPL, &config.ConnectionUrl)
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.HandleFunc("/", landingPage)
	http.HandleFunc("/rubber", rubberController.RubbersPage)
	http.HandleFunc("/rubber-sync", rubberController.SynchronizeRubbers)
	http.HandleFunc("/blades", bladesPage)
	http.HandleFunc("/blades-sync", bladesSync)
}

func main() {
	routes()
	log.Println("Table tennis DB crawler started...")
	http.ListenAndServe(":8080", nil)
}

func landingPage(w http.ResponseWriter, _ *http.Request) {
	err := config.TPL.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func bladesPage(w http.ResponseWriter, _ *http.Request) {
	err := config.TPL.ExecuteTemplate(w, "blades.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func bladesSync(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/blades", http.StatusSeeOther)
}