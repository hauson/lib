package types

import (
	"testing"
	"fmt"
)

type Class struct {
	Name string
	P    *Student
}

type Student struct{}

func TestFieldValues(t *testing.T) {
	c := &Class{Name: "num one class"}
	values := FieldValues(c)
	for _, v := range values {
		fmt.Println(v)
	}
}

func TestIsTypeInitValue(t *testing.T) {
	var P *Class

	result := IsTypeInitValue(P)
	fmt.Println("IsTypeInitValue:", result)
}
