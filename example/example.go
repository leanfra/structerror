package main

import (
	"errors"
	"fmt"

	"github.com/leanfra/structerror"
)

var errFactory structerror.CodeErrorFactory

func main() {
	m := map[int]string{
		1: "error1",
		2: "error2",
	}

	errFactory = structerror.NewCodeErrorFactory(m)
	e := test()
	if ce, ok := e.(structerror.StructError); ok {
		j := ce.ErrorCode()
		fmt.Printf("%v\n", j)
	}

	if ce, ok := e.(structerror.StructJsonError); ok {
		j, _ := ce.JSON()
		fmt.Printf("%v\n", j)
	}

	ce1 := e.(structerror.StructError)
	e2 := test2()
	ce2 := e2.(structerror.StructError)
	fmt.Printf("%v\n", ce1.ErrorCode() == ce2.ErrorCode())
}

func test() error {
	error1 := errors.New("error old 1")
	error2 := errors.New("error old 2")
	return errFactory.Error(1, "found error1", error1, error2)
}

func test2() error {
	error1 := errors.New("error old 1")
	error2 := errors.New("error old 2")
	return errFactory.Error(1, "found error1 again", error1, error2)
}
