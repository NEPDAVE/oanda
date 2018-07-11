package oanda

import (
	"encoding/json"
	"fmt"
)

/*
***************************
orders
***************************
*/

//Orders represents entire order(s) object when submiting and order to Oanda
type Orders struct {
	Order Order `json:"order"`
}

//Order represents single order to Oanda
type Order struct {
	Price            float64          `json:"prices"`
	StopLossOnFill   StopLossOnFill   `json:"stopLossOnFill"`
	TakeProfitOnFill TakeProfitOnFill `json:"takeProfitOnFill"`
	TimeInForce      string           `json:"timeInForce"`
	Instrument       string           `json:"instrument"`
	Units            int              `json:"units"`
	Type             string           `json:"type"`
	PositionFill     string           `json:"positionFill"`
}

//StopLossOnFill represents stop loss parameters for an Order
type StopLossOnFill struct {
	TimeInForce string  `json:"timeInForce"`
	Price       float64 `json:"price"`
}

//TakeProfitOnFill represents take profit parameters for an Order
type TakeProfitOnFill struct {
	TimeInForce string  `json:"timeInForce"`
	Price       float64 `json:"price"`
}

//MarshalOrders marshals order data into []byte for making API requests
func (o Orders) MarshalOrders(orders Orders) []byte {

	ordersByte, err := json.Marshal(orders)

	if err != nil {
		fmt.Println(err)
	}

	return ordersByte
}
