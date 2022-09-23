package ejercicio2

import "fmt"

/*
Ejercicio 2 - Préstamo

Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes
cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

var (
	edad       uint8
	empleado   bool
	antiguedad uint8
	sueldo     uint
)

func OtorgarPrestamo() {
	if edad > 22 {
		fmt.Println("Tiene mas de 21")
		if empleado {
			fmt.Println("Está empleado")
			if antiguedad > 1 {
				fmt.Println("Tiene mas de un año de antigüedad")
				if sueldo > 100000 {
					fmt.Println("No le cobra intereses")
				}
			} else {
				fmt.Println("No tiene mas de 1 año de antiguedad")
			}
		} else {
			fmt.Println("No está empleado")
		}
	} else {
		fmt.Println("No tiene mas de 21")
	}
}
