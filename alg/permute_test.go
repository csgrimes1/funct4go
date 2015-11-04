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


func TestPermute1(t *testing.T) {
	l1 := generateList().Take(2)
	l2 := generateList().Take(3)
	l3 := generateList().Take(4)
	y := alg.Permute(
		l1,
		l2,
		l3,
	)

	fmt.Println(y.RawList())
	assert.Equal(t, 24, y.RawList().Count())
}
