package urlparser

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// a must not be ptr
func URLEncode(a interface{}) string {
	var params []string
	rv := reflect.ValueOf(a)
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		v := rv.Field(i)
		f := rt.Field(i)
		tagValue := f.Tag.Get("json")
		switch tagValue {
		case "-":
			break
		case "":
			params = append(params, fmt.Sprintf("%v=%v", f.Name, v))
			break
		default:
			tags := strings.Split(tagValue, ",")
			tagName := tags[0]
			for _, item := range tags {
				if strings.HasPrefix(item, "default:") {
					defaults := strings.Split(item, ":")
					v = reflect.ValueOf(defaults[1])
					break
				}
			}

			params = append(params, fmt.Sprintf("%v=%v", tagName, v))
		}
	}

	if len(params) == 0 {
		return ""
	}

	return "?" + strings.Join(params, "&")
}

// a must be ptr
// "ID=7&name=xiaofing&start=2&limit=20"
func Decode(url string, a interface{}) error {
	kv := make(map[string]string)
	for _, str := range strings.Split(url, "&") {
		ss := strings.Split(str, "=")
		kv[ss[0]] = ss[1]
	}

	rv := reflect.ValueOf(a).Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		v := rv.Field(i)
		f := rt.Field(i)
		tagValue := f.Tag.Get("json")
		switch tagValue {
		case "-":
			break
		case "":
			if s, ok := kv[f.Name]; ok {
				x, err := convert(s, f.Type)
				if err != nil {
					return err
				}

				v.Set(x)
			}
		default:
			tags := strings.Split(tagValue, ",")
			tagName := tags[0]
			if _, ok := kv[tagName]; !ok {
				for _, item := range tags {
					if strings.HasPrefix(item, "default:") {
						defaults := strings.Split(item, ":")
						kv[tagName] = defaults[1]
						break
					}
				}
			}

			if s, ok := kv[tagName]; ok {
				x, err := convert(s, f.Type)
				if err != nil {
					return err
				}

				v.Set(x)
			}
		}
	}

	return nil
}

func convert(s string, typ reflect.Type) (reflect.Value, error) {
	switch typ.Kind() {
	case reflect.Bool:
		if strings.ToLower(s) == "false" {
			return reflect.ValueOf(false), nil
		} else {
			return reflect.ValueOf(true), nil
		}
	case reflect.Int:
		a, err := strconv.Atoi(s)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(a), nil
	case reflect.Int8:
		a, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(a), nil
	case reflect.Int16:
		a, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(a), nil
	case reflect.Int32:
		a, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(a), nil
	case reflect.Int64:
		a, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(a), nil
	case reflect.Uint:
		a, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(a), nil
	case reflect.Uint8:
		a, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(a), nil
	case reflect.Uint16:
		a, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(a), nil
	case reflect.Uint32:
		a, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(a), nil
	case reflect.Uint64:
		a, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(a), nil
	case reflect.Float32:
		a, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(a), nil
	case reflect.Float64:
		a, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(a), nil
	case reflect.String:
		return reflect.ValueOf(s), nil
	}

	return reflect.Value{}, errors.New("the kind no handler")
}

