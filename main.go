package main

import (
	"net/http"
	"log"
	"github.com/YaroslavDev/table-tennis-db-crawler/rubber"
	"github.com/YaroslavDev/table-tennis-db-crawler/config"
)

func main() {
	rubberController := rubber.NewRubberController(config.TPL)
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.HandleFunc("/", rubberController.IndexPage)
	log.Println("Table tennis DB crawler started...")
	http.ListenAndServe(":8080", nil)
}