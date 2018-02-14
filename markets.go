package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"strconv"
	"time"
)

const (
	coinMarketCapUrl = "https://api.coinmarketcap.com/v1/ticker/"
	csvLocation      = "/csv/"
)

var today string = time.Now().Format("2006-01-02")

// specify markets you want to buy from
var Markets = []string{
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
	Id                  string  `json:"id"`
	Name                string  `json:"name"`
	Symbol              string  `json:"symbol"`
	Rank                int     `json:"rank,string"`
	PriceUsd            float32 `json:"price_usd,string"`
	PriceBtc            float32 `json:"price_btc,string"`
	Volume24hUsd        float32 `json:"24h_volume_usd,string"`
	MarketCap           float32 `json:"24h_volume_usd,string"`
	AvailableSupply     float32 `json:"available_supply,string"`
	TotalSupply         float32 `json:"total_supply,string"`
	PercentageChange1h  float32 `json:"percent_change_1h,string"`
	PercentageChange24h float32 `json:"percent_change_24h,string"`
	PercentageChange7d  float32 `json:"percent_change_7d,string"`
	LastUpdated         int     `json:"last_updated,string"`
}

func getSpecifiedMarkets() (markets []CoinMarketCapMarket, err error) {

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

/*
	we will use this to get data to come up with an estimated value
*/
func writeMarketValueToFile(market CoinMarketCapMarket) {

	if dir, e := homedir.Dir(); e == nil {
		csvFilePath := path.Join(dir, csvLocation)
		// if file doesn't exist, add header
		path := csvFilePath + "/" + market.Symbol + "-" + today + ".csv"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// path/to/whatever does not exist
			file, err := os.Create(path)
			if err != nil {
				fmt.Println("Cannot create file", err)
			}
			defer file.Close()

			writer := csv.NewWriter(file)
			defer writer.Flush()

			data := [][]string{{"Price", "Time"}}
			for _, record := range data {
				if err := writer.Write(record); err != nil {
					fmt.Println("Cannot add csv headers")
				}
			}

		} else {
			// it exists, then insert values regularly
			file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
			if err != nil {
				fmt.Println("Cannot open file", err)
			}
			defer file.Close()

			writer := csv.NewWriter(file)
			defer writer.Flush()

			BTCValue := strconv.FormatFloat(float64(market.PriceBtc), 'f', -1, 32)
			currentTime := time.Now().Format("15:04:05")
			data := [][]string{{BTCValue, currentTime}}
			for _, record := range data {
				if err := writer.Write(record); err != nil {
					fmt.Println("Cannot add csv record")
				}
			}
		}
	}
}

func returnMarketsToBuy() (markets map[string]CoinMarketCapMarket, err error) {

	allMarkets, err := getSpecifiedMarkets()

	sort.Slice(allMarkets, func(i, j int) bool {
		return allMarkets[i].PercentageChange1h < allMarkets[j].PercentageChange1h
	})

	// TODO:
	// for market, _ := range Markets {

	// }

	return
}
