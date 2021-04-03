# go dynamic function call

call a go function by string function name

[foo.go](https://github.com/huahuayu/go-dynamic-call/blob/master/foo/foo.go) have three similar function.

```go
import "fmt"

type Foo struct{}

func (foo Foo)FuncA(arg0 string, arg1 int)(res0 string, res1 string, err error){
	fmt.Println("FuncA been called")
	return "funcA_result0","funcA_result1",nil
}

func (foo Foo)FuncB(arg0 string, arg1 int)(res0 string, res1 string, err error){
	fmt.Println("FuncB been called")
	return "funcB_result0","funcB_result1",nil
}

func (foo Foo)FuncC(arg0 string, arg1 int)(res0 string, res1 string, err error){
	fmt.Println("FuncC been called")
	return "funcC_result0","funcC_result1",nil
}
```


[fakeDynamicCall](https://github.com/huahuayu/go-dynamic-call/blob/1b28c655e79218b66c4db1d391c717aaae9e103f/main.go#L26) have some extent flexible, but it's based on `switch case`

```go
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

while the [trueDynamicCall](https://github.com/huahuayu/go-dynamic-call/blob/1b28c655e79218b66c4db1d391c717aaae9e103f/main.go#L40) is base on `reflect`

```go
func tureDynamicCall(obj interface{}, fn string, args map[string]interface{}) (res []reflect.Value){
    method := reflect.ValueOf(obj).MethodByName(fn)
    var inputs []reflect.Value
    for _, v := range args {
        inputs = append(inputs, reflect.ValueOf(v))
    }
    return method.Call(inputs)
}
```

## reference

more [example](https://medium.com/@vicky.kurniawan/go-call-a-function-from-string-name-30b41dcb9e12)

thanks [@miguelmota](https://github.com/miguelmota) for the `reflect` solution!

