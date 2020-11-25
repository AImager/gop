package gop

import (
	"fmt"
	"reflect"

	"bou.ke/monkey"
)

const methodLocation = "%s.%s.%s"

// register method point
func RegisterMethodPoint(pointType reflect.Type) {
	pkgPth := pointType.PkgPath()
	receiverName := pointType.Name()
	if pointType.Kind() == reflect.Ptr {
		pkgPth = pointType.Elem().PkgPath()
		receiverName = pointType.Elem().Name()
	}
	for i := 0; i < pointType.NumMethod(); i++ {
		method := pointType.Method(i)
		location := fmt.Sprintf(methodLocation, pkgPth, receiverName, method.Name)
		putMatchAspect(location)
		if len(getAspect(location)) < 1 {
			continue
		}
		var guard *monkey.PatchGuard
		var proxy = func(in []reflect.Value) []reflect.Value {
			guard.Unpatch()
			defer guard.Restore()
			receiver := in[0]
			point := newMethodJoinPoint(receiver, in[1:], method)
			defer finallyProcessed(point, location)
			if !beforeProcessed(point, location) {
				return point.Result
			}
			point.Result = receiver.MethodByName(method.Name).Call(in[1:])
			afterProcessed(point, location)
			return point.Result
		}
		// dynamic create proxy function
		proxyFn := reflect.MakeFunc(method.Func.Type(), proxy)
		// change original function call to proxy call
		guard = monkey.PatchInstanceMethod(pointType, method.Name, proxyFn.Interface())
	}
}

// new method join point
func newMethodJoinPoint(receiver interface{}, params []reflect.Value, method reflect.Method) *JoinPoint {
	point := &JoinPoint{
		Receiver: receiver,
		Params:   params,
		Method:   method,
		Func:     method.Func,
	}
	fn := method.Func
	fnType := fn.Type()
	nout := fnType.NumOut()
	point.Result = make([]reflect.Value, nout)
	for i := 0; i < nout; i++ {
		point.Result[i] = reflect.Zero(fnType.Out(i))
	}
	return point
}
