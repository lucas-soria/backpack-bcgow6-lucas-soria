package ejercicio2

import (
	"errors"
	"fmt"
	"os"
)

const (
	ErrorMinimum = "Error: the salary entered does not meet the taxable minimum"
)

/*
Ejercicio 2 - Impuestos de salario #2

Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,
se implemente “errors.New()”.
*/

func MinimumNonTaxableNew() {
	salary := 30_000
	if salary < 150_000 {
		fmt.Println(errors.New(ErrorMinimum))
		os.Exit(1)
	}
	fmt.Println("Has to pay taxes")
}
