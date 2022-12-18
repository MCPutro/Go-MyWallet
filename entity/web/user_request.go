package web

type UserRegistration struct {
	Username string `validate:"alphanum,required,max=15,min=6"`
	FullName string `validate:"required,max=25,min=3"`
	Password string `validate:"required,max=100,min=3"`
	Email    string `validate:"required,email"`
	Imei     string
	DeviceId string
}

type UserRegistrationResp struct {
	UID      string
	Username string `validate:"alphanum,required,max=15,min=6"`
	Password string `validate:"required,max=100,min=3"`
}

type UserLogin struct {
	Account  string `validate:"alphanum,required,max=15,min=6"`
	Password string `validate:"required,max=100,min=3"`
}
