package ejercicio1

import (
	"fmt"
	"os"
)

const (
	ErrorMinimum = "Error: the salary entered does not meet the taxable minimum"
)

/*
Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error:
el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000.
Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

type minimumTaxable struct {
	message string
}

func (err minimumTaxable) Error() string {
	return err.message
}

func MinimumNonTaxableError() {
	salary := 300_000
	if salary < 150_000 {
		fmt.Println(minimumTaxable{ErrorMinimum}.Error())
		os.Exit(1)
	}
	fmt.Println("Has to pay taxes")
}
