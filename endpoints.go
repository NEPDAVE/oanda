package oanda

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var oandaUrl string = "https://api-fxpractice.oanda.com/v3"
var bearer string = "Bearer " + os.Getenv("OANDA_TOKEN")
var accountId string = os.Getenv("OANDA_ACCOUNT_ID")

//FIXME starting to see a lot of repeat code, might want to write a "communications check"
//callback function that prints out and logs the good stuff, IE request, Response
//error codes, maybe even the time
func CommCheck() {
	fmt.Println("Communications Check")
}

/*
***************************
prices
***************************
*/

func GetPricing(instruments ...string) ([]byte, error) {
	client := &http.Client{}
	queryValues := url.Values{}
	instrumentsEncoded := strings.Join(instruments, ",")
	queryValues.Add("instruments", instrumentsEncoded)

	req, err := http.NewRequest("GET", oandaUrl+"/accounts/"+accountId+
		"/pricing?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		return []byte{}, errors.New("GetPricing Error")
	}

	if resp, err := client.Do(req); err != nil {
		defer resp.Body.Close()
		pricesByte, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Request: %s\n", req)
		fmt.Printf("Response: %s\n", string(pricesByte))
		fmt.Printf("Status Code: %s\n", resp.StatusCode)
		fmt.Printf("GetPricing Response Error: %s\n", err)
		return []byte{}, errors.New("GetPricing Error")
	} else {
		defer resp.Body.Close()
		pricesByte, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Request: %s\n", req)
		fmt.Printf("Response: %s\n", string(pricesByte))
		fmt.Printf("Status Code: %s\n", resp.StatusCode)
		fmt.Printf("GetPricing Response Error: %s\n", err)
		return pricesByte, nil
	}
}

func GetCandles(instrument string, count string, granularity string) ([]byte, error) {
	client := &http.Client{}
	queryValues := url.Values{}
	queryValues.Add("instruments", instrument)
	queryValues.Add("count", count)
	queryValues.Add("granularity", granularity)

	req, err := http.NewRequest("GET", oandaUrl+"/instruments"+"/"+instrument+
		"/candles?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		return []byte{}, errors.New("GetPricing Error")
	}

	if resp, err := client.Do(req); err != nil {
		defer resp.Body.Close()
		pricesByte, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Request: %s\n", req)
		fmt.Printf("Response: %s\n", string(pricesByte))
		fmt.Printf("Status Code: %s\n", resp.StatusCode)
		fmt.Printf("GetPricing Response Error: %s\n", err)
		return []byte{}, errors.New("GetCandles Error")
	} else {
		defer resp.Body.Close()
		pricesByte, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Request: %s\n", req)
		fmt.Printf("Response: %s\n", string(pricesByte))
		fmt.Printf("Status Code: %s\n", resp.StatusCode)
		fmt.Printf("GetCandles Response Error: %s\n", err)
		return pricesByte, nil
	}
}
