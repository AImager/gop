package gop

import (
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
	// monkey.GetPatchsOriginalBytes(reflect.ValueOf(fn).Pointer())
	aop.originFn = reflect.NewAt(reflect.TypeOf(fn), fnPtr)

	v := reflect.MakeFunc(fnType, aop.Process)

	// fnValue.Set(v)
	monkey.PatchWithValue(fn, v)
	// aop.originFn = rep.GetReplacement()

	// a := monkey.GetPatch()
	// fmt.Println(a)
}

type Aop struct {
	asp      Aspect
	originFn reflect.Value
}

func (aop *Aop) Process(args []reflect.Value) []reflect.Value {
	aop.asp.Before(args)
	defer aop.asp.After(args)

	return aop.asp.Around(aop.originFn, args)
}

// func (aop *Aop) Original(args []reflect.Value) []reflect.Value {
// 	return []reflect.Value{}
// }
