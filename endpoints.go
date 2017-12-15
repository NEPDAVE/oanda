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

func GetPricing(instruments ...string) []byte {
	client := &http.Client{}
	queryValues := url.Values{}

	for _, v := range instruments {
		queryValues.Add("instruments", v)
	}

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
	fmt.Println(string(pricesByte))
	fmt.Println(req)
	fmt.Println(resp.StatusCode)
	return pricesByte
}

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
	fmt.Println(string(pricesByte))
	fmt.Println(req)
	fmt.Println(resp.StatusCode)
	return pricesByte
}
