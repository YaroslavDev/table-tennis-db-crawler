package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestIndexPage(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	app := App{}

	app.IndexPage(responseWriter, request)

	response := responseWriter.Result()
	if response.StatusCode != 200 {
		t.Fatalf("Received %s but expected 200\n", response.StatusCode)
	}
}
