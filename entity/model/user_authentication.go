package model

type UserAuthentication struct {
	UserId       string `json:"-"`
	Password     string `json:"-"`
	Token        string
	RefreshToken string
}
