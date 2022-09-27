package ejercicio4

import (
	"errors"
	"fmt"
)

const (
	MINIMUM = "minimum"
	AVERAGE = "average"
	MAXIMUM = "maximum"
)

/*
Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los
alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y
que devuelva otra función (y un mensaje en caso que el cálculo no esté definido)
que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior
*/

func maximum(grades ...float64) (max float64) {
	max = grades[0]
	for _, grade := range grades {
		if grade > max {
			max = grade
		}
	}
	return
}

func average(grades ...float64) (average float64) {
	for _, grade := range grades {
		average += grade
	}
	return average / float64(len(grades))
}

func minimum(grades ...float64) (min float64) {
	min = grades[0]
	for _, grade := range grades {
		if grade < min {
			min = grade
		}
	}
	return
}

func handler(operation string) (function func(grades ...float64) float64, err error) {
	switch operation {
	case MINIMUM:
		function = minimum
	case MAXIMUM:
		function = maximum
	case AVERAGE:
		function = average
	default:
		err = errors.New("operación no permitida")
	}
	return
}

func GradeCalculatorTool() {
	if minFunc, err := handler(MINIMUM); err == nil {
		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println("Mínimo:", minValue)
	} else {
		fmt.Println("Error:", err)
	}
	if averageFunc, err := handler(AVERAGE); err == nil {
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Promedio:", averageValue)
	} else {
		fmt.Println("Error:", err)
	}
	if maxFunc, err := handler(MAXIMUM); err == nil {
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Máximo:", maxValue)
	} else {
		fmt.Println("Error:", err)
	}
	if maxFunc, err := handler("alksnd"); err == nil {
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Máximo:", maxValue)
	} else {
		fmt.Println("Error:", err)
	}
}
