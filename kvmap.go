package kvmap

import (
	"fmt"
	"reflect"
)

// Tag is the tag name used as the key for the struct fields. This defaults to kv.
var Tag = "kv"

// ToKV converts the data to a map[string]interface{} using the key parameter as the
// base key.
func ToKV(key string, data interface{}) map[string]interface{} {

	switch reflect.TypeOf(data).Kind() {
	case reflect.Ptr:
		indData := reflect.Indirect(reflect.ValueOf(data)).Interface()
		return ToKV(key, indData)
	case reflect.Struct:
		return toStructKV(key, data)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return toSliceKV(key, data)
	case reflect.Map:
		return toMapKV(key, data)
	case reflect.Chan:
		return map[string]interface{}{key: nil}
	default:
		return map[string]interface{}{key: data}
	}

}

func toStructKV(key string, data interface{}) map[string]interface{} {
	v := reflect.Indirect(reflect.ValueOf(data))
	t := v.Type()

	result := map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {

		newKey := t.Field(i).Tag.Get(Tag)
		if newKey == "" {
			newKey = t.Field(i).Name
		}
		newKey = fmt.Sprintf("%s/%s", key, newKey)

		r := ToKV(newKey, v.Field(i).Interface())
		for k, v := range r {
			result[k] = v
		}

	}
	return result
}

func toSliceKV(key string, data interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	v := reflect.ValueOf(data)
	for i := 0; i < v.Len(); i++ {
		d := v.Index(i)
		newKey := fmt.Sprintf("%s/%d", key, i)

		r := ToKV(newKey, d.Interface())
		for k, v := range r {
			result[k] = v
		}
	}
	return result
}

func toMapKV(key string, data interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	v := reflect.ValueOf(data)

	keys := v.MapKeys()
	for _, k := range keys {
		newKey := fmt.Sprintf("%s/%v", key, k)
		val := v.MapIndex(k).Interface()
		r := ToKV(newKey, val)
		for kk, vv := range r {
			result[kk] = vv
		}
	}
	return result
}
