package model

type ActivityCategory struct {
	CategoryCode uint               `json:"CategoryCode,omitempty"`
	CategoryName string             `json:"CategoryName,omitempty"`
	Multiplier   int32              `json:"Multiplier,omitempty"`
	SubCategory  []ActivityCategory `json:"SubCategory,omitempty"`
}
