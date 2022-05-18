package main

import (
	"log"
	"net/http"
)

func main() {
	setEndpoints()

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
