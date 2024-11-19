package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Testing Golang APP! v20\n")
	}

	http.HandleFunc("/", helloHandler)
    log.Println("Listing for v20 requests at http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
