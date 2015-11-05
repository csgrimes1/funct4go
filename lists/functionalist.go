package lists

import (
	"reflect"
	"fmt"
)


type functionallist struct {
	BasicList
}

func newList(node ListNode) List {
	return functionallist{
		BasicList: linkedlist{
			InnerCollection: node,
		},
	}
}

func newEmptyList() List {
	return functionallist{
		BasicList: linkedlist{
			InnerCollection: emptyNode(),
		},
	}
}


func (list functionallist) Map(mapper func(interface{}) interface{}) List {
	return newList(newMapNode(list.node(), mapper))
}

func (list functionallist) MapT(mapperFunc interface{}) List {
	v := reflect.ValueOf(mapperFunc)
	foo := func(item interface{}) interface{} {
		items := []reflect.Value {reflect.ValueOf(item)}
		return v.Call(items)[0].Interface()
	}
	return list.Map(foo)
}

func (list functionallist) Filter(predicate func(interface{}) bool) List {
	return newList(newFilterNode(list.node(), predicate))
}

func (list functionallist) FilterT(predicate interface{}) List {
	v := reflect.ValueOf(predicate)
	foo := func(item interface{}) bool {
		items := []reflect.Value {reflect.ValueOf(item)}
		return v.Call(items)[0].Bool()
	}
	return list.Filter(foo)
}


func (list functionallist) Count() int {
	return list.count(0)
}

func (list functionallist) count(countBefore int) int {
	return list.Fold(0, func(value interface{}, accumulation interface{}) interface{}{
		return accumulation.(int) + 1
	}).(int)
}

func (list functionallist) Take(length int) List {
	return newList( newLimiter(0, list.node(), func(index int, v interface{}) bool {
			return index < length
		}))
}

func (list functionallist) TakeWhile( predicate func(value interface{}) bool ) List {
	return newList( newLimiter(0, list.node(), func(index int, v interface{}) bool{
			return predicate(v)
		}))
}

func (list functionallist) TakeWhileT( predicate interface{} ) List {
	v := reflect.ValueOf(predicate)
	foo := func(item interface{}) bool {
		items := []reflect.Value {reflect.ValueOf(item)}
		return v.Call(items)[0].Bool()
	}
	return list.TakeWhile(foo)
}

func (list functionallist) skip(index int, number int) List {
	ok, next := list.Next()
	if !ok {
		return newList( emptyNode() )
	} else if index >= number {
		return list
	}
	return next.(functionallist).skip(index + 1, number)
}

func (list functionallist) Skip(number int) List {
	return list.skip(0, number)
}

func (list functionallist) Fold(initialValue interface{}, accumulator func(value interface{}, accumulation interface{}) interface{}) interface{} {
	var node List
	accum := initialValue
	for node = list; !node.Done(); _, node = node.Next() {
		_, val := node.Value()
		accum = accumulator(val, accum)
	}
	return accum
}

func (list functionallist) FoldT(initialValue interface{}, accumulator interface{}) interface{} {
	v := reflect.ValueOf(accumulator)
	foo := func(value interface{}, accumulation interface{}) interface{} {
		items := []reflect.Value {reflect.ValueOf(value), reflect.ValueOf(accumulation)}
		return v.Call(items)[0].Interface()
	}
	return list.Fold(initialValue, foo)
}

func (list functionallist) FlatMap() List {
	return newList( newChildNodes(list.node()))
}

func toString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func (list functionallist) String() string {
	return list.Fold("[", func (value interface{}, accumulation interface{}) interface{} {
		sAccum := accumulation.(string)
		if len(sAccum) <= 1 {
			return sAccum + toString(value)
		}
		return sAccum + "," + toString(value)
	}).(string) + "]"
}
