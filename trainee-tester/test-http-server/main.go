package main

import (
	"encoding/json"
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

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Answer string `json:"answer"`
	}

	p := User{
		Answer: "NAME",
	}

	// Set response header
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&p)
	if err != nil {
		//... handle error
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/json", jsonHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
