package oanda

import (
	"encoding/json"
	"log"
)

//MarshalOrders marshals order data into []byte for making API requests
func MarshalOrders(orders Orders) []byte {

	ordersByte, err := json.Marshal(orders)

	if err != nil {
		log.Println(err)
	}

	return ordersByte
}

//MarshalClosePositions marshals Close data into []byte for making API requests
func MarshalClosePositions(close Close) []byte {
	closePositionsByte, err := json.Marshal(close)

	if err != nil {
		log.Println(err)
	}

	return closePositionsByte
}
