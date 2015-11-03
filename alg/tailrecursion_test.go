package alg_test

import (
	"github.com/csgrimes1/funct4go/alg"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func fibonacci(arg int, accum int) alg.RecurseResult {
	fmt.Printf("%v  %v", arg, accum)

	if accum > 21 {
		return alg.Done(accum)
	}

	return alg.Recurse(accum, arg + accum)
}

func TestTailRecursion(t *testing.T) {
	fib := alg.RunT(0, 1, fibonacci)

	assert.Equal(t, 34, fib)
}
