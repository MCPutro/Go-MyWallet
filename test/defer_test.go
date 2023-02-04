package test

import (
	"fmt"
	"testing"
)

func TestDefer1(t *testing.T) {
	nama := "aaaa"

	defer fmt.Println("nama : ", nama)
	nama = "bbb"
	fmt.Println("nama : ", nama)
}

func TestDefer2(t *testing.T) {
	nama := "aaaa"

	defer func() {
		fmt.Println("nama : ", nama)
	}()
	nama = "bbb"
	fmt.Println("nama : ", nama)
}
