// Utilización de Type Casting (encasillamiento de un tipo de dato)

package main

import "fmt"

func main() {

	var suma int = 100
	var importe int = 19
	var promedio float32

	// Convertimos explícitamente los valores de las variables
	// suma e importe para que ambas sean del tipo de dato float32.
	promedio = float32(suma) / float32(importe)

	fmt.Printf("Promedio = %f\n", promedio)
}
