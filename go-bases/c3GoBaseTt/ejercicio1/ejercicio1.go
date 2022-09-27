package ejercicio1

import "fmt"

/*
Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando
información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo
lugar en memoria para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.
*/

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func cambiarNombre(nuevoNombre string, usuario *Usuario) {
	usuario.Nombre = nuevoNombre
}

func cambiarEdad(nuevaEdad int, usuario *Usuario) {
	usuario.Edad = nuevaEdad
}

func cambiarCorreo(nuevoCorreo string, usuario *Usuario) {
	usuario.Correo = nuevoCorreo
}

func cambiarContrasena(nuevaContrasena string, usuario *Usuario) {
	usuario.Contrasena = nuevaContrasena
}

func RedSocial() {
	usuario := Usuario{
		"Lucas",
		"Soria",
		22,
		"lucas.soria@mercadolibre.com",
		"Ya quisieras saberla",
	}
	fmt.Println("Usuario antes de cambiar:", usuario)
	cambiarNombre("Damian", &usuario)
	cambiarEdad(21, &usuario)
	cambiarCorreo("lucas@correo.com", &usuario)
	cambiarContrasena("1234superseguro", &usuario)
	fmt.Println("Usuario despues de cambiar:", usuario)
}
