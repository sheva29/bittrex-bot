package main

import (
	"errors"
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	"path"
	"os"
	"github.com/shopspring/decimal"
	"io/ioutil"	
)

// specify your path relative to your user
const (
	DefaultConfigPath = ".config/bittrex-bot"
	OrdersPath = ".config/bittrex-orders"
)

var SellRatePercentage = decimal.NewFromFloat(1.06) //
var BidRatePercentage = decimal.NewFromFloat(1.05) // we want to sell below our sell rate percentage
var NumberOfTransacions = 4 // we want to handle only 4 transaction at the time given our budget

type Config struct {
	Key string
	Secret string 
}

type Orders struct {
	Ids map[string]string `json:"orderuuids"`
}

func(oids *Orders)addId(key string, value string){
	oids.Ids[key] = value
}

var (
	ErrMissingConfigFile = errors.New("couldn't find config file")
	ErrMissingBittrexTokens = errors.New("missing Bittrex key and secret") 
	ErrMissingOrderIdsFile = errors.New("Missing orderIds file or corrupted file")
	ErrNoOrderIdsPresent = errors.New("No orderIds present in file")
)

func readOrderIds() (orderIds Orders, err error){
	if dir, e := homedir.Dir(); e == nil {
		ordersPath := path.Join(dir, OrdersPath)
		if fConf, e := os.Open(ordersPath); e == nil {
			defer fConf.Close()
			err = json.NewDecoder(fConf).Decode(&orderIds.Ids)
		}
	} else {
		err = ErrMissingOrderIdsFile
		return
	}

	if len(orderIds.Ids) == 0 {
		err = ErrNoOrderIdsPresent
		return
	}

	return
}

func writeToOrdersFile(orderIds Orders) (err error) {

	if dir, e := homedir.Dir(); e == nil {
		ordersPath := path.Join(dir, OrdersPath)
		ordersJason, _ := json.Marshal(orderIds.Ids)
		err = ioutil.WriteFile(ordersPath, ordersJason, 0644)
		return
	} else {
		err = errors.New("Couldn't write to Orders file")
		return
	}

}


func readBittrexCredentials() (conf Config, err error ){

	if dir, e := homedir.Dir(); e == nil {
		expandedPath := path.Join(dir, DefaultConfigPath)
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