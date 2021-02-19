package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.Write([]byte("GET Hello World!"))
	case "POST":
		w.Write([]byte("POST Hello World!"))
	default:
		w.Write([]byte("UNKNOWN Hello World!"))
	}
}
