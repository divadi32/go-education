package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/divadi32/go-education/trainee-tester/test-http-server/cdbeasy"
)

// TODO: get path from akton

const cdbDirPath = "/user/dmitry/code/horisen/cdbdata"

//CdbSourceResponse is struct for /cdbsource endpoint response

type CdbSourceResponse struct {
	Number string `json:"number"`
	Mcc    string `json:"mcc"`
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
	switch name {
	case "akton":
		response, err = cdbSourceSeach("akton.cdb", number)
	default:
		response = &CdbSourceResponse{
			Number: number,
		}

	}
	return
}

func cdbSourceSeach(cdbFile, number string) (response *CdbSourceResponse, err error) {
	response = &CdbSourceResponse{
		Number: number,
		Ported: false,
	}

	cdbFilePath := cdbDirPath + "/" + cdbFile

	//  +12,15:385992026820->mcc=219&mnc=010

	rec, err := cdbeasy.FindOne(cdbFilePath, number)
	if err != nil || rec == "" {
		return
	}

	log.Println(rec)

	recMap, err := url.ParseQuery(rec)
	if err != nil {
		return
	}
	log.Println(recMap)

	// mcc-210&mnc=010
	response.Mcc = recMap["mcc"][0]
	response.Mnc = recMap["mnc"][0]
	response.Ported = true

	return
}
