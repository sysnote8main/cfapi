package serializer

import (
	"fmt"
	"reflect"
)

func ObjectToQuery(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%+v / %v\n", t.Field(i), v.Field(i))
	}
}

func a(rv reflect.Value) {
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Type().Field(i)
		fmt.Printf("field name: %s, tag: %v\n", f.Name, f.Tag.Get("json"))
	}
}
