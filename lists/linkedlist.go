package lists

import (
	"reflect"
)


type linkedlist struct {
    InnerCollection ListNode;
}


func NewList(source interface{}) List {
	if nil == source {
		return newList( newSliceNode([]interface{}{}))
	}

	switch reflect.ValueOf(source).Kind() {
	case reflect.Slice:
		return newList( newSliceNode(source))
	}
	return newList(emptyNode())
}

func NewGeneratedList(initialValue interface{}, generator func(interface{}) interface{}) List {
	return newList( newGeneratorNode(initialValue, generator) )
}

func NewGeneratedListT(initialValue interface{}, generator interface{}) List {
	v := reflect.ValueOf(generator)
	foo := func(previous interface{}) interface{} {
		args := []reflect.Value {reflect.ValueOf(previous)}
		return v.Call(args)[0].Interface()
	}
	return NewGeneratedList(initialValue, foo)
}


func (list linkedlist) Done() bool {
	if nil == list.InnerCollection {
		return true
	}
	return list.InnerCollection.Done()
}

func (list linkedlist) Value() (bool, interface{}) {
	if list.Done() {
		return false, 0
	}
	return true, list.node().Value()
}

func (list linkedlist) Next() (bool,List) {
	if list.Done() {
		return false, newEmptyList()
	}
    return true, newList( list.node().Tail())
}

func (list linkedlist) node() ListNode {
	return list.InnerCollection
}