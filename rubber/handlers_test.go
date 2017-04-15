package rubber

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"html/template"
	"github.com/antchfx/xquery/html"
)

func TestRubberController_RubbersPage(t *testing.T) {
	request, err := http.NewRequest("GET", "/rubber", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	rubberController := NewRubberController(template.Must(template.ParseGlob("../templates/*.html")), nil)
	rubberController.finder = FinderStub{}
	rubberController.rubberRepository = RubberRepositoryStub{}

	rubberController.RubbersPage(responseWriter, request)

	response := responseWriter.Result()
	expectedStatusCode := http.StatusOK
	if response.StatusCode != expectedStatusCode {
		t.Fatalf("Received %d but expected %d\n", response.StatusCode, expectedStatusCode)
	}
	root, err := htmlquery.Parse(responseWriter.Body)
	if err != nil {
		t.Fatal(err)
	}
	trNodes := htmlquery.Find(root, "//tr")
	if len(trNodes) != 3 {
		t.Fatalf("Rubbers page should have 3 tr elements(1 for header and 2 for rubbers), but has %d\n", len(trNodes))
	}
	expectedRubber1Name := "Donic Acuda S2"
	actualRubber1Name := htmlquery.InnerText(htmlquery.FindOne(trNodes[1], "//td"))
	if actualRubber1Name != expectedRubber1Name {
		t.Fatalf("First rubber should be %s, but was %s\n", expectedRubber1Name, actualRubber1Name)
	}
	expectedRubber2Name := "Butterfly Tenergy 05"
	actualRubber2Name := htmlquery.InnerText(htmlquery.FindOne(trNodes[2], "//td"))
	if actualRubber2Name != expectedRubber2Name {
		t.Fatalf("Second rubber should be %s, but was %s\n", expectedRubber2Name, actualRubber2Name)
	}
}

func TestRubberController_SynchronizeRubbers(t *testing.T) {
	request, err := http.NewRequest("GET", "/rubber-sync", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	rubberController := NewRubberController(template.Must(template.ParseGlob("../templates/*.html")), nil)
	rubberController.finder = FinderStub{}
	rubberController.rubberRepository = RubberRepositoryStub{}

	rubberController.SynchronizeRubbers(responseWriter, request)

	response := responseWriter.Result()
	expectedStatusCode := http.StatusSeeOther
	if response.StatusCode != expectedStatusCode {
		t.Fatalf("Received %d but expected %d\n", response.StatusCode, expectedStatusCode)
	}
	if response.Header.Get("Location") != "/rubber" {
		t.Fatal("Location header should be /rubber")
	}
}

type FinderStub struct {}

func (service FinderStub) FindRubbers() ([]*Rubber, error) {
	rubbers := []*Rubber{
		{Name: "Donic Acuda S2", Speed: 8.8},
		{Name: "Butterfly Tenergy 05", Speed: 9.5},
	}
	return rubbers, nil
}

type RubberRepositoryStub struct {}

func (r RubberRepositoryStub) GetRubbers() ([]Rubber, error) {
	rubbers := []Rubber{
		{Name: "Donic Acuda S2", Speed: 8.8},
		{Name: "Butterfly Tenergy 05", Speed: 9.5},
	}
	return rubbers, nil
}

func (r RubberRepositoryStub) SaveRubber(rubber *Rubber) error {
	return nil
}