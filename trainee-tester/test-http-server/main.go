package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello")

}

func pingHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "PING")

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name") //name="NAME"
	w.WriteHeader(200)
	w.Write([]byte(name))

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
