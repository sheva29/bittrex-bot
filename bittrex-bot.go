package main

import (
	"fmt"
	// "encoding/json"
	// "os"
	"github.com/toorop/go-bittrex"
	"reflect"
	"github.com/shopspring/decimal"
	// "errors"
)

type CurrentBalance struct {

	Currency string
	Balance decimal.Decimal
	BTHValue decimal.Decimal
	OrderUuid string
	AmountToSell decimal.Decimal
}

func NewCurrentBalance() CurrentBalance {
	return CurrentBalance {
	}
}

func (c *CurrentBalance)NewAmountToSell(sellPercentage decimal.Decimal) {
	c.AmountToSell = c.BTHValue.Mul(SellRatePercentage)
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

	if len(orderIds.Ids) == 0 {
		//then GetBalances and populate object from it
		fmt.Print("No existing balances \n")
	}

	fmt.Printf("balance Ids len: %v\n", len(orderIds.Ids))

	//init bittrex client
	bittrex := bittrex.New(config.Key, config.Secret)
	fmt.Printf("bittrex is %v \n", reflect.TypeOf(bittrex))

	// create slice of balances to populate
	currentBalances := returnBalances(orderIds, bittrex)
	fmt.Printf("Balances: %+v\n", reflect.TypeOf(currentBalances))

	if len(currentBalances) > 0 {
		sellCurrentCoins(&currentBalances, bittrex, &orderIds)
	}

	// let's check if we have a BTC balance
			// check if there's any BTC balance and populate []CurrentBalance
	BTC, err := bittrex.GetBalance("BTC")
	if err != nil {
		fmt.Print(err, "\n")
	}
	// populate BTC into balance
	var BTCBalance = NewCurrentBalance()
	BTCBalance.Balance = BTC.Balance

	// if balance is greater than 0 for BTC, check for markets and buy

	fmt.Printf("BTC balance: %+v\n", BTCBalance)
	// populate CurrentBalance object

	writeToOrdersFile(orderIds)

	if err != nil {

	} else {
		fmt.Print("successfully saved orderIds")
	}
	
}

func sellCurrentCoins(balances *[]CurrentBalance, bittrex *bittrex.Bittrex, orders *Orders){


	for _, balance := range *balances {
		currentSellPrice, err := bittrex.GetTicker(balance.Currency)
		if err != nil{
			fmt.Print(err,"\n")
		}
		// check if value is greater
		checkIfCanSell := balance.AmountToSell.Cmp(currentSellPrice.Bid)
		if checkIfCanSell < 1 {
			fmt.Printf("we can sell for %+v, to sell: %+v - bid: %+v \n", balance.Currency, balance.AmountToSell, currentSellPrice.Bid )

			amount, _ := balance.Balance.Float64()
			bid, _ := currentSellPrice.Bid.Float64() 
			// put sell order
			uuid, err := bittrex.SellLimit(balance.Currency, amount, bid)
			if err != nil {
				fmt.Print(err)
			}
			// assign value to map
			orders.Ids[uuid] = uuid
			
		}else{
			fmt.Printf("we can't sell for %+v, to sell: %+v - bid: %+v \n", balance.Currency, balance.AmountToSell, currentSellPrice.Bid )
		}
	}
}

func returnBalances(orders Orders, bittrex *bittrex.Bittrex) (balances []CurrentBalance) {

	if ids, ok := orders.Ids["orderuuids"].(map[string]interface{}); ok{		

		for k, v := range ids {
			fmt.Printf("orderuuid: key: %s - value: %s\n", k , v)
			order, err := bittrex.GetOrder(k)

			if err != nil{
				fmt.Print(err, "\n")
			}
			currentBalance := CurrentBalance{
				Currency : order.Exchange,
				Balance : order.Quantity,
				BTHValue : order.Limit,
				OrderUuid : order.OrderUuid,
			}

			currentBalance.NewAmountToSell(SellRatePercentage)
			fmt.Printf("currentBalance: %+v\n", currentBalance)
			balances = append(balances, currentBalance)
		}
	}
	return
}