package calculadora

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	valuePairFloats       = [][2]float64{{1.3, 2.1}, {2.5, 2.3}, {10.9, 20.1}, {3.5, 2.9}, {4.2, 1}}
	valuePairInts         = [][2]float64{{1, 2}, {2, 2}, {10, 20}, {3, 2}, {4.2, 1.2}}
	valuePairErrZero      = [][2]float64{{2, 0}, {5, 0}, {7, 0}, {22, 0}, {124, 0}, {57, 0}}
	valuePairMixedErrZero = [][2]float64{{2, 4}, {5, 0}, {7, 8}, {22, 0}, {124, 0}, {57, 0}}
)

func TestRestarGood(t *testing.T) {
	expected := []float64{-0.8, 0.2, -9.2, .6, 3.2}
	for i, pair := range valuePairFloats {
		result := Restar(pair[0], pair[1])
		assert.Equal(t, fmt.Sprintf("%.2f", expected[i]), fmt.Sprintf("%.2f", result))
	}
}

func TestRestarGoodType(t *testing.T) {
	expected := []int{-1, 0, -10, 1, 3}
	for i, pair := range valuePairInts {
		result := Restar(pair[0], pair[1])
		assert.EqualValues(t, expected[i], result)
	}
}

func TestRestarBad(t *testing.T) {
	expected := []float64{-19, 32, -.2, 1.6, 6.2}
	for i, pair := range valuePairFloats {
		result := Restar(pair[0], pair[1])
		assert.NotEqual(t, fmt.Sprintf("%.2f", expected[i]), fmt.Sprintf("%.2f", result))
	}
}

func TestRestarBadType(t *testing.T) {
	expected := []int{-2, 1, -11, 4, 9}
	for i, pair := range valuePairInts {
		result := Restar(pair[0], pair[1])
		assert.NotEqualValues(t, expected[i], result)
	}
}

func TestDividirFloatsGood(t *testing.T) {
	expected := []float64{0.62, 1.09, 0.54, 1.21, 4.2}
	for i, pair := range valuePairFloats {
		result, _ := Dividir(pair[0], pair[1])
		assert.Equal(t, fmt.Sprintf("%.2f", expected[i]), fmt.Sprintf("%.2f", result))
	}
}

func TestDividirZeroGood(t *testing.T) {
	for _, pair := range valuePairErrZero {
		_, err := Dividir(pair[0], pair[1])
		assert.Equal(t, ErrZeroDenominator, err)
	}
}

func TestDividirZeroMixedGood(t *testing.T) {
	expected := []float64{0.5, 0, 0.875, 0, 0, 0}
	for i, pair := range valuePairMixedErrZero {
		if result, err := Dividir(pair[0], pair[1]); err != nil {
			assert.Equal(t, ErrZeroDenominator, err)
		} else {
			assert.Equal(t, expected[i], result)
		}
	}
}
