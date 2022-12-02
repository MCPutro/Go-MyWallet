package test

import (
	"fmt"
	"github.com/MCPutro/Go-MyWallet/entity/model"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"testing"
)

func Test_User(t *testing.T) {
	users := model.Users{
		UserId:    "1",
		Username:  "aku",
		FirstName: "lu",
		LastName:  "pa",
	}

	validate := validator.New()

	err := validate.Struct(users)

	if err != nil {
		//fmt.Println("error :", err) //show message

		//show more detail
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error : ", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
			fmt.Println(fieldError.Value())       //value dari variable yang di validasi
			fmt.Println(fieldError.Param())       //nilai dari requirement
			fmt.Println(fieldError.StructField()) //nama variable yang di validasi
		}
	}

}

func TestCheckIsEmail(t *testing.T) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newUUID.String())
	fmt.Println(newUUID.String())
	fmt.Println(newUUID.String())
	fmt.Println(uuid.New().String())
}
