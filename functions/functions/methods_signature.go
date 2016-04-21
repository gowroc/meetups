package main

import (
	"fmt"
	"reflect"
)

// START OMIT

type X struct{}

func (x X) ValueMethod(a int, b string) (float64, error) {
	return 0.0, nil
}

func (x *X) PointerMethod(a int, b string) (float64, error) {
	return 0.0, nil
}

// END OMIT

func methods(v interface{}) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		fmt.Printf("&%v\n", t.Elem())
	} else {
		fmt.Printf("%v.%v\n", t.PkgPath(), t.Name())
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("\t%v %v\n", m.Name, m.Type)
	}
}

func main() {
	methods(X{})
	methods(&X{})
}
