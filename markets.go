package main

import (
	"net/http"
	"net/url"
	"log"
	"import"
	"os"	
)

const url string = "https://api.coinmarketcap.com/v1/ticker/" 

func GetSpecifiedMarkets() {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	os.Stdout.Write(resp.Body)
}