package main

import (
	"net/http"
	"log"
	"github.com/YaroslavDev/table-tennis-db-crawler/rubber"
)

func main() {
	app := rubber.NewAppController("templates")
	http.HandleFunc("/", app.IndexPage)
	http.HandleFunc("/favicon.ico", http.NotFound)
	log.Println("Table tennis DB crawler started...")
	http.ListenAndServe(":8080", nil)
}