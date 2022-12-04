package web

type ActivityType struct {
	Code        uint           `json:"Code,omitempty"`
	Name        string         `json:"Name,omitempty"`
	SubCategory []ActivityType `json:"SubCategory,omitempty"`
}
