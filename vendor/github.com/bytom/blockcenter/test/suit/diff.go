package suit

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var dmp = diffmatchpatch.New()

// Diff o1, o2 is can nil
func Diff(o1, o2 interface{}, ignoreFields ...string) string {
	if isNil(o1) && isNil(o2) {
		return ""
	}

	var kind reflect.Kind
	switch {
	case !isNil(o1): kind = reflect.TypeOf(o1).Kind()
	case !isNil(o2): kind =  reflect.TypeOf(o2).Kind()
	}

	switch kind {
	case reflect.Slice:
		return DiffSlice(o1, o2, ignoreFields...)
	case reflect.Map:
		return DiffMap(o1, o2, ignoreFields...)
	default:
		return DiffObj(o1, o2, ignoreFields...)
	}
}

// DiffObj o1, o2 must be ptr
func DiffObj(o1, o2 interface{}, ignoreFields ...string) string {
	if isNil(o1) || isNil(o2) {
		return diffText(o1, o2)
	}

	clearFields(o1, ignoreFields...)
	clearFields(o2, ignoreFields...)
	return diff(o1, o2)
}

// DiffSlice o1,o2 must be []ptr
func DiffSlice(o1, o2 interface{}, ignoreFields ...string) string {
	if isNil(o1) || isNil(o2) {
		return diffText(o1, o2)
	}

	s1 := toSlice(o1)
	clearSlice(s1, ignoreFields...)

	s2 := toSlice(o2)
	clearSlice(s2, ignoreFields...)
	return diff(s1, s2)
}

// DiffMap m1,m2 must be map[interface{}]ptr
func DiffMap(o1, o2 interface{}, ignoreFields ...string) string {
	if isNil(o1) || isNil(o2) {
		return diffText(o1, o2)
	}

	m1 := toMap(o1)
	m2 := toMap(o2)
	if len(m1) != len(m2) {
		return fmt.Sprintf(" m1 len %d, m2 len %d ", len(m1), len(m2))
	}

	var isDiff bool
	for k, v1 := range m1 {
		v2 := m2[k]
		clearFields(v1, ignoreFields...)
		clearFields(v2, ignoreFields...)
		if !reflect.DeepEqual(v1, v2) {
			isDiff = true
		}
	}

	if isDiff {
		diffs := dmp.DiffMain(mapString(m1), mapString(m2), false)
		return dmp.DiffPrettyText(diffs)
	}
	return ""
}

// o must be ptr
func clearFields(o interface{}, fields ...string) {
	if isNil(o) {
		return
	}

	v := elemValue(o)
	for _, name := range fields {
		field := v.FieldByName(name)
		if !field.IsValid() || isNil(field) {
			continue
		}

		field.Set(reflect.Zero(field.Type()))
	}
}

// ExistMap make key map, key must be key able
func ExistMap(ss ...interface{}) map[interface{}]bool {
	m := make(map[interface{}]bool)
	for _, s := range ss {
		m[s] = true
	}
	return m
}

func clearSlice(s []interface{}, fields ...string) {
	for _, o := range s {
		clearFields(o, fields...)
	}
}

func diff(o1, o2 interface{}) string {
	if reflect.DeepEqual(o1, o2) {
		return ""
	}

	return diffText(o1, o2)
}

func mapString(m map[interface{}]interface{}) string {
	o := make(map[string]interface{})
	for k, v := range m {
		bs, err := json.Marshal(k)
		if err != nil {
			return err.Error()
		}

		o[string(bs)] = v
	}

	bytes, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

func diffText(o1, o2 interface{}) string {
	bytes1, err := json.MarshalIndent(o1, "", "  ")
	if err != nil {
		return err.Error()
	}

	bytes2, err := json.MarshalIndent(o2, "", "  ")
	if err != nil {
		return err.Error()
	}

	diffs := dmp.DiffMain(string(bytes1), string(bytes2), false)
	return dmp.DiffPrettyText(diffs)
}

func toSlice(i interface{}) []interface{} {
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

func toMap(i interface{}) map[interface{}]interface{} {
	iType := reflect.TypeOf(i)
	if iType.Kind() != reflect.Map {
		return map[interface{}]interface{}{i: i}
	}

	iValue := reflect.ValueOf(i)
	m := make(map[interface{}]interface{})
	for _, k := range iValue.MapKeys() {
		m[k.Interface()] = iValue.MapIndex(k).Interface()
	}
	return m
}

func elemType(i interface{}) reflect.Type {
	iType := reflect.TypeOf(i)
	if iType.Kind() == reflect.Ptr {
		return iType.Elem()
	}

	return iType
}

func elemValue(i interface{}) reflect.Value {
	iValue := reflect.ValueOf(i)
	iType := reflect.TypeOf(i)
	if iType.Kind() == reflect.Ptr {
		return iValue.Elem()
	}

	return iValue
}

func isNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if !vi.IsValid() {
		return true
	}

	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}

	return false
}
