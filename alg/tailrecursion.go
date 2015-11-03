package alg

import (
	"reflect"
)

type RecurseResult struct {
	arg interface{}
	returnValue interface{}
	done		bool
}

func Done(returnValue interface{}) RecurseResult {
	return RecurseResult{
		returnValue: returnValue,
		done: true,
	}
}

func Recurse(arg interface{}, accumulation interface{}) RecurseResult {
	return RecurseResult{
		arg: arg,
		returnValue: accumulation,
		done: false,
	}
}

func Run(initialArg interface{}, initialValue interface{}, callback func(arg interface{}, accumulation interface{}) RecurseResult) interface{} {
	var rr RecurseResult
	for rr = callback(initialArg, initialValue); !rr.done; rr = callback(rr.arg, rr.returnValue) {
	}
	return rr.returnValue
}

func RunT(initialArg interface{}, initialValue interface{}, callback interface{}) interface{} {
	v := reflect.ValueOf(callback)
	foo := func(arg interface{}, accumulation interface{}) RecurseResult {
		args := []reflect.Value {reflect.ValueOf(arg), reflect.ValueOf(accumulation)}
		return v.Call(args)[0].Interface().(RecurseResult)
	}
	return Run(initialArg, initialValue, foo)
}
