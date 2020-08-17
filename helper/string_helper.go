package helper

import (
	"encoding/json"
	"reflect"
	"strings"
)

//将[id,id,id]字符串转换成id数组
func TransformStringToInt64Arr(idsString string) ([]int64, error) {
	resourceIdList := make([]int64, 0)
	dec := json.NewDecoder(strings.NewReader(idsString))
	dec.UseNumber()
	errJ := dec.Decode(&resourceIdList)
	return resourceIdList, errJ
}

//结构体转为map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
