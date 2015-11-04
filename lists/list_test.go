package lists_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/csgrimes1/funct4go/lists"
)

func makeBaseList() lists.List {
	return lists.NewList([]int{0, 1, 2})
}

func makeMappedList() lists.List{
	return makeBaseList().
		MapT( func(n int) int {
			return n+10
		});
}

func generateList() lists.List {
	//Should yield 0..99
	return lists.NewGeneratedListT(-1, func(n int) int{
		return n + 1
	})
}

func TestMapping(t *testing.T) {
	result := makeMappedList()
	initialCount := result.Count()

	for n := 0; n<3; n++ {
		ok, value := result.Value()
		assert.Equal(t, true, ok)
		assert.Equal(t, n+10, value)
		_, result = result.Next()
	}

	stillOk, _ := result.Value()
	assert.Equal(t, false, stillOk)
	assert.Equal(t, 3, initialCount)
}

func TestFiltering(t *testing.T) {
	const length = 100
	filteredList := generateList().
		Take(length).
		FilterT( func(n int) bool {
			return n % 2 == 0
		}).Take(length * 2000)

	_, val := filteredList.Value()
	assert.Equal(t, 0, val)
	assert.Equal(t, length/2, filteredList.Count())
}

func TestTakeWhile(t *testing.T) {
	const length = 150
	list := generateList().TakeWhileT( func(val int) bool {
		return val < length
	} )
	assert.Equal(t, length, list.Count())
}

func TestFlatMap(t *testing.T) {
	col := []lists.List{
		generateList().Take(1),
		generateList().Take(0),
		generateList().Take(10),
	}
	outerList := lists.NewList(col)
	fm := outerList.FlatMap()
	assert.Equal(t, 11, fm.Count())
}

func TestFolds(t *testing.T) {
	const length = 1000000
	result := generateList().
		Take(length).
		FoldT(-1, func(val int, accumulation int) int {
			if val % 100 == 0  &&  val > accumulation {
				return val
			}
			return accumulation
		})
	assert.Equal(t, length - 100, result)
}

func TestList2String(t *testing.T) {
	s := generateList().
		Take(10).
		String()
	assert.Equal(t, "[0,1,2,3,4,5,6,7,8,9]", s)
}
