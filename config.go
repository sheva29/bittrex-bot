package main

import (
	"errors"
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	"path"
	"os"
)

// specify your path relative to your user
const (
	DefaultConfigPath = ".config/bittrex-bot"
)

// specify markets you want to buy from

var Markets = []string { "BTC", "ETC"}

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