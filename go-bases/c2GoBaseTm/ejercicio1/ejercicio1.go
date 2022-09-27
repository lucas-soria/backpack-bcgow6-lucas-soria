package ejercicio1

import (
	"errors"
	"fmt"
)

/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000
se le descontará además un 10%.
*/

func tax(salary int) (percentage float64, err error) {
	if salary < 0 {
		return 0, errors.New("salario negativo")
	}
	if salary > 50_000 {
		percentage = 0.17
	}
	if salary > 150_000 {
		percentage *= 0.10
	}
	return float64(salary) * percentage, nil
}

func TaxCalculator() {
	if tax, err := tax(2_048_284); err == nil {
		fmt.Println("El impuesto a pagar es de: $", tax)
	} else {
		fmt.Println("Error:", err)
	}
	if tax, err := tax(-2_048_284); err == nil {
		fmt.Println("El impuesto a pagar es de: $", tax)
	} else {
		fmt.Println("Error:", err)
	}
}
