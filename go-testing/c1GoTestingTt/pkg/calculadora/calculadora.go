package calculadora

import "errors"

var (
	ErrZeroDenominator = errors.New("denominator cannot be zero")
)

func Restar(a, b float64) float64 {
	return a - b
}

func Dividir(a, b float64) (division float64, err error) {
	if b == 0 {
		err = ErrZeroDenominator
		return
	}
	division = a / b
	return
}
