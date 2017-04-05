package web

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
	app := NewAppController("../templates")

	app.IndexPage(responseWriter, request)

	response := responseWriter.Result()
	expectedStatusCode := 200
	if response.StatusCode != expectedStatusCode {
		t.Fatalf("Received %d but expected %d\n", response.StatusCode, expectedStatusCode)
	}
}
