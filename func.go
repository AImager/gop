package gop

import (
	"fmt"
	"reflect"

	"bou.ke/monkey"
)

const funcLocation = "%s.%s"

// register func point
func RegisterFuncPoint(fn interface{}) {
	var guard *monkey.PatchGuard
	pointType := reflect.TypeOf(fn)
	pkgPth := pointType.PkgPath()
	location := fmt.Sprintf(funcLocation, pkgPth, pointType.Name())
	putMatchAspect(location)
	if len(getAspect(location)) < 1 {
		return
	}
	var proxy = func(in []reflect.Value) []reflect.Value {
		guard.Unpatch()
		defer guard.Restore()
		point := newFuncJoinPoint(reflect.ValueOf(fn), in)
		defer finallyProcessed(point, location)
		if !beforeProcessed(point, location) {
			return point.Result
		}
		point.Result = point.Func.Call(in)
		afterProcessed(point, location)
		return point.Result
	}
	// dynamic create proxy function
	proxyFn := reflect.MakeFunc(pointType, proxy)
	// change original function call to proxy call
	guard = monkey.Patch(fn, proxyFn.Interface())
}

// new func join point
func newFuncJoinPoint(fn reflect.Value, params []reflect.Value) *JoinPoint {
	point := &JoinPoint{
		Params: params,
		Func:   fn,
	}
	fnType := fn.Type()
	nout := fnType.NumOut()
	point.Result = make([]reflect.Value, nout)
	for i := 0; i < nout; i++ {
		point.Result[i] = reflect.Zero(fnType.Out(i))
	}
	return point
}
