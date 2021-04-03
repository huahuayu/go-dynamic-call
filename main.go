package main

import (
	"errors"
	"fmt"
	"github.com/huahuayu/go-dynamic-call/foo"
	"reflect"
)

func main() {
	funcName := "FuncA"
	args := make(map[string]interface{}, 0)
	args["arg0"] = "arg0"
	args["arg1"] = 0
	fakeDynamicCall(funcName, args)

	foo := new(foo.Foo)
	res := tureDynamicCall(foo, funcName, args)
	fmt.Println(res[0].String())
	fmt.Println(res[1].String())
	if err, ok := res[2].Interface().(error); ok && err != nil {
		fmt.Println(err)
	}
}

func fakeDynamicCall(fn string, args map[string]interface{}) (res0 string, res1 string, err error) {
	foo := new(foo.Foo)
	switch fn {
	case "FuncA":
		return foo.FuncA(args["arg0"].(string), args["arg1"].(int))
	case "FuncB":
		return foo.FuncB(args["arg0"].(string), args["arg1"].(int))
	case "FuncC":
		return foo.FuncC(args["arg0"].(string), args["arg1"].(int))
	default:
		return "", "", errors.New("no such function")
	}
}

func tureDynamicCall(obj interface{}, fn string, args map[string]interface{}) (res []reflect.Value) {
	method := reflect.ValueOf(obj).MethodByName(fn)
	var inputs []reflect.Value
	for _, v := range args {
		inputs = append(inputs, reflect.ValueOf(v))
	}
	return method.Call(inputs)
}
