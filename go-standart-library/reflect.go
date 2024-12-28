package main

import (
	"fmt"
	"reflect"
)

func main() {
	sampleRequest := User{Name: "", Age: 20}
	isValidUser := isValid(sampleRequest)
	fmt.Println(isValidUser)
}

type User struct {
	Name string `required:"true"`
	Age  int
}

func isValid(value interface{}) bool {
	t := reflect.TypeOf(value)
	numberOfFields := t.NumField()

	for i := 0; i < numberOfFields; i++ {
		f := t.Field(i)
		if (f.Tag.Get("required") == "true") {
			data := reflect.ValueOf(value).Field(i).Interface()
			return data != ""
		}
	}
	
	return true
}