package test

import (
	"context"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/app"
	"github.com/MCPutro/Go-MyWallet/repository"
	"github.com/MCPutro/Go-MyWallet/service"
	"testing"
)

func TestGetActType(t *testing.T) {

	db, err := app.InitDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	begin, err := db.Begin()

	activityRepository := repository.NewActivityRepository()

	_, err = activityRepository.GetActivityTypes(context.Background(), begin)

	if err != nil {
		begin.Rollback()
	}
	begin.Commit()
}

func TestKedua(t *testing.T) {
	db, err := app.InitDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	activityRepository := repository.NewActivityRepository()
	activityService := service.NewActivityService(activityRepository, db)

	activityService.GetActivityType(context.Background())

}

func TestK(t *testing.T) {
	var k string = "a"
	var K string = "A"

	fmt.Println(k, K, k == K)
}
