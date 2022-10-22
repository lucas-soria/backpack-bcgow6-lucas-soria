package c3goTestingTt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 9223372036854775807
var testCases = map[string][]int64{
	"input":  {0, 1, 2, 3, 4, 7, 9, 23, 30},
	"output": {0, 1, 1, 2, 3, 13, 34, 28657, 832040},
}

func TestFibonacciFor(t *testing.T) {
	for i, j := range testCases["input"] {
		result, err := FibonacciFor(j)
		assert.Nil(t, err)
		assert.Equal(t, testCases["output"][i], result)
	}
}

func TestFibonacciForNegative(t *testing.T) {
	result, err := FibonacciFor(-1)
	assert.Equal(t, int64(0), result)
	assert.EqualError(t, err, NegativeNumber.Error())
}

func TestFibonacciRecursive(t *testing.T) {
	for i, j := range testCases["input"] {
		result, err := FibonacciRecursive(j)
		assert.Nil(t, err)
		assert.Equal(t, testCases["output"][i], result)
	}
}

func TestFibonacciNegative(t *testing.T) {
	result, err := FibonacciRecursive(-1)
	assert.Equal(t, int64(0), result)
	assert.EqualError(t, err, NegativeNumber.Error())
}

func BenchmarkFibonacciFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciFor(19)
	}
}
func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciRecursive(19)
	}
}
