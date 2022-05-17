package main

import (
	"encoding/json"
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

	time.Sleep(200 * time.Millisecond)
	sendJSON(w, map[string]string{"answer": "pong"})

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
