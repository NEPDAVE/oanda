package oanda

import (
	"encoding/json"
	"log"
)

/*
***************************
orders
***************************
*/

//Orders represents entire order(s) object when creating an Oanda order
type Orders struct {
	Order Order `json:"order"`
}

//Order represents single order to Oanda
type Order struct {
	Price            string           `json:"prices"`
	StopLossOnFill   StopLossOnFill   `json:"stopLossOnFill"`
	TakeProfitOnFill TakeProfitOnFill `json:"takeProfitOnFill"`
	TimeInForce      string           `json:"timeInForce"`
	Instrument       string           `json:"instrument"`
	Units            string           `json:"units"`
	Type             string           `json:"type"`
	PositionFill     string           `json:"positionFill"`
}

//StopLossOnFill represents stop loss parameters for an Order
type StopLossOnFill struct {
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
}

//TakeProfitOnFill represents take profit parameters for an Order
type TakeProfitOnFill struct {
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
}

//MarshalOrders marshals order data into []byte for making API requests
func (o Orders) MarshalOrders(orders Orders) []byte {

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
