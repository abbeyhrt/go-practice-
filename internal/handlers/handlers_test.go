package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abbeyhrt/keep-up-graphql/internal/handlers"
)

func TestSlashRoute(t *testing.T) {
	handler := handlers.New()
	srv := httptest.NewServer(handler)

	defer srv.Close()

	res, err := http.Get(srv.URL)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("Received non 200 response: %d\n", res.StatusCode)
	}

	expected := "hello, world"

	actual, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	if expected != string(actual) {
		t.Errorf("Expected the message '%s' got '%s'\n", expected, actual)
	}

}
