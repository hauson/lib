package types

import "reflect"

func ElemType(i interface{}) reflect.Type {
	iType := reflect.TypeOf(i)
	if iType.Kind() == reflect.Ptr {
		return iType.Elem()
	}

	return iType
}

func IsTypeInitValue(i interface{}) bool {
	elemType := ElemType(i)
	v := reflect.New(elemType).Elem().Interface()
	return reflect.DeepEqual(i, v)
}

func FieldNames(i interface{}) []string {
	var columns []string
	elemType := ElemType(i)
	for j := 0; j < elemType.NumField(); j++ {
		columns = append(columns, elemType.Field(j).Name)
	}
	return columns
}

func FieldValues(i interface{}) []interface{} {
	var values []interface{}
	iValue := reflect.ValueOf(i)
	for i := 0; i < iValue.NumField(); i++ {
		values = append(values, iValue.Field(i).Interface())
	}
	return values
}

func ToSlice(i interface{}) []interface{} {
	iType := reflect.TypeOf(i)
	if iType.Kind() != reflect.Slice {
		return []interface{}{i}
	}

	iValue := reflect.ValueOf(i)
	num := iValue.Len()
	ss := make([]interface{}, num)
	for j := 0; j < num; j++ {
		ss[j] = iValue.Index(j).Interface()
	}
	return ss
}
