package ordenamiento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	unsortedSlices = [][]int{
		{1, 5, 6, 3, 7, 8, 3, 8, 1, 1982, 7, 37, 2, 56, 567},
		{9, 4, 2, 6, 8, 0, 3, 1, 7, 5},
		{49, 20, 23, 34, 6, 29, 35},
	}
	sortedSlices = [][]int{
		{1, 1, 2, 3, 3, 5, 6, 7, 7, 8, 8, 37, 56, 567, 1982},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{6, 20, 23, 29, 34, 35, 49},
	}
)

func TestInsertionSortGood(t *testing.T) {
	var copySlice [][]int
	copy(unsortedSlices, copySlice)
	for i, slice := range copySlice {
		result := InsertionSort(slice)
		assert.Equal(t, sortedSlices[i], result)
	}
}

func TestInsertionAndBubbleSameGood(t *testing.T) {
	var copySlice [][]int
	copy(unsortedSlices, copySlice)
	for _, slice := range copySlice {
		assert.Equal(t, InsertionSort(slice), BubbleSort(slice))
	}
}

func TestInsertionAndSelectionSameGood(t *testing.T) {
	var copySlice [][]int
	copy(unsortedSlices, copySlice)
	for _, slice := range copySlice {
		assert.Equal(t, InsertionSort(slice), SelectionSort(slice))
	}
}

func TestInsertionSortBad(t *testing.T) {
	var copySlice [][]int
	copy(unsortedSlices, copySlice)
	for i, slice := range copySlice {
		result := InsertionSort(slice)
		assert.NotEqual(t, unsortedSlices[i], result)
	}
}
