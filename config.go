package main

import (
	"errors"
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	"path"
	"os"
	// "github.com/shopspring/decimal"
)

// specify your path relative to your user
const (
	DefaultConfigPath = ".config/bittrex-bot"
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

const SellRatePercentage = 5.0

type Config struct {
	Key string `json:"key"`
	Secret string `json:"secret"`
}

var (
	ErrMissingConfigFile = errors.New("couldn't find config file")
	ErrMissingBittrexTokens = errors.New("missing Bittrex key and secret") 
)


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