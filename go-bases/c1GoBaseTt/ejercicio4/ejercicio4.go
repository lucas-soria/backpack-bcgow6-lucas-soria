package ejercicio4

import "fmt"

var (
	employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
)

/*
Ejercicio 4 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa,
ayuda a imprimir la edad de Benjamin.

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario:
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.
*/

func CalcularEdad() {
	fmt.Println("La edad de Benjamin es: ", employees["Benjamin"])
	var c uint8
	for _, element := range employees {
		if element > 21 {
			c++
		}
	}
	fmt.Println("La cantidad de empleados mayores a 21 es: ", c)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}
