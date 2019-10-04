package main

import (
	"log"
	"time"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"strings"
	"os"
)

func StoreUpdater(store PriceStore, apiKey string) {
	client := &http.Client{}

	for (true) {
		request, _ := http.NewRequest("GET",
			"https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest",
			nil)
		request.Header.Add("X-CMC_PRO_API_KEY", apiKey)
		request.Header.Add("Accept", "application/json")
		request.Header.Add("Limit", "200")
		
		resp, err := client.Do(request)

		if (err != nil) {
			log.Println("error:", err)
		}

		prices := &PriceRequest{}
		body, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, prices)

		if (err != nil) {
			log.Println("error", err)
		} else {
			for _, price := range prices.Data {
				store.Put( DtoFromItem(price))
			}
			time.Sleep(10 * time.Second)
		}
	}
}

func main() {
	log.Println("starting pricing service")

	apiKey := os.Getenv("COINMARKETCAP_KEY")
	if (strings.Compare(apiKey, "") == 0) {
		log.Fatal("CoinMarketCap Key not set")
	}
	
	store := PriceStore{}
	store.Init()

	handler := ApiHandler {
		ApiStore: store,
	}

	go StoreUpdater(store, apiKey)
	
	server := &http.Server{
		Addr:           ":6668",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 32,
	}

	log.Fatal(server.ListenAndServe())
}


