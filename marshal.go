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

type Orders struct {
	Order Order `json:"order"`
}

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

type StopLossOnFill struct {
	TimeInForce string  `json:"timeInForce"`
	Price       float64 `json:"price"`
}

type TakeProfitOnFill struct {
	TimeInForce string  `json:"timeInForce"`
	Price       float64 `json:"price"`
}

/*
FIXME here as an example
user := &User{Name: "Frank"}
    b, err := json.Marshal(user)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(b))
*/

func (o Orders) MarshalOrders(orders Orders) []byte {

	ordersByte, err := json.Marshal(orders)

	if err != nil {
		fmt.Println(err)
	}

	return ordersByte
}
