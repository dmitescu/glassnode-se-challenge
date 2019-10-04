package main

type PriceItem struct {
	Rank int `json:"rank"`
	Symbol string `json:"symbol"`
	Price float64 `json:"price"`
}

type ItemPriceDTO struct {
	Symbol string `json:"symbol"`
	Price float64 `json:"price"`
}

type ItemRankDTO struct {
	Symbol string `json:"symbol"`
	Rank int `json:"rank"`
}
