package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name") //name="NAME"
	w.WriteHeader(200)
	w.Write([]byte(name))

}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
