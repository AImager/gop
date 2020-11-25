package gop

import (
	"fmt"
	"testing"
)

func test() {
	fmt.Println("nihao test")
}

func TestRegisterFuncPoint(t *testing.T) {
	RegisterAspect([]AspectInterface{&Aspect1{}, &Aspect{}})
	RegisterFuncPoint(test)

	test()
}
