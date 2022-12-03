package web

type Response struct {
	Status  string
	Message interface{} `json:"Message,omitempty"`
	Data    interface{} `json:"Data,omitempty"`
}
