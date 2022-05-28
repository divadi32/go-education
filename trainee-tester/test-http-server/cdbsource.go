package main

import "net/http"

func cdbSourceHeandler(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, map[string]string{"answer": "pong"})
}
