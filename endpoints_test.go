package oanda

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

//TestGetPricing validates GetPricing returns pricesByte and nil error
func TestGetPricing(t *testing.T) {
	_, err := GetPricing("EUR_USD", "USD_JPY", "GBP_USD", "USD_CHF")

	t.Log("Given the need to test Oanda API endpoints")
	t.Log("\tWhen checking GetPricing function")

	if err != nil {
		t.Fatal("\t\tShould return nil error", ballotX, err)
	}
	t.Log("\t\tShould return nil error", checkMark)
}

//TestGetCandles validates GetPricing returns pricesByte and nil error
func TestGetCandles(t *testing.T) {
	_, err := GetPricing("EUR_USD", "50", "D")

	t.Log("Given the need to test Oanda API endpoints")
	t.Log("\tWhen checking GetCandles function")

	if err != nil {
		t.Fatal("\t\tShould return nil error", ballotX, err)
	}
	t.Log("\t\tShould return nil error", checkMark)
}
