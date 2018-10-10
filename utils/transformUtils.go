package utils

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"reflect"
)

func TransformStructToMap(st interface{}) map[string]interface{} {
	t := reflect.TypeOf(st)
	v := reflect.ValueOf(st)

	var params = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		params[t.Field(i).Name] = v.Field(i).String()
	}
	return params
}

func TransformInterfaceToMap(origin interface{}) map[string]interface{} {
	var obj map[string]interface{}
	err := json.Unmarshal(origin.([]byte), &obj)
	if err != nil {
		return nil
	}
	return obj
}

func TransformStructToStr(model interface{}) string {
	if params, err := json.Marshal(model); err != nil {
		return err.Error()
	} else {
		return string(params[:])
	}
}

func TransformByteToJSON(str []byte) interface{} {
	var tsJson interface{}
	if err := json.Unmarshal(str, &tsJson); err != nil {
		beego.Error(err.Error())
		return nil
	} else {
		beego.Info(tsJson)
		return tsJson
	}
}
