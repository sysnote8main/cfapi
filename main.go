package main

import (
	"fmt"
	"reflect"

	"github.com/sysnote8main/cfapi/model"
	"github.com/sysnote8main/cfapi/serializer"
)

func main() {
	d := model.SearchModQuery{
		GameId:  12,
		ClassId: 23,
	}
	serializer.ObjectToQuery(d)
	a(reflect.ValueOf(d))
}

func a(rv reflect.Value) {
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Type().Field(i)
		fmt.Printf("field name: %s, tag: %v\n", f.Name, f.Tag.Get("json"))
	}
}
