package foo

import "fmt"

type Foo struct{}

func (foo Foo) FuncA(arg0 string, arg1 int) (res0 string, res1 string, err error) {
	fmt.Println("FuncA been called")
	return "funcA_result0", "funcA_result1", nil
}

func (foo Foo) FuncB(arg0 string, arg1 int) (res0 string, res1 string, err error) {
	// different logic from FuncA
	fmt.Println("FuncB been called")
	return "funcB_result0", "funcB_result1", nil
}

func (foo Foo) FuncC(arg0 string, arg1 int) (res0 string, res1 string, err error) {
	// different logic from FuncA & FuncB
	fmt.Println("FuncC been called")
	return "funcC_result0", "funcC_result1", nil
}
