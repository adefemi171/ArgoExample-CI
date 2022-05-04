package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	helloWorld(response, request)
	got := response.Body.String()
	want := "Test CI Webserver 3"

	if got != want {
		t.Error("Wanted:", want, "Got:", got)
	}

}
