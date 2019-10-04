package main

import (
	"log"
	"time"
	"net/http"
	"strings"
	"os"
)

func main() {
	log.Println("starting ranking service")

	key := os.Getenv("CRYPTOCOMPARE_KEY")


	if(strings.Compare(key, "") == 0) {
		log.Fatal("No CryptoCompare key set!")
	}
	
	handler := ApiHandler {ApiKey: key}
	
	server := &http.Server{
		Addr:           ":6669",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 32,
	}

	log.Fatal(server.ListenAndServe())
}


