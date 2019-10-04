package main

type PriceRequest struct {
	Data []PriceItem `json:"data"`
}

type PriceItem struct {
	Symbol string `json:"symbol"`
	Quote PriceQuote `json:"quote"`
}

type PriceQuote struct {
	USD USDQuote `json:"USD"`
}

type USDQuote struct {
	Price float64 `json:"price"`
}

type PriceItemDTO struct {
	Symbol string `json:"symbol"`
	Price float64 `json:"price"`
}

func DtoFromItem(item PriceItem) PriceItemDTO {
	return PriceItemDTO {
		Symbol: item.Symbol,
		Price: item.Quote.USD.Price,
	}
}
