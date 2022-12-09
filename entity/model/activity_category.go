package model

type ActivityCategory struct {
	CategoryId      uint32
	Type            string
	CategoryName    string
	SubCategoryName string
	IsActive        string `json:"IsActive,omitempty"`
	//Type         string             `json:"Type,omitempty"`
	//CategoryCode uint               `json:"CategoryCode,omitempty"`
	//CategoryName string             `json:"CategoryName,omitempty"`
	//Multiplier   int                `json:"Multiplier,omitempty"`
	//SubCategory  []ActivityCategory `json:"SubCategory,omitempty"`
}
