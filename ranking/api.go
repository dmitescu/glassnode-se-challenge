package main

import (
	"net/http"
	"net/url"
	
	"encoding/json"
	"io/ioutil"
	"strings"
	
	"log"
)

type ApiHandler struct {
	ApiKey string
}

func (a ApiHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	log.Println("serving request")
	
	encoder := json.NewEncoder(w)
	client := &http.Client{}

	w.Header().Set("Content-Type", "application/json")

	err := request.ParseForm()

	if (err != nil || request.Form["syms"] == nil) {
		log.Println("messed up form data")
	}
	
	form := url.Values {
		"fsyms": strings.Split(request.Form["syms"][0], ","),
		"tsyms": {"USD"},
	}

	apiRequest, _ := http.NewRequest("GET",
		strings.Join(
			[]string{
				"https://min-api.cryptocompare.com/data/pricemultifull?",
				form.Encode(),
			}, ""),
		nil)

	apiRequest.Header.Add("Authorization",
		strings.Join(
			[]string{
				"Apikey",
				a.ApiKey,
			},
		" "))
	apiRequest.Header.Add("Accept", "application/json")
	apiRequest.Header.Add("Limit", "200")

	resp, err := client.Do(apiRequest)

	if (err != nil) {
		log.Println("error:", err)
		return
	}

	priceResponse := &PriceResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, priceResponse)

	log.Println(string(body))

	if (err != nil) {
		log.Println("error", err)
		return
	}

	var prices []Price
	
	for _, price := range priceResponse.Raw {
		prices = append(prices, price.USD)
	}
	
	err = encoder.Encode(DtosFromPrices(prices))

	if (err != nil) {
		log.Println("error encoding: ", err)
		return
	}
}

