package types

import (
	"errors"
	"reflect"
)

// MergeMaps merge maps, the map key -> value must be have Merge method
func MergeMaps(out interface{}, ms ...interface{}) error {
	if reflect.TypeOf(out).Kind() != reflect.Map {
		return errors.New("arg0 must be map")
	}

	for _, m := range ms {
		if reflect.TypeOf(m).Kind() != reflect.Map {
			return errors.New("ms must be map")
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

	return nil
}
