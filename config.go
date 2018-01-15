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

var SellRatePercentage = decimal.NewFromFloat(1.05)

type Config struct {
	Key string `json:"key"`
	Secret string `json:"secret"`
}

type OrderIds struct {
	Ids []string `json:"orderuuids"`
}

func(oids *OrderIds)addId(id string){
	oids.Ids = append(oids.Ids, id)
}

var (
	ErrMissingConfigFile = errors.New("couldn't find config file")
	ErrMissingBittrexTokens = errors.New("missing Bittrex key and secret") 
	ErrMissingOrderIdsFile = errors.New("Missing orderIds file")
	ErrNoOrderIdsPresent = errors.New("No orderIds present in file")
)

func readOrderIds() (orderIds OrderIds, err error){
	if dir, e := homedir.Dir(); e == nil {
		ordersPath := path.Join(dir, OrdersPath)
		if fConf, e := os.Open(ordersPath); e == nil {
			defer fConf.Close()
			err = json.NewDecoder(fConf).Decode(&orderIds)
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

func writeToOrdersFile(orderIds OrderIds) (err error) {

	if dir, e := homedir.Dir(); e == nil {
		ordersPath := path.Join(dir, OrdersPath)
		ordersJason, _ := json.Marshal(orderIds)
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