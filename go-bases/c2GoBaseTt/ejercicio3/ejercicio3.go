package ejercicio3

import "fmt"

const (
	PEQUENO = "pequeño"
	MEDIANO = "mediano"
	GRANDE  = "grande"
)

/*
Ejercicio 3 - Productos
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y
retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
* Pequeño: El costo del producto (sin costo adicional)
* Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
* Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
* Crear una estructura “tienda” que guarde una lista de productos.
* Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
* Crear una interface “Producto” que tenga el método “CalcularCosto”
* Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
* Interface Producto:
* El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales
si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
*/

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto) []Producto
}

type producto struct {
	p              Producto
	tipoDeProducto string
	nombre         string
	precio         float64
}

type tienda struct {
	t         Ecommerce
	productos []Producto
}

func (prod producto) CalcularCosto() (costoTotal float64) {
	switch prod.tipoDeProducto {
	case PEQUENO:
		costoTotal = prod.precio
	case MEDIANO:
		costoTotal = prod.precio * 1.03
	case GRANDE:
		costoTotal = prod.precio*1.06 + 2500
	}
	return
}

func (tienda tienda) Total() (total float64) {
	for _, producto := range tienda.productos {
		total += producto.CalcularCosto()
	}
	return
}

func (tienda *tienda) Agregar(p Producto) []Producto {
	tienda.productos = append(tienda.productos, p)
	return tienda.productos
}

func nuevoProducto(tipo string, nombre string, precio float64) Producto {
	return &producto{
		tipoDeProducto: tipo,
		nombre:         nombre,
		precio:         precio,
	}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func CrearTiendasYProductos() {
	pCelu := nuevoProducto(PEQUENO, "Samsung S22 Ultra", 300_000)
	pHeladera := nuevoProducto(GRANDE, "Heladera Liliana No Frost", 500_000)
	t := nuevaTienda()
	t.Agregar(pCelu)
	fmt.Println("Agregamos un celu: ")
	fmt.Println(t.Total())
	t.Agregar(pHeladera)
	fmt.Println("Agregamos una heladera: ")
	fmt.Println(t.Total())
	fmt.Println("El estado de la tienda es: ")
	fmt.Println(t)
}
