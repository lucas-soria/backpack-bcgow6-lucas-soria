package ejercicio2

import (
	"errors"
	"fmt"
)

/*
Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y
un error en caso que uno de los números ingresados sea negativo
*/

func Average(grades ...float64) (average float64, err error) {
	for _, grade := range grades {
		if grade < 0 {
			return 0, errors.New("nota negativa")
		} else {
			average += float64(grade)
		}
	}
	average /= float64(len(grades))
	return
}

func AverageCalculator() {
	if average, err := Average(6, 7, 5, 3, 2, 5, 2, 5, 3, 5, 3); err == nil {
		fmt.Println("El promedio de notas es de:", average)
	} else {
		fmt.Println("Error:", err)
	}
	if average, err := Average(6, 7, 5, 3, 2, 5, 2, 5, 3, 5, -3); err == nil {
		fmt.Println("El promedio de notas es de:", average)
	} else {
		fmt.Println("Error:", err)
	}
}
