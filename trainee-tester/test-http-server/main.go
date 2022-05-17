package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		w,
		`<doctype html>
		<html>
		<head>
		<title>Hello</title>
		</head>
		<body>
		<H1>
		Hello!!!
		</H1>
		</body>
		</html>`,
	)

}

func pingHandler(w http.ResponseWriter, r *http.Request) {

	sendJSON(w, map[string]string{"answer": "pong"})

	//fmt.Fprintf(w, "PING")

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, map[string]string{"answer": "NAME"})

}

func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	// start time duration
	t0 := time.Now()
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
	// finish time duration
	t1 := time.Now()
	// Get duration.
	d := t1.Sub(t0)
	fmt.Println("Duration", d)
}

// Added func sendJSON
func sendJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		log.Println("ERROR:", err)
	}
	return err
}

func main() {
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/timeout", timeoutHandler)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
