package web

import "time"

type ActivityResponse struct {
	ActivityId   int       `json:"ActivityId"`
	Type         string    `json:"Type"`
	CategoryId   string    `json:"CategoryId"`
	WalletIdFrom int       `json:"WalletIdFrom"`
	WalletIdTo   int       `json:"WalletIdTo"`
	ActivityDate time.Time `json:"ActivityDate"`
	Amount       int       `json:"Amount"`
}
