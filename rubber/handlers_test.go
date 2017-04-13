package rubber

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"html/template"
)

func TestIndexPage(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	rubberController := NewRubberController(template.Must(template.ParseGlob("../templates/*.html")))
	rubberController.fetcher = FetcherStub{}

	rubberController.RubbersPage(responseWriter, request)

	response := responseWriter.Result()
	expectedStatusCode := 200
	if response.StatusCode != expectedStatusCode {
		t.Fatalf("Received %d but expected %d\n", response.StatusCode, expectedStatusCode)
	}
}

type FetcherStub struct {}

func (service FetcherStub) FetchRubbers() ([]*Rubber, error) {
	rubbers := []*Rubber{
		{Name: "Donic Acuda S2", Speed: 8.8},
		{Name: "Butterfly Tenergy 05", Speed: 9.5},
	}
	return rubbers, nil
}
