package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func query(url string) (body string, ok bool) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("Authorization", "correct-password")
	response, err := (&http.Client{}).Do(request)
	if err != nil {
		panic(err)
	}

	bodyChars, _ := ioutil.ReadAll(response.Body)
	body = string(bodyChars)
	ok = response.StatusCode >= 200 && response.StatusCode < 300
	return
}

func main() {
	body, ok := query("http://localhost:8081/foo")
	if ok {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failure!")
		fmt.Println(body)
	}
}
