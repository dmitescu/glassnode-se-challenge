package main

import (
	"log"
	"github.com/patrickmn/go-cache"
	"time"
)

type PriceStore struct {
	Prices *cache.Cache
}

func (store *PriceStore) Init() {
	store.Prices = cache.New(cache.NoExpiration, time.Minute)
}

func (store *PriceStore) Put(item PriceItemDTO) {
	log.Println("inserting ", item)
	store.Prices.Set(item.Symbol, item, cache.DefaultExpiration)
}

func (store *PriceStore) GetAll() []PriceItemDTO {
	var priceItems []PriceItemDTO
	items := store.Prices.Items()

	for _, i := range items {
		priceItem := i.Object.(PriceItemDTO)
		priceItems = append(priceItems, priceItem)
	}

	return priceItems
}

