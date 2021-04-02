# go dynamic function call

foo.go

```go
import "fmt"

type Foo struct{}

func (foo Foo)FuncA(arg0 string, arg1 int)(res0 string, res1 string, err error){
	fmt.Println("FuncA been called")
	return "funcA_result0","funcA_result1",nil
}

func (foo Foo)FuncB(arg0 string, arg1 int)(res0 string, res1 string, err error){
	// different logic from FuncA
	fmt.Println("FuncB been called")
	return "funcB_result0","funcB_result1",nil
}

func (foo Foo)FuncC(arg0 string, arg1 int)(res0 string, res1 string, err error){
	// different logic from FuncA & FuncB
	fmt.Println("FuncC been called")
	return "funcC_result0","funcC_result1",nil
}
```

main.go

```go
package main

import (
	"fmt"
	"github.com/huahuayu/dynamic-call/foo"
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
```

How to replace the `switch...case statement` to more generic way for dynamic call function? Something like this

```go
func trueDynamicCall(fn string, args map[string]interface{}){
	foo := new(foo.Foo)
	foo.$fn(args["arg0"].(string),args["arg1"].(int))
}
```

Here is a good [example](https://medium.com/@vicky.kurniawan/go-call-a-function-from-string-name-30b41dcb9e12), but it's the function & function caller are in same package, what if there are in different package as my example?
