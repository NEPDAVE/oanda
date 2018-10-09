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

//ClientOrders represents entire Orders object for submiting/creating Orders
//with Oanda
type ClientOrders struct {
	Orders Orders `json:"order"`
}

//Orders represents single order to Oanda
type Orders struct {
	Price            string           `json:"price"`
	StopLossOnFill   StopLossOnFill   `json:"stopLossOnFill"`
	TakeProfitOnFill TakeProfitOnFill `json:"takeProfitOnFill"`
	TimeInForce      string           `json:"timeInForce"`
	Instrument       string           `json:"instrument"`
	Units            string           `json:"units"`
	Type             string           `json:"type"`
	PositionFill     string           `json:"positionFill"`
	TradeID          string           `json:"tradeID"`
	Distance         string           `json:"distance"`
	TriggerCondition string           `json:"triggerCondition"`
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

//MarshalClientOrders marshals order data into []byte for making API requests
func (c ClientOrders) MarshalClientOrders(clientOrders ClientOrders) []byte {

	clientOrdersByte, err := json.Marshal(clientOrders)

	if err != nil {
		log.Println(err)
	}

	return clientOrdersByte
}

//MarshalClosePositions marshals Close data into []byte for making API requests
func MarshalClosePositions(close Close) []byte {
	closePositionsByte, err := json.Marshal(close)

	if err != nil {
		log.Println(err)
	}

	return closePositionsByte
}
