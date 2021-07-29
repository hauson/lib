package types

import (
	"reflect"
)

// MergeMaps merge maps, the map key -> value must be have Merge method
func MergeMaps(out interface{}, ms ...interface{}) {
	if reflect.TypeOf(out).Kind() != reflect.Map {
		panic("arg0 must be map")
	}

	for _, m := range ms {
		if reflect.TypeOf(m).Kind() != reflect.Map {
			panic("ms must be map")
		}
	}

	outMap := reflect.ValueOf(out)
	for _, m := range ms {
		iterMap := reflect.ValueOf(m)
		for _, key := range iterMap.MapKeys() {
			kValue := iterMap.MapIndex(key)
			if totalValue := outMap.MapIndex(key); totalValue.IsValid() {
				params := []reflect.Value{kValue}
				mergeV := totalValue.MethodByName("Merge").Call(params)[0]
				outMap.SetMapIndex(key, mergeV)
			} else {
				outMap.SetMapIndex(key, kValue)
			}
		}
	}
}

// ConvMapToSlice convert map to slice, m must be map, s must be slice ptr
func ConvMapToSlice(m interface{}, s interface{}) {
	if reflect.TypeOf(m).Kind() != reflect.Map {
		panic("m must be map")
	}

	if reflect.TypeOf(s).Kind() != reflect.Ptr {
		panic("s must be ptr")
	}

	slice := reflect.ValueOf(s).Elem()
	if slice.Kind() != reflect.Slice {
		panic("s elem must be slice")
	}

	mv := reflect.ValueOf(m)
	for _, key := range mv.MapKeys() {
		value := mv.MapIndex(key)
		slice.Set(reflect.Append(slice, value))
	}
}
