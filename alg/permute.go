package alg
import (
	"github.com/csgrimes1/funct4go/lists"
	"reflect"
)

//A collector for permutations that allows yielding a new list or fold
//from the combinations.
type Yield struct {
	permuteResults lists.List
}

func (y Yield) RawResults() lists.List {
	return y.permuteResults
}

//Based on Scala for comprehension
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

func insert(first interface{}, tail []interface{}) []interface{} {
	buffer := make([]interface{}, 0, len(tail) + 1)
	return append(append(buffer, first), tail...)
}

//accumulator: func(accumulation, arg1, arg2, ... , argN)
func (y Yield) Fold(initialValue interface{}, accumulator interface{}) interface{} {
	accum := reflect.ValueOf(accumulator)
	foo := func(row interface{}, accumulation interface{}) interface{} {
		rowT := insert(accumulation, row.([]interface{}))
		rowValue := rawValueToSliceOfValues(rowT)
		return accum.Call(rowValue)[0].Interface()
	}
	return y.RawList().Fold(initialValue, foo)
}

func buildSlice2(row []interface{}, accum []reflect.Value, index int, capacity int) []reflect.Value {
	accum2 := append(accum, reflect.ValueOf(row[index]))
	if index >= capacity -1 {
		return accum2
	}
	return buildSlice2(row, accum2, index+1, capacity)
}

func rawValueToSliceOfValues(row []interface{}) []reflect.Value {
	targetLength := len(row)
	newSlice := make([]reflect.Value, 0, targetLength)
	return buildSlice2(row, newSlice, 0, targetLength)
}

func (y Yield) Emit(mapper interface{}) lists.List {
	lambda := reflect.ValueOf(mapper)
	mapper2 := func(row interface{}) interface{} {
		rowValue := rawValueToSliceOfValues(row.([]interface{}))
		return lambda.Call(rowValue)[0].Interface()
	}
	return y.RawList().Map(mapper2)
}

