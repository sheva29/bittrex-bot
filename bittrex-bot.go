package main

import (
	"fmt"
	// "encoding/json"
	// "os"
	"github.com/toorop/go-bittrex"
	"reflect"
	"github.com/shopspring/decimal"
	"errors"
)

type CurrentBalance struct {

	Currency string
	Balance decimal.Decimal
	BTHValue float32
	OrderUuid string
	MartketTicker string 
	InitialLimitBuy decimal.Decimal
	AmountToSell decimal.Decimal
}

func NewCurrentBalance() CurrentBalance {
	return CurrentBalance {
		MartketTicker: "BTC-",
	}
}

func (c *CurrentBalance)NewAmountToSell(sellPercentage decimal.Decimal) {
	c.AmountToSell = c.InitialLimitBuy.Mul(SellRatePercentage)
}


func main() {

	fmt.Println("started successfully")

	// get keys
	config, err := readBittrexCredentials()
	if  err != nil{
		fmt.Println(err)
	}

	// read order numbers
	orderIds, err := readOrderIds()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("order Ids: %+v", orderIds)
	orderIds.addId("thisIsANewOrderId")

	f := writeToOrdersFile(orderIds)
	if f != nil {
		fmt.Print(f, "\n")
	}
	/*

		init bittrex client 

	*/
	bittrex := bittrex.New(config.Key, config.Secret)
	fmt.Printf("bittrex is %v \n", reflect.TypeOf(bittrex))


	/*
		Printing current balances
	*/
	// fmt.Printf("Markets: %+v \n", markets)
	fmt.Print("\n")
	balances, err := bittrex.GetBalances()
	if err != nil{
		fmt.Println(err)
	}

	/*
	
	Get User balance
		
	*/
	// fmt.Printf("balance %+v\n", balances)
	userBalances, err := checkBalances(balances)

	if err != nil {
		fmt.Printf( "Balance error: %+v \n", err)
	}

	for _, userBalance := range userBalances{
		orderHistory, err := bittrex.GetOrderHistory(userBalance.MartketTicker)
		if err != nil{
			fmt.Println(err)
		}
		if len (orderHistory) > 0 {

			for _, order := range orderHistory {
				if userBalance.Balance.Equals(order.Quantity) {
					userBalance.OrderUuid = order.OrderUuid
					userBalance.InitialLimitBuy = order.Limit
					userBalance.NewAmountToSell(SellRatePercentage)
				}
			}
		} 
		fmt.Printf("UserBalance Object: %+v\n",userBalance)

		ticker, err := bittrex.GetTicker(userBalance.MartketTicker)

		if err != nil {
			fmt.Print(err,"\n")
		}

		fmt.Printf("ticker for %v : %+v \n", userBalance.Currency, ticker)
		fmt.Printf("amount to buy: %v ticker bid: %v \n",  userBalance.AmountToSell, ticker.Bid)

		checkIfCanSell := userBalance.AmountToSell.Cmp(ticker.Bid)
		fmt.Printf("checkIfCanSell: %v\n", checkIfCanSell)
		if  checkIfCanSell == 0 || checkIfCanSell < 1 {
			fmt.Print("We can buy \n")
		} else {
			fmt.Print("We aren't buying \n")
		}

	}	
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
			var cb CurrentBalance = NewCurrentBalance()
			if balance.Currency != "BTC"{
				cb.Balance = balance.Balance
				cb.Currency = balance.Currency
				cb.MartketTicker += balance.Currency
				currentBalances = append(currentBalances, cb)
				// fmt.Printf("Currency: %v \n", balance.Currency)
				// fmt.Printf("Balance: %v \n", balance.Balance)
			}
		}

	}else{
		err = errors.New("Balance is empty. check your Bittrex account for more information")
		return
	}

	return
}