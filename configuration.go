package oanda

import "os"

var (
	Host       = os.Getenv("OANDA_HOST")
	StreamHost = os.Getenv("STREAM_OANDA_HOST")
	AccountID  = os.Getenv("OANDA_ACCOUNT_ID")
	Bearer     = "Bearer " + os.Getenv("OANDA_TOKEN")
)

func Init() {
	Host = os.Getenv("OANDA_HOST")
	StreamHost = os.Getenv("STREAM_OANDA_URL")
	Bearer = "Bearer " + os.Getenv("OANDA_TOKEN")
	AccountID = os.Getenv("OANDA_ACCOUNT_ID")
}
