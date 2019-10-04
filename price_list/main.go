package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("serving on port 6667")
	
	handler := ApiHandler {}
	
	server := &http.Server{
		Addr:           ":6667",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 32,
	}

	log.Fatal(server.ListenAndServe())
}
