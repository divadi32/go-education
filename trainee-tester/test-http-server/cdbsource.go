package main

import "net/http"

//CdbSourceResponse is struct for /cdbsource endpoint response

type CdbSourceResponse struct {
	Number string `json:"number"`
	MCC    string `json:"mcc"`
	Mnc    string `json:"mnc"`
	Ported bool   `json:"ported"`
}

func cdbSourceHeandler(w http.ResponseWriter, r *http.Request) {
	// request: /cdbsource?name=akton&number=123456789
	name := r.URL.Query().Get("name")
	number := r.URL.Query().Get("number")

	// TODO: logic
	response, err := cdbSourceLogic(name, number)
	if err != nil {
		sendJSON(http.StatusInternalServerError, w, nil)
	}
	// response: JSON {"number": "123456789", "mcc": MCC, "mnc": MNC, "ported": isPortede }
	sendJSON(http.StatusOK, w, response)
}

func cdbSourceLogic(name, number string) (response *CdbSourceResponse, err error) {
	response = &CdbSourceResponse{
		Number: number,
	}
	return
}
