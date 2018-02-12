package main

import (
	"net/http"
		// "net/url"
	"log"
	"encoding/json"
		// "os"
		// "reflect"
	// "fmt"
	"io/ioutil"
    // "errors"
    "sort"
)
const coinMarketCapUrl string = "https://api.coinmarketcap.com/v1/ticker/"

// specify markets you want to buy from
var Markets = []string { 
	"BTC",
	"XRP",
	"ETH",
	"BCH",
	"ADA",
	"LTC",
	"MIOTA",
	"XEM",
	"XLM",
	"DASH",
	"XMR",
	"NEO",
	"EOS",
	"BTG",
	"QTUM",
	"XRB",
	"TRX",
	"ETC",
	"ICX",
	"LSK",
}

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

func getSpecifiedMarkets() (markets []CoinMarketCapMarket, err error){

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

	}
	// handle error response
	if resp.StatusCode != 200 {
  		bodyError, _ := ioutil.ReadAll(resp.Body)
  		log.Fatal(string(bodyError))
  		return
	}

	//decode payload
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&markets)
	// fmt.Printf("Markets: %+v",)

	return
}

func returnMarketsToBuy() (markets map[string]CoinMarketCapMarket, err error) {

	allMarkets, err := getSpecifiedMarkets()

	sort.Slice(allMarkets, func(i, j int) bool {
  		return allMarkets[i].PercentageChange1h < allMarkets[j].PercentageChange1h
	})

	// for market, _ := range Markets {
		
	// }

	return  
}