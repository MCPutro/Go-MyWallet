package test

//func TestGetActType(t *testing.T) {
//
//	db, err := app.InitDatabase()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	begin, err := db.Begin()
//
//	activityRepository := repository.NewActivityc()
//
//	s, err := activityRepository.FindActivityCategory(context.Background(), begin)
//
//	for _, category := range s {
//		fmt.Println(category)
//	}
//
//	if err != nil {
//		begin.Rollback()
//	}
//	begin.Commit()
//}
//
//func TestKedua(t *testing.T) {
//	db, err := app.InitDatabase()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	activityRepository := repository.NewActivityRepository()
//	activityService := service.NewActivityService(activityRepository, repository.NewWalletRepository(), db)
//
//	activityService.GetActivityCategory(context.Background())
//
//}
//
//func TestGetActCategoryByIdService(t *testing.T) {
//	db, err := app.InitDatabase()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	walletRepository := repository.NewWalletRepository()
//	activityRepository := repository.NewActivityRepository()
//	activityService := service.NewActivityService(activityRepository, walletRepository, db)
//
//	responseActivityType, err := activityService.GetActivityCategoryById(context.Background(), 23)
//
//	fmt.Println(err)
//	fmt.Println(responseActivityType)
//
//}
//
//func BenchmarkAppendStruct(b *testing.B) {
//	var a []model.ActivityCategory
//
//	for i := 0; i < b.N; i++ {
//		a = append(a, model.ActivityCategory{})
//	}
//
//}
//
//func BenchmarkAppendPointer(b *testing.B) {
//	var a []*model.ActivityCategory
//
//	for i := 0; i < b.N; i++ {
//		a = append(a, &model.ActivityCategory{})
//	}
//
//}
//
//func TestGetActCategoryByIdRepository(t *testing.T) {
//	ctx := context.Background()
//
//	db, err := app.InitDatabase()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	tx, err := db.Begin()
//
//	//walletRepository := repository.NewWalletRepository()
//	activityRepository := repository.NewActivityRepository()
//	//activityService := service.NewActivityService(activityRepository, walletRepository, db)
//
//	categories, err := activityRepository.FindActivityCategoryById(ctx, tx, 5)
//
//	fmt.Println(categories)
//
//	//for i, category := range categories {
//	//	fmt.Println(i, category)
//	//}
//
//	if err != nil {
//		tx.Rollback()
//	}
//	tx.Commit()
//}
