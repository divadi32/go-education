package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
