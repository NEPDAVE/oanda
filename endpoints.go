package oanda

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var oandaUrl string = "https://api-fxpractice.oanda.com/v3"
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
order
***************************
*/

//FIXME the error handling in this file is not correct! Your git pull never worked!!!!!!


//FIXME think about creating a CreateBuyOrder and CreateSellOrder func. This will
//make things more readable. Also think about error handling and possibly another
//function that will prepare the order and another function that will execute it
//also take a look at the data structure and make sure it's getting marshalled
//correctly also add a func to Unmarshal the data after placing an order
//total side note also lookinto coding the double bb
func CreateBuyOrder(bid float64, ask float64, instrument string, units int) Orders {
	//FIXME stopLossPrice and takeProfitPrice are hardcoded to certain ratios
	//this may not be the best way...
  targetPrice := bid
	stopLossPrice := bid - .00002
	takeProfitPrice := bid + (ask - bid) - .000002
  stopLoss := StopLossOnFill{TimeInForce: "GTC", Price: stopLossPrice}
	takeProfit := TakeProfitOnFill{TimeInForce: "GTC", Price: takeProfitPrice}
	orderData := Order{
		Price:            targetPrice,
		StopLossOnFill:   stopLoss,
		TakeProfitOnFill: takeProfit,
		TimeInForce:      "FOK",
		Instrument:       instrument,
		Type:             "LIMIT",
		PositionFill:     "DEFAULT"}
	order := OrderBody{Order: orderData}

	jsonOrders, err := json.Marshal(order)
	if err != nil {
		log.Printf("Json Marshal Error: %s\n", err)
	}

	return jsonOrders

}

//FIXME remember that this units param should be negative
func CreateSellOrder(bid float64, ask float64, instrument string, units int) Orders {
	//FIXME stopLossPrice and takeProfitPrice are hardcoded to certain ratios
	//this may not be the best way...
	targetPrice = ask
	stopLossPrice = ask + .00002
	takeProfitPrice = bid - (ask - bid) + .000002
  stopLoss := StopLossOnFill{TimeInForce: "GTC", Price: stopLossPrice}
	takeProfit := TakeProfitOnFill{TimeInForce: "GTC", Price: takeProfitPrice}
	orderData := Order{
		Price:            targetPrice,
		StopLossOnFill:   stopLoss,
		TakeProfitOnFill: takeProfit,
		TimeInForce:      "FOK",
		Instrument:       instrument,
		Type:             "LIMIT",
		PositionFill:     "DEFAULT"}
	order := OrderBody{Order: orderData}

	jsonOrders, err := json.Marshal(order)
	if err != nil {
		log.Printf("Json Marshal Error: %s\n", err)
	}

	return jsonOrders

}

func SubmitOrder(orders Orders) ([]byte, error) {
	body := bytes.NewBuffer(jsonOrders)
	//FIXME these should be moved somewhere or removed
	fmt.Println(body)
	fmt.Println()
	client := &http.Client{}

	req, err := http.NewRequest("POST", oandaUrl+"/accounts/"+accountId+"/orders", body)
	req.Header.Set("Authorization", bearer)
	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	byte, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte))

}


func CreateOrder(bid float64, ask float64, instrument string, side string, units int) {
	var targetPrice float64
	var stopLossPrice float64
	var takeProfitPrice float64

	if side == "buy" {
		targetPrice = bid
		stopLossPrice = bid - .00002
		takeProfitPrice = bid + (ask - bid) - .000002
	} else if side == "sell" {
		targetPrice = ask
		stopLossPrice = ask + .00002
		takeProfitPrice = bid - (ask - bid) + .000002
		units = units - (units * 2)
	}

	stopLoss := StopLossOnFill{TimeInForce: "GTC", Price: stopLossPrice}
	takeProfit := TakeProfitOnFill{TimeInForce: "GTC", Price: takeProfitPrice}
	orderData := Order{
		Price:            targetPrice,
		StopLossOnFill:   stopLoss,
		TakeProfitOnFill: takeProfit,
		TimeInForce:      "FOK",
		Instrument:       instrument,
		Type:             "LIMIT",
		PositionFill:     "DEFAULT"}
	order := OrderBody{Order: orderData}

	jsonOrderBody, _ := json.Marshal(order)
	body := bytes.NewBuffer(jsonOrderBody)
	fmt.Println(body)
	fmt.Println()
	client := &http.Client{}

	req, err := http.NewRequest("POST", oandaUrl+"/accounts/"+accountId+"/orders", body)
	req.Header.Set("Authorization", bearer)
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	byte, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte))
}
