package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Interface2StringList(arr interface{}) []string {
	var s []string
	if arr == nil {
		return nil
	}
	v := reflect.ValueOf(arr)
	switch v.Kind() {
	case reflect.String:
		s = append(s, v.String())
	case reflect.Slice:
		l := v.Len()
		for i := 0; i < l; i++ {
			s = append(s, fmt.Sprint(v.Index(i).Interface()))
		}
	}
	return s
}

func Interface2InterfaceList(arr interface{}) []interface{} {
	var res []interface{}
	if arr == nil {
		return nil
	}
	v := reflect.ValueOf(arr)

	if v.Kind() != reflect.Slice {
		res = append(res, arr)
		return res
	}
	l := v.Len()
	// ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		// ret[i] = v.Index(i).Interface()
		res = append(res, v.Index(i).Interface())
	}
	return res
}

func Interface2Map(m interface{}) map[string]interface{} {

	if m == nil {
		return nil
	}
	v := reflect.ValueOf(m)

	if v.Kind() != reflect.Map {
		return nil
	}
	js, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	var jm map[string]interface{}

	err = json.Unmarshal(js, &jm)
	if err != nil {
		return nil
	}
	return jm
}

func Unmarshal(m interface{}, v interface{}) {
	data, _ := json.Marshal(m)
	json.Unmarshal(data, v)
}
