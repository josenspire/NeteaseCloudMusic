package utils

import "reflect"

func TransformStruct2Map(st interface{}) map[string]interface{} {

	t := reflect.TypeOf(st)
	v := reflect.ValueOf(st)

	var params = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		params[t.Field(i).Name] = v.Field(i).String()
	}

	return params
}
