package main

import (
	"net/http"
	
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
	
	"log"
)

type ApiHandler struct {}

func GetPrices() []ItemPriceDTO {
	client := &http.Client{}
	request, err := http.NewRequest("GET",
		"http://127.0.0.1:6668",
		nil)

	if (err != nil) {
		log.Println("error:", err)
	}
	
	resp, err := client.Do(request)

	if (err != nil) {
		log.Println("error:", err)
		return nil
	}

	var prices []ItemPriceDTO
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &prices)

	if (err != nil) {
		log.Println("error", err)
		return nil
	}

	return prices
}


func GetRankings(symbols []string) []ItemRankDTO {
	client := &http.Client{}
	request, err := http.NewRequest("GET",
		strings.Join(
			[]string{
				"http://127.0.0.1:6669?syms=",
				strings.Join(symbols, ","),
			},""),
		nil)

	if (err != nil) {
		log.Println("error:", err)
	}
	
	resp, err := client.Do(request)

	if (err != nil) {
		log.Println("error:", err)
		return nil
	}

	var ranks []ItemRankDTO
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ranks)

	if (err != nil) {
		log.Println("error", err)
		return nil
	}

	return ranks
}


func (a ApiHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	log.Println("serving request")
	
	encoder := json.NewEncoder(w)
	var limit int
	var err error

	w.Header().Set("Content-Type", "application/json")

	limit, err = strconv.Atoi(request.FormValue("limit"))
	
	if (err != nil) {
		limit = 10
	}

	prices := GetPrices()

	var symbols []string
	var finalItems []PriceItem

	for _, price := range prices {
		symbols = append(symbols, price.Symbol)
	}

	ranks := GetRankings(symbols)

	for _, rank := range ranks {
		var foundPrice float64

		for _, price := range prices {
			if (price.Symbol == rank.Symbol) {
				foundPrice = price.Price
			}
		}

		finalItems = append(finalItems,
			PriceItem {
				Rank: rank.Rank,
				Symbol: rank.Symbol,
				Price: foundPrice,
			})
	}
	
	err = encoder.Encode(finalItems[:
		int(math.Max(
			float64(limit),
			float64(len(finalItems))))])

	if (err != nil) {
		log.Println("error encoding: ", err)
	}
	// encoder.Write()
}

