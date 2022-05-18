package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

// setEndoints (setupRoutes)
func setEndpoints() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/timeout", timeoutHandler)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		w,
		`<!doctype html>
		<html>
		<head>
		<title>Hello</title>
		</head>
		<body>
		<h1>
		Hello!!!
		</h1>
		</body>
		</html>`,
	)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, map[string]string{"answer": "pong"})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		sendJSON(w, map[string]string{"answer": "Param name was not found"})
	}
	sendJSON(w, map[string]string{"answer": name})
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
