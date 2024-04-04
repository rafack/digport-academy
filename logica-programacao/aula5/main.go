package main

import (
	"io"
	"net/http"
)

func main() {
	handleHelloWorld := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello, world")
	}

	http.HandleFunc("/", handleHelloWorld)

	http.ListenAndServe(":8080", nil)
}
