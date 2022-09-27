package ejercicio2

import "fmt"

/*
Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que
represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set: Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho,
si es cuadrática y cuál es el valor máximo.
*/

type Matrix struct {
	values [][]float64
	height int
	width  int
	square bool
	max    float64
}

func (matrix *Matrix) Set(values [][]float64) {
	matrix.values = values
	matrix.height = len(values)
	matrix.width = len(values[0])
	matrix.square = matrix.height == matrix.width
	matrix.max = max(matrix.values)
}

func (matrix Matrix) Print() {
	for _, row := range matrix.values {
		fmt.Println(row)
	}
}

func max(values [][]float64) (max float64) {
	max = values[0][0]
	for _, row := range values {
		for _, element := range row {
			if element > max {
				max = element
			}
		}
	}
	return
}

func Neo() {
	values := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {7, 8, 9}}
	m := Matrix{}
	m.Set(values)
	fmt.Printf("La matriz: ")
	fmt.Println(m)
	fmt.Println("Tiene los siguientes valores (impresos de forma linda):")
	m.Print()
}
