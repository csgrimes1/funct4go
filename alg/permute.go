package alg
import (
	"github.com/csgrimes1/funct4go/lists"
	//"reflect"
)

type Yield struct {
	permuteResults lists.List
}

func (y Yield) RawResults() lists.List {
	return y.permuteResults
}

func Permute(ls ...lists.List) Yield {
	if len(ls) <= 0 {
		return Yield{ permuteResults: lists.CreateOptionalEmpty().List }
	}

	return Yield{permuteResults: permute(ls[0], ls[1:]) }
}


func buildSlice(listPrefix interface{}, appendThis interface{}) []interface{} {
	slice, ok := listPrefix.([]interface{})
	if ok {
		//Let's be defensive in case Go starts to over-allocate the arrays under
		//slices.
		newSlice := make([]interface{}, len(slice) +1, len(slice) + 1)
		copy(newSlice, append(slice, appendThis))
		return newSlice
	}

	//Start the slice
	return []interface{}{listPrefix, appendThis}
}

func permute(list1 lists.List, others []lists.List) lists.List {
	if len(others) <= 0 {
		return list1
	}
	next := list1.Map(func(item interface{})  interface{} {
		//Make a list of all combinations of this item with the next list
		popList := others[0]
		return popList.Map( func(item2 interface{}) interface{}{
			return buildSlice(item, item2)
		})
	}).
		FlatMap()

	return permute(next, others[1:])
}

func (y Yield) Map(mapFunc interface{}) lists.List {
	return y.permuteResults
}

func (y Yield) RawList() lists.List {
	return y.permuteResults
}

func (y Yield) Fold(accumulator interface{}) interface{} {
//	v := reflect.ValueOf(accumulator)
//	foo := func(value interface{}, accumulation interface{}) interface{} {
//		items := []reflect.Value {reflect.ValueOf(value), reflect.ValueOf(accumulation)}
//		return v.Call(items)[0].Interface()
//	}
//	return list.Fold(initialValue, foo)
	return 0
}