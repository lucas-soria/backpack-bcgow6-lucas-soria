package main

import "fmt"

// La interface vacía (The empty interface) es una interface especial
// que es satisfecha por cualquier dato, porque cualquier dato
// implementa por lo menos cero métodos.
// Por lo tanto cuando recibimos un parámetro que debe satisfacer
// la interface vacía, eso significa que el parámetro puede ser
// de cualquier tipo.
func inspeccionar(i interface{}) {

	// Realizamos una aserción de tipo junto con un switch
	switch v := i.(type) {
	case int:
		fmt.Printf("Dos veces %v es %v\n", v, v*2)
	case string:
		fmt.Printf("La cadena %q tiene una longitud de %v\n", v, len(v))
	default:
		fmt.Printf("No conozco los datos de tipo %T!\n", v)
	}
}

func main() {
	inspeccionar(21)
	inspeccionar("hola mundo")
	inspeccionar(true)
	inspeccionar(3.1416)
}
