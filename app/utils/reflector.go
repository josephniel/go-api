package utils

import (
	"fmt"
	"reflect"
)

// CallFuncByName is a helper function that deals with dynamically calling a function by its name.
func CallFuncByName(myStruct interface{}, funcName string, params ...interface{}) (out []reflect.Value) {
	method := reflect.ValueOf(myStruct).MethodByName(funcName)
	if !method.IsValid() {
		panic(fmt.Errorf("Method not found \"%s\"", funcName))
	}
	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}
	return method.Call(in)
}
