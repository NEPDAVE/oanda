package oanda

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPricing(t *testing.T) {
	p, err := ioutil.ReadFile("pricing.json")

	if err != nil {
		log.Fatal(err)
	}

	pStr := string(p)

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, pStr)
	}

	server := httptest.NewServer(http.HandlerFunc(f))
	defer server.Close()
	oandaHost = server.URL


	pricing, err := GetPricing([]string{"EUR_USD"})

	if err != nil {
		t.Fail()
	}

	if pricing.Prices[0].Instrument != "EUR_USD" {
		t.Fail()
	}



}
