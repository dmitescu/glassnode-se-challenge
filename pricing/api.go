package main

import (
	"net/http"
	
	"encoding/json"
	
	"log"
)

type ApiHandler struct {
	ApiStore PriceStore
}

func (a ApiHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	log.Println("serving request")
	
	encoder := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")

	err := encoder.Encode(a.ApiStore.GetAll())

	if (err != nil) {
		log.Println("error encoding: ", err)
		return
	}
}

