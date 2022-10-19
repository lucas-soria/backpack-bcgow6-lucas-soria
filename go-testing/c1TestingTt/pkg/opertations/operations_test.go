package opertations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	valuePair = [][2]int{{1, 2}, {2, 2}, {10, 20}}
	expected  = []int{3, 4, 30}
)

func TestAddGood(t *testing.T) {
	for i, pair := range valuePair {
		result := Add(pair[0], pair[1])
		if result != expected[i] {
			t.Errorf("Expected: '%d'. Got: '%d'", expected[i], result)
		}
	}

}

func TestAddAssertGood(t *testing.T) {
	for i, pair := range valuePair {
		result := Add(pair[0], pair[1])
		assert.Equal(t, expected[i], result)
	}
}
