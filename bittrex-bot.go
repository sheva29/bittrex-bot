package main

import (
	"fmt"
	"encoding/json"
	"os"
	"errors"
	"github.com/mitchellh/go-homedir"
	"path"
)

type Config struct {
	Key string `json:"key"`
	Secret string `json:"secret"`
}

const (
	DefaultConfigPath = ".config/bittrex-bot"
)

var (
	ErrMissingConfigFile = errors.New("couldn't find config file")
	ErrMissingBittrexTokens = errors.New("missing Bittrex key and secret") 
)

func main() {

	fmt.Printf("started successfully")
	if path, er := readBittrexCredentials(DefaultConfigPath); er == nil{
		fmt.Printf("path: %s", path)
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
