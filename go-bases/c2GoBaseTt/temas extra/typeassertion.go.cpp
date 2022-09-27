// Type assertion (aserción de tipo)
package main

import "fmt"

func main() {
	var i interface{} = "hola"

	s := i.(string)
	fmt.Println(s)

	// Cuando hacemos una aserción recibiendo
	// dos valores podemos saber si la aserción
	// fue exitosa sin abortar el programa
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Si la aserción no tiene éxito y la forma
	// usada no recibe dos valores, el programa
	// aborta con un panic()
	f = i.(float64)
	fmt.Println(f)
}
