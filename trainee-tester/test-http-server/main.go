package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(200 * time.Millisecond)
	type User struct {
		Answer string `json:"answer"`
	}

	p1 := User{
		Answer: "NAME_timeout",
	}

	// Set response header
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&p1)
	if err != nil {
		//... handle error
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/timeout", timeoutHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
