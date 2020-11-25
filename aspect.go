package gop

import (
	"reflect"
)

// aspect interface
type AspectInterface interface {
	Before(point *JoinPoint) bool
	After(point *JoinPoint)
	Finally(point *JoinPoint)
	GetAspectExpress() string
}

// join point
type JoinPoint struct {
	Receiver interface{}
	Method   reflect.Method
	Params   []reflect.Value
	Result   []reflect.Value
	Func     reflect.Value
}

// aspect container
var aspectContainer = make([]AspectInterface, 0)

// append aspect to container
func RegisterAspect(aspectList []AspectInterface) {
	aspectContainer = append(aspectContainer, aspectList...)
}
