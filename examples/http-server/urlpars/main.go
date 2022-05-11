package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func hello1Handler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return

	}
	fmt.Fprintf(w, "User ID %d...", id)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name") //name="NAME"
	w.WriteHeader(200)
	w.Write([]byte(name))

}

func main() {
	http.HandleFunc("/test", hello1Handler)
	http.HandleFunc("/hello", helloHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
