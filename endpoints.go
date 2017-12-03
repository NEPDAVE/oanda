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

	fmt.Println("GET PRICING!!!")
	fmt.Println(req.Header)
	fmt.Println(queryValues)


	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	pricesByte, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(pricesByte)
	fmt.Println("ENDPOINTS LINE 47")
	fmt.Println(bodyString)
	fmt.Println(resp.StatusCode)
	return pricesByte
}
