package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	mux := http.NewServeMux()
	mux.Handle("/foo", http.RedirectHandler("/foo/", 301))
	mux.HandleFunc("/foo/", func(w http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")
		if auth == "correct-password" {
			w.WriteHeader(200)
			w.Write([]byte("OK"))
		} else {
			w.WriteHeader(401)
			message := fmt.Sprintf("Unauthorized `%s`", auth)
			w.Write([]byte(message))
		}
	})
	testServer = httptest.NewServer(mux)
	defer testServer.Close()
	status := m.Run()
	os.Exit(status)
}

func TestQueryNoSlash(t *testing.T) {
	_, ok := query(testServer.URL + "/foo")
	if !ok {
		t.Error("Failed with no slash")
	}
}

func TestQueryWithSlash(t *testing.T) {
	_, ok := query(testServer.URL + "/foo/")
	if !ok {
		t.Error("Failed with slash")
	}
}
