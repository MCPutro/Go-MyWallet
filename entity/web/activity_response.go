package web

import "time"

type ActivityResponse struct {
	ActivityId       uint8     `json:"ActivityId"`
	Type             string    `json:"Type"`
	Category         string    `json:"Category"`
	WalletIdFrom     uint      `json:"WalletIdFrom"`
	WalletIdTo       uint      `json:"WalletIdTo"`
	ActivityDate     time.Time `json:"ActivityDate"`
	Amount           uint32    `json:"Amount,omitempty"`
	AmountWalletFrom uint32    `json:"AmountWalletFrom,omitempty"`
	AmountWalletTo   uint32    `json:"AmountWalletTo,omitempty"`
}
