package web

import "time"

type NewActivityResponse struct {
	ActivityId         uint32    `json:"ActivityId"`
	Type               string    `json:"Type"`
	Category           string    `json:"Category"`
	WalletIdFrom       uint32    `json:"WalletIdFrom"`
	WalletIdTo         uint32    `json:"WalletIdTo"`
	ActivityDate       time.Time `json:"ActivityDate"`
	Nominal            uint32    `json:"Nominal,omitempty"`
	AmountWalletIdFrom uint32    `json:"AmountWalletIdFrom"`
	AmountWalletIdTo   uint32    `json:"AmountWalletIdTo"`
	Desc               string
}

type Activity struct {
	ActivityId     uint8     `json:"ActivityId"`
	Type           string    `json:"Type"`
	Category       string    `json:"Category"`
	WalletIdFrom   uint32    `json:"WalletIdFrom"`
	WalletNameFrom string    `json:"WalletNameFrom"`
	WalletIdTo     uint32    `json:"WalletIdTo"`
	WalletNameTo   string    `json:"WalletNameTo"`
	ActivityDate   time.Time `json:"ActivityDate"`
	Nominal        uint32    `json:"Nominal,omitempty"`
	Desc           string
}
