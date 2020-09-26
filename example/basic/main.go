package main

import (
	"fmt"
	"reflect"

	"github.com/aimager/gop"
)

func main() {
	c := &Cong{}
	gop.Annotation(test, c)
	// monkey.Patch(test, func() { fmt.Println("test1") })
	test()
}

func test() {
	fmt.Println("test")
}

type Cong struct{}

func (c *Cong) Before(args []reflect.Value) {
	fmt.Println("before")
}
func (c *Cong) After(args []reflect.Value) {
	fmt.Println("after")
}
func (c *Cong) Around(fn reflect.Value, args []reflect.Value) []reflect.Value {
	fmt.Println("around")
	// fnElem := fn.Elem()
	// fmt.Println(unsafe.Pointer(fn.Pointer()))
	// fnE := fn.Elem()
	// fmt.Println(unsafe.Pointer(fnE.Pointer()))
	// fnC := fnE.Convert(reflect.TypeOf(test))
	// // fnC.Call(args)
	// fnC.Call(args)
	// fmt.Println(unsafe.Pointer(fn.Pointer()))
	// fnElem.Call(args)
	fn.Call(args)
	return []reflect.Value{}
}
