package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/toorop/go-bittrex"
	"reflect"
)

type CurrentBalance struct {
	Currency     string
	Balance      decimal.Decimal
	BTHValue     decimal.Decimal
	OrderUuid    string
	AmountToSell decimal.Decimal
	AmountForBid decimal.Decimal
}

func NewCurrentBalance() CurrentBalance {
	return CurrentBalance{}
}

func (c *CurrentBalance) NewAmountToSell(sellPercentage decimal.Decimal) {
	c.AmountToSell = c.BTHValue.Mul(sellPercentage)
}

func (c *CurrentBalance) NewAmountForBid(bidPercentage decimal.Decimal) {
	c.AmountForBid = c.BTHValue.Mul(bidPercentage)
}

func main() {

	fmt.Println("started successfully")

	// get keys
	config, err := readBittrexCredentials()
	if err != nil {
		fmt.Println(err)
	}

	/*
		get order ids
	*/
	// read order numbers
	orderIds, err := readOrderIds()
	fmt.Println("orderIds:", orderIds)
	if err != nil {
		fmt.Println("Error parsing order Ids:", err)
	}

	if len(orderIds.Ids) == 0 {
		//then GetBalances and populate object from it
		fmt.Print("No existing balances \n")
	}

	fmt.Printf("balance Ids len: %v\n", len(orderIds.Ids))

	//init bittrex client
	bittrex := bittrex.New(config.Key, config.Secret)
	fmt.Printf("bittrex is %v \n", reflect.TypeOf(bittrex))

	/*
		get balances
	*/
	// create slice of balances to populate
	currentBalances := returnBalances(orderIds, bittrex)
	fmt.Printf("Balances: %+v\n", reflect.TypeOf(currentBalances))

	/*
		sell
	*/
	balances := len(currentBalances)
	if balances > 0 {
		sellCurrencies(&currentBalances, bittrex, &orderIds)
	}
	/*
		check BTC balance
	*/
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

	/*
		save orders
	*/
	// we save our Ids in case they have change
	writeToOrdersFile(orderIds)

	if err != nil {

	} else {
		fmt.Print("successfully saved orderIds")
	}

	if markets, err := getSpecifiedMarkets(); err == nil {
		for i := 1; i < 26; i++ {
			writeMarketValueToFile(markets[i])
		}
	} else {
		fmt.Println(err)
	}

}

func sellCurrencies(balances *[]CurrentBalance, bittrex *bittrex.Bittrex, orders *Orders) {

	for _, balance := range *balances {
		currentSellPrice, err := bittrex.GetTicker(balance.Currency)
		if err != nil {
			fmt.Println("couldn't get order for", balance.Currency, err)
		}
		// check if value is greater
		if balance.AmountToSell.Cmp(currentSellPrice.Bid) < 1 {
			fmt.Printf("we can sell for %+v, to sell: %+v - bid: %+v \n", balance.Currency, balance.AmountToSell, currentSellPrice.Bid)

			// put sell order
			uuid, err := bittrex.SellLimit(balance.Currency, balance.Balance, balance.AmountForBid)
			if err != nil {
				fmt.Println(err)
			}
			// assign value to map
			orders.Ids[uuid] = uuid

			//TODO: add mechanism
			// we delete old uuid
			delete(orders.Ids, balance.OrderUuid)

		} else {
			fmt.Printf("we can't sell for %+v, to sell: %+v - bid: %+v \n", balance.Currency, balance.AmountToSell, currentSellPrice.Bid)
		}
	}
}

func returnBalances(orders Orders, bittrex *bittrex.Bittrex) (balances []CurrentBalance) {

	for k, v := range orders.Ids {
		fmt.Printf("orderuuid: key: %s - value: %s\n", k, v)
		order, err := bittrex.GetOrder(k)

		if err != nil {
			fmt.Print(err, "\n")
		}
		currentBalance := CurrentBalance{
			Currency:  order.Exchange,
			Balance:   order.Quantity,
			BTHValue:  order.Limit,
			OrderUuid: order.OrderUuid,
		}

		currentBalance.NewAmountToSell(SellRatePercentage)
		currentBalance.NewAmountForBid(BidRatePercentage)
		fmt.Printf("currentBalance: %+v\n", currentBalance)
		balances = append(balances, currentBalance)
	}
	return
}

func buyCurrencies() {
	//TODO: build method based on expected value analysis
	return
}
