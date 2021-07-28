package batcher

import (
	"reflect"
)

// FillPage fill page by slice
func FillPage(page interface{}, all interface{}, offset, limit int) {
	if reflect.TypeOf(page).Kind() != reflect.Ptr {
		panic("arg1 must be slice ptr")
	}

	if reflect.TypeOf(all).Kind() != reflect.Slice {
		panic("arg2 must be slice")
	}

	allRV := reflect.ValueOf(all)
	if offset > allRV.Len() {
		return
	}

	end := offset + limit
	if end > allRV.Len() {
		end = allRV.Len()
	}

	elem := reflect.ValueOf(page).Elem()
	elem.Set(reflect.AppendSlice(elem, allRV.Slice(offset, end)))
}
