package web

type Response struct {
	Status  string
	Message interface{} `json:"Message,omitempty"`
	Data    interface{} `json:"Data,omitempty"`
}

type ResponseActivityType struct {
	Status   string
	Message  interface{}
	Income   interface{} `json:"IncomeCategory,omitempty"`
	Expense  interface{} `json:"ExpenseCategory,omitempty"`
	Transfer interface{} `json:"TransferCategory,omitempty"`
}
