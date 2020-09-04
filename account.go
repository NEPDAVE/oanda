package oanda

import (
	"encoding/json"
	"time"
)

type AccountPayload struct {
	Account           Account `json:"account"`
	LastTransactionID string  `json:"lastTransactionID"`
}
type Long struct {
	Pl           string `json:"pl"`
	ResettablePL string `json:"resettablePL"`
	Units        string `json:"units"`
	UnrealizedPL string `json:"unrealizedPL"`
}
type Short struct {
	Pl           string `json:"pl"`
	ResettablePL string `json:"resettablePL"`
	Units        string `json:"units"`
	UnrealizedPL string `json:"unrealizedPL"`
}
type Positions struct {
	Instrument   string `json:"instrument"`
	Long         Long   `json:"long"`
	Pl           string `json:"pl"`
	ResettablePL string `json:"resettablePL"`
	Short        Short  `json:"short"`
	UnrealizedPL string `json:"unrealizedPL"`
}
type Account struct {
	NAV                         string      `json:"NAV"`
	Alias                       string      `json:"alias"`
	Balance                     string      `json:"balance"`
	CreatedByUserID             int         `json:"createdByUserID"`
	CreatedTime                 time.Time   `json:"createdTime"`
	Currency                    string      `json:"currency"`
	HedgingEnabled              bool        `json:"hedgingEnabled"`
	ID                          string      `json:"id"`
	LastTransactionID           string      `json:"lastTransactionID"`
	MarginAvailable             string      `json:"marginAvailable"`
	MarginCloseoutMarginUsed    string      `json:"marginCloseoutMarginUsed"`
	MarginCloseoutNAV           string      `json:"marginCloseoutNAV"`
	MarginCloseoutPercent       string      `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue string      `json:"marginCloseoutPositionValue"`
	MarginCloseoutUnrealizedPL  string      `json:"marginCloseoutUnrealizedPL"`
	MarginRate                  string      `json:"marginRate"`
	MarginUsed                  string      `json:"marginUsed"`
	OpenPositionCount           int         `json:"openPositionCount"`
	OpenTradeCount              int         `json:"openTradeCount"`
	Orders                      []Order     `json:"orders"`
	PendingOrderCount           int         `json:"pendingOrderCount"`
	Pl                          string      `json:"pl"`
	PositionValue               string      `json:"positionValue"`
	Positions                   []Positions `json:"positions"`
	ResettablePL                string      `json:"resettablePL"`
	Trades                      []Trade     `json:"trades"`
	UnrealizedPL                string      `json:"unrealizedPL"`
	WithdrawalLimit             string      `json:"withdrawalLimit"`
}

func GetAccountSummary() (*AccountPayload, error) {
	reqArgs := &ReqArgs{
		ReqMethod: "GET",
		URL:       OandaHost + "/accounts/" + accountID + "/summary",
	}

	accountBytes, err := MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	accountPayload := &AccountPayload{}
	err = json.Unmarshal(accountBytes, accountPayload)

	if err != nil {
		return nil, err
	}

	return accountPayload, nil
}

func GetAccount() (*AccountPayload, error) {
	reqArgs := &ReqArgs{
		ReqMethod: "GET",
		URL:       OandaHost + "/accounts/" + accountID,
	}

	accountBytes, err := MakeRequest(reqArgs)

	if err != nil {
		return nil, err
	}

	accountPayload := &AccountPayload{}
	err = json.Unmarshal(accountBytes, accountPayload)

	if err != nil {
		return nil, err
	}

	return accountPayload, nil
}
