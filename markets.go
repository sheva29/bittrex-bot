package main

import (
	"net/http"
		// "net/url"
	"log"
	"encoding/json"
		// "os"
		// "reflect"
	// "fmt"
	// "io/ioutil"
    // "errors"
)
const coinMarketCapUrl string = "https://api.coinmarketcap.com/v1/ticker/"

type CoinMarketCapMarket struct {
	Id 	string	`json:"id"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Rank int `json:"rank,string"`
	PriceUsd float32 `json:"price_usd,string"`
	PriceBtc float32 `json:"price_btc,string"`
	Volume24hUsd float32 `json:"24h_volume_usd,string"`
	MarketCap float32 `json:"24h_volume_usd,string"`
	AvailableSupply float32 `json:"available_supply,string"`
	TotalSupply float32 `json:"total_supply,string"`
	PercentageChange1h float32 `json:"percent_change_1h,string"`
	PercentageChange24h float32 `json:"percent_change_24h,string"`
	PercentageChange7d float32 `json:"percent_change_7d,string"`
	LastUpdated int `json:"last_updated,string"`
} 

func GetSpecifiedMarkets() (markets []CoinMarketCapMarket, err error){

	// build request
	req, er := http.NewRequest("GET", coinMarketCapUrl, nil)
	if er != nil {
		log.Fatal("NewRequest: ", er)
		err = er
		return
	}

	// get response
	client := &http.Client{}
	resp, er := client.Do(req)
	if er != nil {
		log.Fatal("Do: ", er)
		err = er
		return
	}

	//decode payload
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&markets)
	// fmt.Printf("Markets: %+v",)

	return
}