package ejercicio3

import (
	"fmt"
	"os"
)

const (
	ErrorMinimum = "Error: the salary entered does not meet the taxable minimum"
)

/*
Ejercicio 3 - Impuestos de salario #3
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por
parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola
deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el
valor de tipo int pasado por parámetro).
*/

func MinimumNonTaxableErrorf() {
	salary := 30_000
	if salary < 150_000 {
		err := fmt.Errorf(
			"%s\n\tMore info: the minimum taxable amount is 150.000, but the salary entered was: %d",
			ErrorMinimum,
			salary,
		)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Has to pay taxes")
}
