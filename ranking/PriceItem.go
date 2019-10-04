package main

import "sort"

type PriceResponse struct {
	Raw map[string]USDPrice `json:"RAW"`
}

type USDPrice struct {
	USD Price `json:"USD"`
}

type Price struct {
	Symbol string `json:"FROMSYMBOL"`
	Price float64 `json:"PRICE"`
}
	
type PriceItemDTO struct {
	Symbol string `json:"symbol"`
	Rank int `json:"rank"`
}

type ByPrice []Price

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].Price > a[j].Price }

func DtosFromPrices(items []Price) []PriceItemDTO {
	sort.Sort(ByPrice(items))
	var dtos []PriceItemDTO
	
	for index, item := range items {
		dtos = append(dtos,
			PriceItemDTO {
				Symbol: item.Symbol,
				Rank: index + 1,
			})
	}

	return dtos
}
