package model

type UserAuthentication struct {
	UserId       string `json:"UserId,omitempty"`
	Password     string `json:"Password,omitempty"`
	Token        string `json:"Token,omitempty"`
	RefreshToken string `json:"RefreshToken,omitempty"`
}
