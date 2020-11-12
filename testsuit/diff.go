package testsuit

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/lib/types"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var dmp = diffmatchpatch.New()

// Diff diff interface{}
func Diff(o1, o2 interface{}, ignoreFields ...string) string {
	switch reflect.TypeOf(o1).Kind() {
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
	clearFields(o1, ignoreFields...)
	clearFields(o2, ignoreFields...)
	return diff(o1, o2)
}

// DiffSlice o1,o2 must be []ptr
func DiffSlice(o1, o2 interface{}, ignoreFields ...string) string {
	s1 := types.ToSlice(o1)
	clearSlice(s1, ignoreFields...)

	s2 := types.ToSlice(o2)
	clearSlice(s2, ignoreFields...)
	return diff(s1, s2)
}

// DiffMap m1,m2 must be map[interface{}]ptr
func DiffMap(o1, o2 interface{}, ignoreFields ...string) string {
	m1 := types.ToMap(o1)
	m2 := types.ToMap(o2)
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
	if o == nil {
		return
	}

	v := types.ElemValue(o)
	for _, name := range fields {
		field := v.FieldByName(name)
		if !field.IsValid() || types.IsNil(field) {
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
