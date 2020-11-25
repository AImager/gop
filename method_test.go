package gop

import (
	"fmt"
	"reflect"
	"testing"
)

type Aspect struct{}

func (a *Aspect) Before(point *JoinPoint) bool {
	fmt.Println("before")
	return true
}

func (a *Aspect) After(point *JoinPoint) {
	fmt.Println("after")
}

func (a *Aspect) Finally(point *JoinPoint) {
	fmt.Println("finally")
}

func (a *Aspect) GetAspectExpress() string {
	return ".*\\.HelloAop"
}

type Aspect1 struct{}

func (a *Aspect1) Before(point *JoinPoint) bool {
	fmt.Println("before1")
	return true
}

func (a *Aspect1) After(point *JoinPoint) {
	fmt.Println("after1")
}

func (a *Aspect1) Finally(point *JoinPoint) {
	fmt.Println("finally1")
}

func (a *Aspect1) GetAspectExpress() string {
	return ".*"
}

type HelloAop struct {
}

func (h *HelloAop) HelloAop() {
	fmt.Println("helloAop")
}

func TestRegisterMethodPoint(t *testing.T) {
	RegisterAspect([]AspectInterface{&Aspect1{}, &Aspect{}})
	RegisterMethodPoint(reflect.TypeOf((*HelloAop)(nil)))

	h := &HelloAop{}
	h.HelloAop()
}
