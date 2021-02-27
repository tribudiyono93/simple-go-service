package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listring for request at http://127.0.0.1:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
