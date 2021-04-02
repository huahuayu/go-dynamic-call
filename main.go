package main

import (
	"fmt"
	"github.com/huahuayu/go-dynamic-call/foo"
)

func main(){
	funcName := "FuncA"
	args := make(map[string]interface{},0)
	args["arg0"] = "arg0"
	args["arg1"] = 0
	fakeDynamicCall(funcName,args)
}

// the case statements are similar, which repeated myself
func fakeDynamicCall(fn string, args map[string]interface{}){
	foo := new(foo.Foo)
	switch fn {
	case "FuncA":
		foo.FuncA(args["arg0"].(string),args["arg1"].(int))
	case "FuncB":
		foo.FuncB(args["arg0"].(string),args["arg1"].(int))
	case "FuncC":
		foo.FuncC(args["arg0"].(string),args["arg1"].(int))
	default:
		fmt.Println("no such function")
	}
}

/** can I have a more generic dynamic call like below?

func trueDynamicCall(fn string, args map[string]interface{}){
	foo := new(foo.Foo)
	foo.$fn(args["arg0"].(string),args["arg1"].(int))
}

 */



