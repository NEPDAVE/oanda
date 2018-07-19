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
	Price            string           `json:"prices"`
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
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
}

//TakeProfitOnFill represents take profit parameters for an Order
type TakeProfitOnFill struct {
	TimeInForce string `json:"timeInForce"`
	Price       string `json:"price"`
}

//MarketBuyOrder builds struct needed for marshaling data into a []byte
func (o Order) MarketBuyOrder(bid float64, ask float64, instrument string, units int) Order {
	//tp/sl ratio is 3 to 1
	stopLossPrice := fmt.Sprintf("%.6f", bid-(ask*.005))
	takeProfitPrice := fmt.Sprintf("%.6f", ask+(ask*.015))
	stopLoss := StopLossOnFill{TimeInForce: "GTC", Price: stopLossPrice}
	takeProfit := TakeProfitOnFill{TimeInForce: "GTC", Price: takeProfitPrice}
	orderData := Order{
		StopLossOnFill:   stopLoss,
		TakeProfitOnFill: takeProfit,
		TimeInForce:      "FOK",
		Instrument:       instrument,
		Type:             "MARKET",
		PositionFill:     "DEFAULT"}
	order := Orders{Order: orderData}

	return order
}

//MarketSellOrder builds struct needed for marshaling data into a []byte
func MarketSellOrder(bid float64, ask float64, instrument string, units int) Order {
	//tp/sl ratio is 3 to 1
	stopLossPrice := fmt.Sprintf("%.6f", ask+(bid*.005))
	takeProfitPrice := fmt.Sprintf("%.6f", bid-(ask*.015))
	stopLoss := StopLossOnFill{TimeInForce: "GTC", Price: stopLossPrice}
	takeProfit := TakeProfitOnFill{TimeInForce: "GTC", Price: takeProfitPrice}
	orderData := Order{
		StopLossOnFill:   stopLoss,
		TakeProfitOnFill: takeProfit,
		TimeInForce:      "FOK",
		Instrument:       instrument,
		Type:             "MARKET",
		PositionFill:     "DEFAULT"}
	order := Orders{Order: orderData}

	return order
}

//MarshalOrders marshals order data into []byte for making API requests
func (o Orders) MarshalOrders(orders Orders) []byte {

	ordersByte, err := json.Marshal(orders)

	if err != nil {
		fmt.Println(err)
	}

	return ordersByte
}
