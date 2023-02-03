package test

//func Test_User(t *testing.T) {
//	users := model.Users{
//		UserId:    "1",
//		Username:  "aku",
//		FirstName: "lu",
//		LastName:  "pa",
//	}
//
//	validate := validator.New()
//
//	err := validate.Struct(users)
//
//	if err != nil {
//		//fmt.Println("error :", err) //show message
//
//		//show more detail
//		validationErrors := err.(validator.ValidationErrors)
//		for _, fieldError := range validationErrors {
//			fmt.Println("error : ", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
//			fmt.Println(fieldError.Value())       //value dari variable yang di validasi
//			fmt.Println(fieldError.Param())       //nilai dari requirement
//			fmt.Println(fieldError.StructField()) //nama variable yang di validasi
//		}
//	}
//
//}
//
//func TestCheckIsEmail(t *testing.T) {
//	newUUID, err := uuid.NewUUID()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(newUUID.String())
//	fmt.Println(newUUID.String())
//	fmt.Println(newUUID.String())
//	fmt.Println(uuid.New().String())
//}
//
//func TestMap(t *testing.T) {
//	userMap := make(map[string]model.Users)
//
//	userMap["duwa"] = model.Users{
//		UserId:      "1",
//		Username:    "1",
//		FirstName:   "1",
//		LastName:    "1",
//		CreatedDate: time.Time{},
//		Status:      "1",
//		Authentication: model.UserAuthentication{
//			UserId:       "1",
//			Password:     "11",
//			Token:        "11",
//			RefreshToken: "11",
//		},
//		Data: nil,
//	}
//
//	users, ok1 := userMap["satu"]
//
//	fmt.Println("1", ok1)
//	fmt.Println("2", users)
//
//	users2, ok2 := userMap["duwa"]
//
//	fmt.Println("1", ok2)
//	fmt.Println("2", users2)
//}
//
//func TestStructLagi(t *testing.T) {
//
//	users := model.Users{
//		UserId:         "",
//		Username:       "",
//		FirstName:      "",
//		LastName:       "",
//		CreatedDate:    time.Time{},
//		Status:         "",
//		Authentication: model.UserAuthentication{},
//		Data:           nil,
//	}
//
//	fmt.Println(users.Data)
//
//}
//
//func TestFindAllUser(t *testing.T) {
//
//	db, err := app.InitDatabase()
//
//	if err != nil {
//		fmt.Println("error : ", err)
//		return
//	}
//
//	//load all
//	tx, err := db.Begin()
//
//	userRepository := repository.NewUserRepository()
//
//	findAll, err := userRepository.FindAll(context.Background(), tx)
//
//	fmt.Println(findAll)
//
//	//userByUsernameOrEmail, err := userRepository.FindByUsernameOrEmail(context.Background(), tx, "akulup2a3")
//	//
//	//fmt.Println(userByUsernameOrEmail)
//
//}
