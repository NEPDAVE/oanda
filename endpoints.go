package oanda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)


var oandaUrl string = "https://api-fxtrade.oanda.com/v3"
var bearer string = "Bearer " + os.Getenv("OANDA_TOKEN")
var accountId string = os.Getenv("OANDA_ACCOUNT_ID")

/*
***************************
prices
***************************
*/

func Pricing(instrument string) []byte {
        client := &http.Client{}
        queryValues := url.Values{}

        req, err := http.NewRequest("GET", oandaUrl+"/accounts/"+accountId+"/pricing"+queryValues.Encode(), nil)
				queryValues.Add("instruments", instrument)
				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", bearer)


				fmt.Println("TEST TEST TEST")
				fmt.Println(req.Header)
				fmt.Println(queryValues)

				resp, err := client.Do(req)

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        defer resp.Body.Close()
        byte, _ := ioutil.ReadAll(resp.Body)
        return byte
}
