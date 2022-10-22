package c3goTestingTt

import "errors"

var (
	NegativeNumber = errors.New("negative values are not allowed")
)

func FibonacciFor(loops int64) (result int64, err error) {
	if loops <= 1 {
		if loops > -1 {
			result = loops
			return
		}
		return 0, NegativeNumber
	}
	var previous int64 = 1
	var previousOfPrevious int64 = 0
	var i int64 = 1
	for ; i < loops; i++ {
		result = previous + previousOfPrevious
		previousOfPrevious = previous
		previous = result
	}
	return
}

func FibonacciRecursive(loops int64) (result int64, err error) {
	if loops <= 1 {
		if loops > -1 {
			result = loops
			return
		}
		err = NegativeNumber
		return
	}
	newResult1, _ := FibonacciRecursive(loops - 1)
	newResult2, _ := FibonacciRecursive(loops - 2)
	result = newResult1 + newResult2
	return
}
