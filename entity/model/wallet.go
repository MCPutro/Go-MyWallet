package model

type Wallet struct {
	UserId   string `json:"-" validate:"required"`
	WalletId uint32 `json:"WalletId,omitempty"`
	Name     string `json:"Name,omitempty" validate:"required,max=25,min=3"`
	Type     string `json:"Type,omitempty" validate:"required,max=3,min=3"`
	IsActive string `json:"-" default:"Y"`
	Amount   uint32 `json:"Amount" validate:"numeric,gte=0,lte=4294967295"` //4294967295
}
