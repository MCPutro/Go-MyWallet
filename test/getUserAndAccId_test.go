package test

import (
	"fmt"
	"github.com/MCPutro/Go-MyWallet/helper"
	"testing"
)

func TestGetUserAndAccId(t *testing.T) {
	data := "1-2-3-4-5-6-7"
	id, s := helper.GetUserIdAndAccountId(data)
	fmt.Println(data)
	fmt.Println(id)
	fmt.Println(s)
}
