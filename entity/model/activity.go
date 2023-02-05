package model

import "time"

type Activity struct {
	ActivityId   uint32 //increment db
	UserId       string
	WalletIdFrom uint32
	WalletIdTo   uint32
	CategoryId   uint
	Period       string
	ActivityDate time.Time
	Nominal      uint32
	Desc         string
}
