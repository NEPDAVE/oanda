package oanda

import (
	"net/http"
	"os"
)

var (
	client   = &http.Client{}
	oandaURL = os.Getenv("OANDA_URL")
	//streamoandaURL = os.Getenv("STREAM_OANDA_URL")
	bearer    = "Bearer " + os.Getenv("OANDA_TOKEN")
	accountID = os.Getenv("OANDA_ACCOUNT_ID")
)
