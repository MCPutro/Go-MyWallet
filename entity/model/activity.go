package model

import "time"

type Activity struct {
	ActivityId uint8 //increment db
	UserId     string
	//Type         string //income, expense or transfer
	WalletIdFrom uint
	WalletIdTo   uint
	CategoryId   uint
	Period       string
	ActivityDate time.Time
	Amount       int32
}
