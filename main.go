package main

import (
	"net/http"
	"log"
	"github.com/table-tennis-db-crawler/web"
)

func main() {
	app := web.NewAppController("templates")
	http.HandleFunc("/", app.IndexPage)
	log.Println("Table tennis DB crawler started...")
	http.ListenAndServe(":8080", nil)
}