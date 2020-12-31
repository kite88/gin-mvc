package test

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println("fun name:", t.Name())
}
