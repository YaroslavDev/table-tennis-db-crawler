package rubber

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
	app.RubberFetchingService = FetchingServiceMock{}

	app.IndexPage(responseWriter, request)

	response := responseWriter.Result()
	expectedStatusCode := 200
	if response.StatusCode != expectedStatusCode {
		t.Fatalf("Received %d but expected %d\n", response.StatusCode, expectedStatusCode)
	}
}

type FetchingServiceMock struct {}

func (service FetchingServiceMock) FetchRubbers() ([]*Rubber, error) {
	rubbers := []*Rubber{
		{Name: "Donic Acuda S2", Speed: 8.8},
		{Name: "Butterfly Tenergy 05", Speed: 9.5},
	}
	return rubbers, nil
}
