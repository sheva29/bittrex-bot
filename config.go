package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"os"
	"path"
)

// specify your path relative to your user
const (
	DefaultConfigPath = ".config/bittrex-bot"
	OrdersPath        = ".config/bittrex-orders"
)

var SellRatePercentage = decimal.NewFromFloat(1.06) //
var BidRatePercentage = decimal.NewFromFloat(1.05)  // we want to sell below our sell rate percentage
var NumberOfTransacions = 4                         // we want to handle only 4 transaction at the time given our budget

type Config struct {
	Key    string
	Secret string
}

type Orders struct {
	Ids map[string]string `json:"orderuuids"`
}

func (oids *Orders) addId(key string, value string) {
	oids.Ids[key] = value
}

var (
	ErrMissingConfigFile    = errors.New("couldn't find config file")
	ErrMissingBittrexTokens = errors.New("missing Bittrex key and secret")
	ErrMissingOrderIdsFile  = errors.New("Missing orderIds file or corrupted file")
	ErrNoOrderIdsPresent    = errors.New("No orderIds present in file")
	ErrCouldntParseFile     = errors.New("Can't parse JSON file for orders")
)

func readOrderIds() (orderIds Orders, err error) {
	if dir, e := homedir.Dir(); e == nil {
		ordersPath := path.Join(dir, OrdersPath)
		if ordersJson, e := os.Open(ordersPath); e == nil {
			defer ordersJson.Close()
			err = json.NewDecoder(ordersJson).Decode(&orderIds)
			fmt.Println("Finishied decoding")
			if err != nil {
				fmt.Println("Inside error")
				err = ErrCouldntParseFile

			}
			// file, err := ioutil.ReadFile(ordersPath)
			// if err != nil {
			// 	fmt.Println("error when converting orders to []byte", err)
			// }
			// err = json.Unmarshal(file, &orderIds)
			// if err != nil {
			// 	fmt.Println("error when unmarshalling orders to", err)
			// }
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
		ordersJason, _ := json.Marshal(orderIds)
		err = ioutil.WriteFile(ordersPath, ordersJason, 0644)
		return
	} else {
		err = errors.New("Couldn't write to Orders file")
		return
	}

}

func readBittrexCredentials() (conf Config, err error) {

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
