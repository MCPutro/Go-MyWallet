package model

type Wallet struct {
	UserId   string `json:"UserId,omitempty" validate:"required"`
	WalletId uint   `json:"WalletId,omitempty"`
	Name     string `json:"Name,omitempty" validate:"required,max=25,min=3"`
	Type     string `json:"Type,omitempty" validate:"required,max=3,min=3"`
	IsActive string `json:"-"`
	Amount   uint32 `json:"Amount" validate:"required,numeric,gte=0"`
}
