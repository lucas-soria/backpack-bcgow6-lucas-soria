package clase

func TestPunteros() {
	variable := "Hola"
	puntero := &variable
	println("Valor de la variable:", variable)
	println("Direccion de la variable:", &variable)
	println("Valor del puntero (direccion de la variable)", puntero)
	println("Direccion del puntero:", &puntero)
	println("Valor de la variable a la que apunta el puntero:", *puntero)
}
