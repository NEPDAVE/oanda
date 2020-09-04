package oanda

import (
	"net/http"
	"os"
)

var (
	client         = &http.Client{}
	OandaHost      = os.Getenv("OANDA_HOST")
	streamOandaURL = os.Getenv("STREAM_OANDA_HOST")
	bearer         = "Bearer " + os.Getenv("OANDA_TOKEN")
	accountID      = os.Getenv("OANDA_ACCOUNT_ID")
)

func Init() {
	OandaHost = os.Getenv("OANDA_HOST")
	streamOandaURL = os.Getenv("STREAM_OANDA_URL")
	bearer = "Bearer " + os.Getenv("OANDA_TOKEN")
	accountID = os.Getenv("OANDA_ACCOUNT_ID")

}
