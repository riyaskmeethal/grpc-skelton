package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// pass key as "struct.structfield.fieldname" for nested struct.
func GetValueByField(source interface{}, key string) (valInterface interface{}, err error) {

	obj := reflect.ValueOf(source)
	if obj.Kind() == reflect.Ptr {
		obj = obj.Elem()
	}

	if obj.Kind() != reflect.Struct {
		err = fmt.Errorf("expecting struct")
		return
	}

	for _, fn := range strings.Split(key, ".") {
		obj = obj.FieldByName(fn)
		if !obj.IsValid() {
			err = fmt.Errorf("field doesn't exist")
			return
		}
	}
	valInterface = obj.Interface()
	return
}
