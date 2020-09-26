package gop

import (
	"fmt"
	"reflect"
	"unsafe"

	"bou.ke/monkey"
)

type Aspect interface {
	Before(args []reflect.Value)
	After(args []reflect.Value)
	Around(fn reflect.Value, args []reflect.Value) []reflect.Value
}

func Annotation(fn interface{}, asp Aspect) {
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		panic("annotation fn not func")
	}
	// fnValue := reflect.ValueOf(fn)
	// fnValue.Set(v)
	// fnElem := fnValue.Elem()

	aop := &Aop{
		asp: asp,
	}

	// v := reflect.MakeFunc(reflect.TypeOf(original), original)
	// monkey.Patch(original, fn)

	fnPtr := unsafe.Pointer(reflect.ValueOf(fn).Pointer())
	fmt.Println(fnPtr)

	// newVal := reflect.NewAt(fnType, fnPtr)
	// aop.originFn = newVal
	// fmt.Println(unsafe.Pointer(newVal.Pointer()))
	// fnG := reflect.NewAt(fnType, fnPtr)

	// monkey.PatchWithDoubleValue(fnG.Elem(), reflect.ValueOf(fn))
	// newValue := reflect.NewAt(fnType, fnPtr)
	// fmt.Println(unsafe.Pointer(newValue.Pointer()))
	// aop.originFn = newValue

	// vc := reflect.MakeFunc(fnType, Test)

	// fnG := fnE.Convert(reflect.TypeOf(fn))
	// fmt.Println(unsafe.Pointer(fnG.Pointer()))
	// aop.originFn = fnG
	// monkey.GetPatchsOriginalBytes(reflect.ValueOf(fn).Pointer())
	// aop.originFn = reflect.NewAt(reflect.TypeOf(fn), fnPtr)

	v := reflect.MakeFunc(fnType, aop.Process)
	fmt.Println(unsafe.Pointer(v.Pointer()))
	fmt.Println(unsafe.Pointer(reflect.ValueOf(aop.Process).Pointer()))

	// fnValue.Set(v)
	monkey.PatchWithValue(fn, v)
	// aop.originFn = rep.GetReplacement()

	monkey.Patch(Test, fn)
	aop.originFn = reflect.ValueOf(Test)

	// a := monkey.GetPatch()
	fmt.Println("ok")
}

type Aop struct {
	asp      Aspect
	originFn reflect.Value
}

func (aop *Aop) Process(args []reflect.Value) []reflect.Value {
	fmt.Println("ok2")
	aop.asp.Before(args)
	defer aop.asp.After(args)

	fmt.Println("ok3")

	return aop.asp.Around(aop.originFn, args)
}

func Test() {
	fmt.Println("ok1")
}

// func (aop *Aop) Original(args []reflect.Value) []reflect.Value {
// 	return []reflect.Value{}
// }
