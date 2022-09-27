package ejercicio3

import (
	"fmt"
	"math"
)

/*
Ejercicio 3 - Calcular Precio
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y
Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se
realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega
a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando
el total de los 3).
*/

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicio struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func sumarProductos(productos *[]Producto, canal chan float64) {
	total := 0.0
	for _, producto := range *productos {
		total += producto.Precio * float64(producto.Cantidad)
	}
	canal <- total
}

func sumarServcios(servicios *[]Servicio, canal chan float64) {
	total := 0.0
	for _, servicio := range *servicios {
		total += servicio.Precio * (math.Ceil(float64(servicio.MinutosTrabajados) / 30))
	}
	canal <- total
}

func sumarMantenimiento(mantenimientos *[]Mantenimiento, canal chan float64) {
	total := 0.0
	for _, mantenimiento := range *mantenimientos {
		total += mantenimiento.Precio
	}
	canal <- total
}

func sacarTotal(productos *[]Producto, servicios *[]Servicio, mantenimientos *[]Mantenimiento) (total float64) {
	canal := make(chan float64)
	go sumarProductos(productos, canal)
	go sumarServcios(servicios, canal)
	go sumarMantenimiento(mantenimientos, canal)
	for i := 0; i < 3; i++ {
		total += <-canal
	}
	return
}

func CalcularPrecio() {
	productos := []Producto{{Precio: 2.2, Cantidad: 2}, {Precio: 5.6, Cantidad: 1}}
	servicios := []Servicio{{Precio: 2, MinutosTrabajados: 27}, {Precio: 7, MinutosTrabajados: 200}}
	mantenimientos := []Mantenimiento{{Precio: 2}, {Precio: 5}}
	fmt.Println("El monto final es de:", sacarTotal(&productos, &servicios, &mantenimientos))
}
