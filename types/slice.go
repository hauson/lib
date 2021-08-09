package types

import (
	"encoding/json"
	"errors"
	"reflect"
	"sort"
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

// s is slice, convert strings and sort
func Strings(s interface{}) ([]string, error) {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		return nil, errors.New("args must slice")
	}

	var lines []string
	slice := reflect.ValueOf(s)
	for i := 0; i < slice.Len(); i++ {
		elem := slice.Index(i).Interface()
		bytes, err := json.Marshal(elem)
		if err != nil {
			return nil, err
		}

		lines = append(lines, string(bytes))
	}

	sort.Strings(lines)
	return lines, nil
}

// ConvSliceToMap,s must be slice, and  slice item must be Key() and Value() method,  m must be  map[Key()]Value() if args does not conform to specification, will panic
func ConvSliceToMap(s interface{}, m interface{}) {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		panic("args s must slice")
	}

	if reflect.TypeOf(m).Kind() != reflect.Map {
		panic("m must be map")
	}

	outMap := reflect.ValueOf(m)
	slice := reflect.ValueOf(s)
	for i := 0; i < slice.Len(); i++ {
		elem := slice.Index(i)

		key := elem.MethodByName("Key").Call(nil)[0]
		value := elem.MethodByName("Value").Call(nil)[0]
		outMap.SetMapIndex(key, value)
	}
}

// ConvSliceToLookupTable convert slice to look up table
func ConvSliceToLookupTable(s interface{}) map[interface{}]bool {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		panic("args s must slice")
	}

	lookupTable := make(map[interface{}]bool)
	slice := reflect.ValueOf(s)
	for i := 0; i < slice.Len(); i++ {
		elem := slice.Index(i).Interface()
		lookupTable[elem] = true
	}
	return lookupTable
}

// SumSlice sum Slice, s is slice, item have Key() and Value() value method, the value have Sum Method
func SumSlice(s interface{}, m interface{}) {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		panic("s must be slice")
	}

	if reflect.TypeOf(m).Kind() != reflect.Map {
		panic("m must be map")
	}

	outMap := reflect.ValueOf(m)
	slice := reflect.ValueOf(s)
	for i := 0; i < slice.Len(); i++ {
		elem := slice.Index(i)
		key := elem.MethodByName("Key").Call(nil)[0]
		value := elem.MethodByName("Value").Call(nil)[0]

		if kValue := outMap.MapIndex(key); kValue.IsValid() {
			newValue := sum(value, kValue)
			outMap.SetMapIndex(key, newValue)
		} else {
			outMap.SetMapIndex(key, value)
		}
	}
}

func sum(a1, a2 reflect.Value) reflect.Value {
	kind := a1.Kind()
	if kind != a2.Kind() {
		panic("a1 and s2 not same kind")
	}

	switch kind {
	case reflect.String:
		return reflect.ValueOf(a1.String() + a2.String())
	case reflect.Float32:
		return reflect.ValueOf(float32(a1.Float() + a2.Float()))
	case reflect.Float64:
		return reflect.ValueOf(a1.Float() + a2.Float())
	case reflect.Int:
		return reflect.ValueOf(int(a1.Int() + a2.Int()))
	case reflect.Int8:
		return reflect.ValueOf(int8(a1.Int() + a2.Int()))
	case reflect.Int16:
		return reflect.ValueOf(int16(a1.Int() + a2.Int()))
	case reflect.Int32:
		return reflect.ValueOf(int32(a1.Int() + a2.Int()))
	case reflect.Int64:
		return reflect.ValueOf(a1.Int() + a2.Int())
	case reflect.Uint:
		return reflect.ValueOf(uint(a1.Int() + a2.Int()))
	case reflect.Uint8:
		return reflect.ValueOf(uint8(a1.Int() + a2.Int()))
	case reflect.Uint16:
		return reflect.ValueOf(uint16(a1.Int() + a2.Int()))
	case reflect.Uint32:
		return reflect.ValueOf(uint32(a1.Int() + a2.Int()))
	case reflect.Uint64:
		return reflect.ValueOf(uint64(a1.Int() + a2.Int()))
	case reflect.Struct, reflect.Ptr:
		params := []reflect.Value{a2}
		return a1.MethodByName("Sum").Call(params)[0]
	default:
		panic("can not deal this kind " + kind.String())
	}
}
