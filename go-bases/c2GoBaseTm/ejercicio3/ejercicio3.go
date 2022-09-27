package ejercicio3

import (
	"errors"
	"fmt"
)

/*
Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas
trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y
la categoría, y que devuelva su salario.
*/

func salaryCalculator(minutes int, category string) (salary float64, err error) {
	switch category {
	case "A":
		salary = float64(minutes) / 60 * 3_000 * 1.5
	case "B":
		salary = float64(minutes) / 60 * 1_500 * 1.2
	case "C":
		salary = float64(minutes) / 60 * 1_000
	default:
		err = errors.New("categoria no válida")
	}
	return
}

func SalaryCalculatorTool() {
	if salary, err := salaryCalculator(95, "A"); err == nil {
		fmt.Println("El salario del empleado es:", salary)
	} else {
		fmt.Println("Error:", err)
	}
	if salary, err := salaryCalculator(765, "D"); err == nil {
		fmt.Println("El salario del empleado es:", salary)
	} else {
		fmt.Println("Error:", err)
	}
}
