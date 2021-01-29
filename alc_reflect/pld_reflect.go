package alc_reflect

import "reflect"

func IsPtr(source interface{}) bool {
	rv := reflect.ValueOf(source)
	return rv.Kind() == reflect.Ptr
}
func IsStruct(source interface{}) bool {
	rv := reflect.ValueOf(source)
	rv = rv.Elem()
	return rv.Kind() == reflect.Struct
}

func StructFieldIdxMap(t reflect.Type) map[string]int {
	rte := t.Elem()
	fieldMap := make(map[string]int)
	for i := 0; i < rte.NumField(); i++ {
		field := rte.Field(i)
		fieldMap[field.Name] = i
	}
	return fieldMap
}
