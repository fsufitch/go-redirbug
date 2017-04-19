package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/foo", http.RedirectHandler("/foo/", 301))
	http.HandleFunc("/foo/", func(w http.ResponseWriter, request *http.Request) {
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
	http.ListenAndServe(":8081", nil)
}
