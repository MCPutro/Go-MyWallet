package model

type ActivityCategory struct {
	Type         string             `json:"Type,omitempty"`
	CategoryCode uint               `json:"CategoryCode,omitempty"`
	CategoryName string             `json:"CategoryName,omitempty"`
	Multiplier   int                `json:"Multiplier,omitempty"`
	SubCategory  []ActivityCategory `json:"SubCategory,omitempty"`
}
