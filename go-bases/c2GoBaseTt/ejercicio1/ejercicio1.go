package ejercicio1

import (
	"fmt"
	"time"
)

/*
Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el
detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que
tenga un método detalle
*/

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    time.Time
}

func (alumno Alumno) detalle() {
	fmt.Printf(
		"Nombre: %+v\nApellido: %+v\nDNI: %+v\nFecha: %+v\n",
		alumno.Nombre,
		alumno.Apellido,
		alumno.DNI,
		alumno.Fecha)
}

func DetallarAlumno() {
	alumno := Alumno{
		Nombre:   "Lucas",
		Apellido: "Soria",
		DNI:      42670450,
		Fecha:    time.Date(2000, 5, 9, 14, 30, 45, 100, time.Local),
	}
	fmt.Println("Los detalles del Alumno son: ")
	alumno.detalle()
}
