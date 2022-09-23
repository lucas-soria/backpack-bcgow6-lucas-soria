package ejercicio1

import "fmt"

var (
	name    = "Lucas Damian Soria Gava"
	address = "*Secret* 276"
)

func NombreYDireccion() {
	fmt.Printf("Nombre: %v\n", name)
	fmt.Printf("Direccion: %v\n", address)
}
