package main

import "net/http"

//CdbSourceResponse is struct for /csbsource endpoint response

func cdbSourceHeandler(w http.ResponseWriter, r *http.Request) {
	// request: /cdbsource?name=akton&number=123456789

	// TODO: logic

	// response: JSON {"number": "123456789", "mcc": MCC, "mnc": MNC, "ported": isPortede }
	sendJSON(w, map[string]string{"answer": "pong"})
}
