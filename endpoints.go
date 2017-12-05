package oanda

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var oandaUrl string = "https://api-fxpractice.oanda.com/v3"
var bearer string = "Bearer " + os.Getenv("OANDA_TOKEN")
var accountId string = os.Getenv("OANDA_ACCOUNT_ID")

/*
***************************
prices
***************************
*/

func GetPricing(instrument string) []byte {
	client := &http.Client{}
	queryValues := url.Values{}
	queryValues.Add("instruments", instrument)

	req, err := http.NewRequest("GET", oandaUrl+"/accounts/"+accountId+
		"/pricing?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	pricesByte, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(pricesByte)
	fmt.Println(queryValues)
	fmt.Println(bodyString)
	fmt.Println(resp.StatusCode)
	return pricesByte
}

//https://api-fxpractice.oanda.com/v3/instruments/candles?count=10&granularity=D&instruments=EUR_USD
//https://api-fxpractice.oanda.com/v3/instruments/instrument/candles?count=10&granularity=D&instruments=EUR_USD
//https://api-fxpractice.oanda.com/v3/instrumentsEUR_USD/candles?count=10&granularity=D&instruments=EUR_USD
//"https://api-fxpractice.oanda.com/v3/instruments/USD_JPY/candles?
//count=10&price=A&from=2016-01-01T00%3A00%3A00.000000000Z&granularity=D"

//"https://api-fxpractice.oanda.com/v3"
func GetCandles(instrument string, count string, granularity string) []byte {
	client := &http.Client{}
	queryValues := url.Values{}
	queryValues.Add("instruments", instrument)
	queryValues.Add("count", count)
	queryValues.Add("granularity", granularity)

	req, err := http.NewRequest("GET", oandaUrl+"/instruments"+"/"+instrument+
		"/candles?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	pricesByte, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(pricesByte)
	fmt.Println(queryValues)
	fmt.Println(bodyString)
	fmt.Println(resp.StatusCode)
	fmt.Println(queryValues)
	fmt.Println(req)
	return pricesByte
}
