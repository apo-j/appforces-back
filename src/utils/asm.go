//array slice map
package utils

import (
	"reflect"
	_"models"
	_ "log"
	"strconv"
)

func GroupById(self interface{}) map[string][]interface{}{
	res := make(map[string][]interface {})
	slice := reflect.ValueOf(self)

	for i:=0; i<slice.Len(); i++ {
		key := strconv.FormatInt(slice.Index(i).FieldByName("ID").Int(), 10)
		res[key] = append(res[key], slice.Index(i).Interface())
	}
	return res
}

func GroupByString(self interface{}, property string) map[string][]interface{}{
	res := make(map[string][]interface {})
	slice := reflect.ValueOf(self)

	for i:=0; i<slice.Len(); i++ {
		key := reflect.ValueOf(slice.Index(i)).FieldByName(property).Elem().String()
		res[key] = append(res[key], slice.Index(i).Interface())
	}
	return res
}

/*func GroupById(self interface{}) map[string][]interface{}{
	res := make(map[string][]interface {})
	slice := reflect.ValueOf(self)

	for i:=0; i<slice.Len(); i++ {
		key := strconv.FormatInt(slice.Index(i).FieldByName("ID").Int(), 10)

		if items, ok := res[key]; ok{
			res[key] = append(items, slice.Index(i).Interface())
		}else{
			res[key] = append(items, slice.Index(i).Interface())
		}
	}
	return res
}*/
