package ejercicio2

import "fmt"

/*
Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del
programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.
*/

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

func nuevoProducto(nombre string, precio float64) Producto {
	return Producto{
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: 1,
	}
}

func agregarProducto(usuario *Usuario, producto *Producto, cantidad int) {
	producto.Cantidad *= cantidad
	usuario.Productos = append(usuario.Productos, *producto)
}

func borrarProductos(usuario *Usuario) {
	usuario.Productos = []Producto{}
}

func Ecommerce() {
	usuario := Usuario{
		Nombre:   "Lucas",
		Apellido: "Soria",
		Correo:   "lucas.soria@mercadolibre.com",
	}
	fmt.Println("Estado inicial del usuario:", usuario)
	producto := nuevoProducto("Computadora", 500_000)
	fmt.Println("Estado inicial del producto:", producto)
	agregarProducto(&usuario, &producto, 7)
	fmt.Println("Estado luego de agregar del usuario:", usuario)
	fmt.Println("Estado luego de agregar del producto:", producto)
	borrarProductos(&usuario)
	fmt.Println("Estado luego de borrar del usuario:", usuario)
	fmt.Println("Estado luego de borrar del producto:", producto)
}
