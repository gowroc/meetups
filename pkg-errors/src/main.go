package main

import (
	"fmt"

	"github.com/mycodesmells/pkg-errors-example/bare"
	"github.com/mycodesmells/pkg-errors-example/concat"
	"github.com/mycodesmells/pkg-errors-example/wrap"
	"github.com/pkg/errors"
)

func main() {
	bareErr := bare.CallA()
	printErr("Bare", bareErr)

	concatErr := concat.CallA()
	printErr("Concat", concatErr)

	wrapErr := wrap.CallA()
	printErr("Wrap", wrapErr)
}

func printErr(name string, err error) {
	fmt.Printf("== %s ==\n", name)
	fmt.Printf("Message: %v\n", err)
	fmt.Printf("Type: %T\n", err)
	fmt.Printf("Original error? %v\n", errors.Cause(err))
	fmt.Printf("Original type? %T\n", errors.Cause(err))
	printStack(err)

	fmt.Println()
}

func printStack(err error) {
	if err, ok := err.(stackTracer); ok {
		for i, f := range err.StackTrace() {
			fmt.Printf("%+s:%d", f, i)
		}
		fmt.Printf("Stack: %+v", err.StackTrace())
	} else {
		fmt.Println("No stack trace...")
	}
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}
