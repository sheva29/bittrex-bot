package main

import (
	"fmt"
	"encoding/json"
	"os"
	"github.com/toorop/go-bittrex"
	"reflect"
	"github.com/shopspring/decimal"
	"errors"
)

// var MarketTickersBTC = map[string]string {
// 	"BTC-ETH": "BTC-ETH",
// 	"BTC-ETC": "BTC-ETC",
// 	"BTC-XRP": "BTC-XRP",
// 	"BTC-LTC": "BTC-LTC",
// }

var MarketTickerBTC string = "BTC-"

type CurrentBalance struct {

	Currency string
	Balance decimal.Decimal
	BTHValue float32
	OrderUuid string

}

func main() {

	fmt.Println("started successfully")
	test()
	config, err := readBittrexCredentials()
	if  err != nil{
		fmt.Println(err)
	}
	bittrex := bittrex.New(config.Key, config.Secret)
	fmt.Print("bittrex is ", reflect.TypeOf(bittrex))

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
	// marketSummaries, err := bittrex.GetMarketSummaries()
	// fmt.Print("Type of markets: ", reflect.TypeOf(marketSummaries))

	// getSelectedMarkets(marketSummaries)

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

	markets, err := GetSpecifiedMarkets()
	marketsFormatted, err := json.MarshalIndent(markets, "", " ") 
	os.Stdout.Write(marketsFormatted)
	fmt.Printf("\n")

	for _, element := range markets {
		fmt.Printf("%s \n", element.Symbol)			
	}
	/*
		Printing current balances
	*/
	// fmt.Printf("Markets: %+v \n", markets)
	fmt.Print("")
	balances, err := bittrex.GetBalances()
	if err != nil{
		fmt.Println(err)
	}

	fmt.Printf("balance %+v\n", balances)
	checkBalances(balances)
	// fmt.Print("balance object: ", reflect.TypeOf(balances))

	// fmt.Printf("balance %+v\n", balances)

	orderHistory, err := bittrex.GetOrderHistory("BTC-ADA")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%+v\n",orderHistory)

	
}

func test() {
	fmt.Print("I'm being called")
}

// func getSelectedMarkets(allMarketSummaries []bittrex.MarketSummary) {
// 	// for
// 	fmt.Print("Inside the function")
// 	for _, market := range allMarketSummaries{
// 		// fmt.Printf("%+v\n", market.MarketName)
// 		if _, ok := MarketTickersBTC[market.MarketName]; ok{
// 			fmt.Printf("%+v\n", market)
// 		}
// 	} 
// }

func checkBalances(balances []bittrex.Balance) ( currentBalances []CurrentBalance, err error){
	if len(balances) > 0 {

		fmt.Printf("we have : %d balances for the following coins: \n", len(balances))

		for _, balance := range balances {
			var cb CurrentBalance
			if balance.Currency != "BTC"{
				cb.Balance = balance.Balance
				cb.Currency = balance.Currency
				currentBalances = append(currentBalances, cb)
				fmt.Printf("Currency: %v \n", balance.Currency)
				fmt.Printf("Balance: %v \n", balance.Balance)
			}
		}
	}else{
		err = errors.New("Balance is empty. check website")
		return
	}




	return
}
