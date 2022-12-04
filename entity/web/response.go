package web

type Response struct {
	Status  string
	Message interface{} `json:"Message,omitempty"`
	Data    interface{} `json:"Data,omitempty"`
}

type ResponseActivityType struct {
	Status  string
	Message interface{} `json:"Message,omitempty"`
	Income  interface{} `json:"Income,omitempty"`
	Expense interface{} `json:"Expense,omitempty"`
}
