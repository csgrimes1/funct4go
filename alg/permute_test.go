package alg_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/csgrimes1/funct4go/lists"
	"github.com/csgrimes1/funct4go/alg"
	"fmt"
)

func generateList() lists.List {
	//Should yield 0..99
	return lists.NewGeneratedListT(-1, func(n int) int{
		return n + 1
	})
}

func makePermutation() alg.Yield {
	l1 := generateList().Take(2)
	l2 := generateList().Skip(1).Take(3)
	l3 := generateList().Skip(2).Take(4)
	return alg.Permute(
		l1,
		l2,
		l3,
	)
}


func TestPermute1(t *testing.T) {
	perm := makePermutation()

	//fmt.Println(y.RawList())
	assert.Equal(t, 24, perm.RawList().Count())
}


func TestFold(t *testing.T) {
	result := makePermutation().
		Fold(-1, func(accum int, a int, b int, c int) interface{}{
			sum := a + b + c
			if sum > accum {
				return sum
			}
			return accum
		})
	assert.Equal(t, 9, result)
}

func TestEmit(t *testing.T) {
	result := makePermutation().
		Emit(func(a int, b int, c int) interface{}{
			return fmt.Sprintf("(%v %v %v)", a, b, c)
		})
	//t.Logf("%v\n", result.String())
	_, v1 := result.Value()
	assert.Equal(t, "(0 1 2)", v1)
}

