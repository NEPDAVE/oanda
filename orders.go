package oanda

import (
//"log"
//"strconv"
//"errors"
//"fmt"
)

//MarketBuyOrder builds struct needed for marshaling data into a []byte
func MarketBuyOrder(bid float64, ask float64, instrument string, units int) Orders {
	stopLossPrice := bid -
	takeProfitPrice := bid + (ask - bid) - .000002
	stopLoss := StopLossOnFill{TimeInForce: "GTC", Price: stopLossPrice}
	takeProfit := TakeProfitOnFill{TimeInForce: "GTC", Price: takeProfitPrice}
	timeInForce :=
	type :=
	positionFill :=
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
