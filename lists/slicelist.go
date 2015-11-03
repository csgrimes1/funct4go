package lists

import (
	"reflect"
)

func buildSlice(target []interface{}, source reflect.Value, sourceIndex int, sourceLength int) []interface{} {
	if sourceIndex >= sourceLength {
		return target
	}
	nextSlice := append(target, source.Index(sourceIndex).Interface())
	return buildSlice(nextSlice, source, sourceIndex + 1, sourceLength)
}

func makeReflectedSliceFrom(source reflect.Value) []interface{} {
	length := source.Len()
	beginSlice := make([]interface{}, 0, length)
	return buildSlice(beginSlice, source, 0, length)
}

func newSliceList(slice interface{}) ListNode {
	return sliceList{
		slice: makeReflectedSliceFrom(reflect.ValueOf(slice)),
	}
}

type sliceList struct {
	slice []interface{}
}

func (node sliceList) Done() bool {
	return len(node.slice) <= 0
}

func (node sliceList) Value() interface{} {
	return node.slice[0]
}

func (node sliceList) Tail() ListNode {
	return sliceList{
		slice: node.slice[1:],
	}
}
