package main

import (
	"fmt"
	"encoding/json"
	"os"
	"errors"
	"github.com/toorop/go-bittrex"
	"github.com/mitchellh/go-homedir"
	"path"
	"reflect"
)

type Config struct {
	Key string `json:"key"`
	Secret string `json:"secret"`
}

const (
	DefaultConfigPath = ".config/bittrex-bot"
)

var MarketTickersBTC = map[string]string {
	"BTC-ETH": "BTC-ETH",
	"BTC-ETC": "BTC-ETC",
	"BTC-XRP": "BTC-XRP",
	"BTC-LTC": "BTC-LTC",
}

var (
	ErrMissingConfigFile = errors.New("couldn't find config file")
	ErrMissingBittrexTokens = errors.New("missing Bittrex key and secret") 
)

func main() {

	fmt.Println("started successfully")
	test()
	config, err := readBittrexCredentials(DefaultConfigPath)
	if  err != nil{
		fmt.Println(err)
	}
	bittrex := bittrex.New(config.Key, config.Secret)

	/*
	 *	Markets
	 {
	  "MarketCurrency": "ENG",
	  "BaseCurrency": "ETH",
	  "MarketCurrencyLong": "Enigma",
	  "BaseCurrencyLong": "Ethereum",
	  "MinTradeSize": "13.06241877",
	  "MarketName": "ETH-ENG",
	  "IsActive": true,
	  "Notice": "",
	  "IsSponsored": false,
	  "LogoUrl": "https://bittrexblobstorage.blob.core.windows.net/public/a7cb51db-1e6d-47f5-89b5-afbfc766b01b.png"
	 }
	 *
	 */
	// markets, err := bittrex.GetMarkets()
	// marketsFormatted, err := json.MarshalIndent(markets, "", " ") 
	// if err == nil{
	// 	 fmt.Println("Error: ", err)
	// }
	// fmt.Print("Type of markets: ", reflect.TypeOf(markets))
	// os.Stdout.Write(marketsFormatted)

	/*
	 *	Market Summary
	  {
	  "MarketName": "ETH-POWR",
	  "High": "0.00127",
	  "Low": "0.00116145",
	  "Ask": "0.00124166",
	  "Bid": "0.0012055",
	  "OpenBuyOrders": 146,
	  "OpenSellOrders": 470,
	  "Volume": "226245.06664889",
	  "Last": "0.00121984",
	  "BaseVolume": "276.2336568",
	  "PrevDay": "0.00123284",
	  "TimeStamp": "2017-12-27T17:22:51.757"
	 }
	 *
	 */
	marketSummaries, err := bittrex.GetMarketSummaries()
	fmt.Print("Type of markets: ", reflect.TypeOf(marketSummaries))

	getSelectedMarkets(marketSummaries)

	// fmt.Printf("Type of +%v", reflect.TypeOf(summary))

	// marketSummariesFormatted, err := json.MarshalIndent(marketSummaries, "", " ")
	// os.Stdout.Write(marketSummariesFormatted)
	// fmt.Printf("marketSummaries: %+v\n", marketSummaries)

	// ETHTicker, err := bittrex.GetTicker(MarketTickersBTC["ETH"])
	// ETHTickerFormatted, err := json.Formattedjson.MarshalIndent(ETHTicker, "", " ")
	
	// if err == nil {
	// 	fmt.Println("Erroron ticker: ", err)
	// }
	// fmt.Printf("ETH: %+v\n", ETHTicker)

	// myBalance, err := bittrex.GetBalances()
	// fmt.Printf("Balance: %+v\n", myBalance)

	// orderHistory, err := bittrex.GetOrderHistory("BTC-ETH")
	// fmt.Printf("orderHistory: %+v\n", orderHistory)

	// orderETH, err := bittrex.GetOrderBookBuySell("BTC-ETH", "buy")
	// fmt.Printf("order buy ETH: : %+v\n", orderb)
	// ETHOrderMarketFormatted , err := json.MarshalIndent(orderETH, "", " ")
	// os.Stdout.Write(ETHOrderMarketFormatted)
}

func test() {
	fmt.Print("I'm being called")
}

func getSelectedMarkets(allMarketSummaries []bittrex.MarketSummary) {
	// for
	fmt.Print("Inside the function")
	for _, market := range allMarketSummaries{
		// fmt.Printf("%+v\n", market.MarketName)
		if _, ok := MarketTickersBTC[market.MarketName]; ok{
			fmt.Printf("%+v\n", market)
		}
	} 
}


func readBittrexCredentials(providedPath string) (conf Config, err error ){

	if dir, e := homedir.Dir(); e == nil {
		expandedPath := path.Join(dir, providedPath)
		if fConf, e := os.Open(expandedPath); e == nil {
			defer fConf.Close()
			err = json.NewDecoder(fConf).Decode(&conf)

		}
	} else {
		err = ErrMissingConfigFile
		return
	}

	if conf.Key == "" {
		err = ErrMissingBittrexTokens
		return
	}

	return
}
