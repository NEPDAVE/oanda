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
