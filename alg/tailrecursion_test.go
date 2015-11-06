package alg_test

import (
	"github.com/csgrimes1/funct4go/alg"
	"testing"
	"github.com/stretchr/testify/assert"
)

func fibonacci(arg int, accum int) alg.RecurseResult {
	if accum > 21 {
		return alg.Done(accum)
	}

	return alg.Recurse(accum, arg + accum)
}

func TestTailRecursion(t *testing.T) {
	fib := alg.RunT(0, 1, fibonacci)

	assert.Equal(t, 34, fib)
}
