package test

import (
	"fmt"
	"testing"
)

func Test_len_key(t *testing.T) {
	s := []byte("KeY-<^>S2qUr1tY){.312wqD.^-^xeSW")

	fmt.Println(len(s))
	fmt.Println(s)
}
