package lists

import (
	"reflect"
	//"github.com/csgrimes1/funct4go/alg"
	"fmt"
)


type List struct {
    InnerCollection ListNode;
}


func NewList(source interface{}) List {
	switch reflect.ValueOf(source).Kind() {
	case reflect.Slice:
		return List{
			InnerCollection: newSliceList(source),
		}
	}
	return List{}
}

func NewGeneratedList(initialValue interface{}, generator func(interface{}) interface{}) List {
	return List {
		InnerCollection: newGeneratorList(initialValue, generator),
	}
}

func NewGeneratedListT(initialValue interface{}, generator interface{}) List {
	v := reflect.ValueOf(generator)
	foo := func(previous interface{}) interface{} {
		args := []reflect.Value {reflect.ValueOf(previous)}
		return v.Call(args)[0].Interface()
	}
	return NewGeneratedList(initialValue, foo)
}


func (list List) Done() bool {
	return list.InnerCollection.Done()
}

func (list List) Value() (bool, interface{}) {
	if list.Done() {
		return false, 0
	}
	return true, list.InnerCollection.Value()
}

func (list List) Next() (bool,List) {
	if list.Done() {
		return false, List{}
	}
    return true, List{
        InnerCollection: list.InnerCollection.Tail(),
    }
}

func (list List) Map(mapper func(interface{}) interface{}) List {
	return List{
		InnerCollection: newMapList(list.InnerCollection, mapper),
	}
}

func (list List) MapT(mapperFunc interface{}) List {
	v := reflect.ValueOf(mapperFunc)
	foo := func(item interface{}) interface{} {
		items := []reflect.Value {reflect.ValueOf(item)}
		return v.Call(items)[0].Interface()
	}
	return list.Map(foo)
}

func (list List) Filter(predicate func(interface{}) bool) List {
	return List{
		InnerCollection: newFilterList(list.InnerCollection, predicate),
	}
}

func (list List) FilterT(predicate interface{}) List {
	v := reflect.ValueOf(predicate)
	foo := func(item interface{}) bool {
		items := []reflect.Value {reflect.ValueOf(item)}
		return v.Call(items)[0].Bool()
	}
	return list.Filter(foo)
}


func (list List) Count() int {
	return list.count(0)
}

func (list List) count(countBefore int) int {
	return list.Fold(0, func(value interface{}, accumulation interface{}) interface{}{
		return accumulation.(int) + 1
	}).(int)
}

func (list List) Take(length int) List {
	return List{
		InnerCollection: newLimiter(0, list.InnerCollection, func(index int, v interface{}) bool {
			return index < length
		}),
	}
}

func (list List) TakeWhile( predicate func(value interface{}) bool ) List {
	return List {
		InnerCollection: newLimiter(0, list.InnerCollection, func(index int, v interface{}) bool{
			return predicate(v)
		}),
	}
}

func (list List) TakeWhileT( predicate interface{} ) List {
	v := reflect.ValueOf(predicate)
	foo := func(item interface{}) bool {
		items := []reflect.Value {reflect.ValueOf(item)}
		return v.Call(items)[0].Bool()
	}
	return list.TakeWhile(foo)
}

/*func (list List) Skip(number int) List {
	ok, next = list.Next()
	if !ok {
		return EmptyList()
	}

}*/

func (list List) Fold(initialValue interface{}, accumulator func(value interface{}, accumulation interface{}) interface{}) interface{} {
	var node List
	accum := initialValue
	for node = list; !node.Done(); _, node = node.Next() {
		_, val := node.Value()
		accum = accumulator(val, accum)
	}
	return accum
}

func (list List) FoldT(initialValue interface{}, accumulator interface{}) interface{} {
	v := reflect.ValueOf(accumulator)
	foo := func(value interface{}, accumulation interface{}) interface{} {
		items := []reflect.Value {reflect.ValueOf(value), reflect.ValueOf(accumulation)}
		return v.Call(items)[0].Interface()
	}
	return list.Fold(initialValue, foo)
}

func (list List) FlatMap() List {
	return List {
		InnerCollection: newChildNodes(list.InnerCollection),
	}
}

func toString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func (list List) String() string {
	return list.Fold("[", func (value interface{}, accumulation interface{}) interface{} {
		sAccum := accumulation.(string)
		if len(sAccum) <= 1 {
			return sAccum + toString(value)
		}
		return sAccum + "," + toString(value)
	}).(string) + "]"
}
