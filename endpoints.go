package oanda

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"bufio"
)

//FIXME currently you're doing error handling for non 200 status requests,
//however there are other responses oanda can deliver, like
//{"errorMessage":"Timeout waiting for response."}
//shit like that needs to be handled too

//also this error handling did not catch the last service outage yo!!!!
//the code was not working, checked http://api-status.oanda.com/
//and sure enough there is an outage that the code did not detect....
//need that to be working yo!!!!!!!!!!!!

var oandaUrl string = "https://api-fxpractice.oanda.com/v3"
var streamOandaUrl string = "https://stream-fxpractice.oanda.com/v3"
var bearer string = "Bearer " + os.Getenv("OANDA_TOKEN")
var accountId string = os.Getenv("OANDA_ACCOUNT_ID")

//callback function for printing out network requests
func LogComms(req *http.Request, pricesByte []byte, statusCode int, err error) {
	log.Printf("Request: %s\n", req)
	log.Printf("Response: %s\n", string(pricesByte))
	log.Printf("Status Code: %s\n", statusCode)
	log.Printf("GetPricing Response Error: %s\n", err)
}

/*
***************************
prices
***************************
*/

func StreamPricing(instruments ...string) ([]byte, error) {
	client := &http.Client{}
	queryValues := url.Values{}
	instrumentsEncoded := strings.Join(instruments, ",")
	queryValues.Add("instruments", instrumentsEncoded)

	//https://stream-fxtrade.oanda.com/v3/accounts/<ACCOUNT>/pricing/stream?instruments=EUR_USD%2CUSD_CAD"
	//https://stream-fxpractice.oanda.com/v3/accounts/101-001-6395930-001/pricing/stream?instruments=EUR_USD
	//https://stream-fxpractice.oanda.com/accounts/101-001-6395930-001/pricing?instruments=EUR_USD

	req, err := http.NewRequest("GET", streamOandaUrl+"/accounts/"+accountId+
		"/pricing/stream?"+queryValues.Encode(), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	if err != nil {
		return []byte{}, errors.New("GetPricing Error")
	}
	fmt.Println(req)

	resp, err := client.Do(req)

	if err != nil {
		//pricesByte, _ := ioutil.ReadAll(resp.Body)
		//LogComms(req, pricesByte, resp.StatusCode, err)
		fmt.Println("error line 67")
		return []byte{}, errors.New("GetPricing Error")
	}
	defer resp.Body.Close()

		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil{
				fmt.Println("77")
				return []byte{}, errors.New("GetPricing Error")
			}
			//pricesByte, _ := ioutil.ReadAll(line)
			fmt.Println(line)
		}
		//LogComms(req, []byte{}, resp.StatusCode, err)
		return []byte{}, nil
}


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
		LogComms(req, pricesByte, resp.StatusCode, err)
		return []byte{}, errors.New("GetPricing Error")
	} else {
		defer resp.Body.Close()
		pricesByte, _ := ioutil.ReadAll(resp.Body)
		LogComms(req, pricesByte, resp.StatusCode, err)
		return pricesByte, nil
	}
}

/*
***************************
history
***************************
*/

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
		LogComms(req, pricesByte, resp.StatusCode, err)
		return []byte{}, errors.New("GetCandles Error")
	} else {
		defer resp.Body.Close()
		pricesByte, _ := ioutil.ReadAll(resp.Body)
		LogComms(req, pricesByte, resp.StatusCode, err)
		return pricesByte, nil
	}
}

/*
***************************
orders
***************************
*/

func SubmitOrder(orders []byte) ([]byte, error) {
	body := bytes.NewBuffer(orders)
	//FIXME these should be moved somewhere or removed
	fmt.Println(body)
	fmt.Println()
	client := &http.Client{}

	req, err := http.NewRequest("POST", oandaUrl+"/accounts/"+accountId+"/orders", body)

	req.Header.Set("Authorization", bearer)
	req.Header.Set("content-type", "application/json")

	if err != nil {
		return []byte{}, errors.New("SubmitOrder Error")
	}

	if resp, err := client.Do(req); err != nil {
		defer resp.Body.Close()
		ordersByte, _ := ioutil.ReadAll(resp.Body)
		LogComms(req, ordersByte, resp.StatusCode, err)
		return []byte{}, errors.New("SubmitOrder Error")
	} else {
		defer resp.Body.Close()
		ordersByte, _ := ioutil.ReadAll(resp.Body)
		LogComms(req, ordersByte, resp.StatusCode, err)
		return ordersByte, nil
	}
}
