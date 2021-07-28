package types

import (
	"errors"
	"reflect"
)

// SlicePluck slice elem one filedName
// s0 := []struct{
//     Name string
//     Gender int
//   }{...}
// var s1 []int
// SlicePluck(s0, "Gender", &s1)
func SlicePluck(s0 interface{}, fieldName string, s1 interface{}) error {
	slice0 := ElemValue(s0)
	if slice0.Kind() != reflect.Slice {
		return errors.New("s0 must be slice")
	}

	slice1 := ElemValue(s1)
	if slice1.Kind() != reflect.Slice {
		return errors.New("s1 must be slice")
	}

	num := slice0.Len()
	for j := 0; j < num; j++ {
		elem := slice0.Index(j)
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		field := elem.FieldByName(fieldName)
		slice1.Set(reflect.Append(slice1, field))
	}

	return nil
}
