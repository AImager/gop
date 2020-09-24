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
	fn.Call(args)
	return []reflect.Value{}
}
