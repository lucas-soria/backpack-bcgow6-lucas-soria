package ejercicio3

import "fmt"

var (
	mes = 5
)

/*
Ejercicio 3 - A qué mes corresponde

Realizar una aplicación que contenga una variable con el número del mes.
Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
Ej: 7, Julio
*/

func QueMesEs() {
	mesesDelAno := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	if mes < 1 || mes > 12 {
		fmt.Println("Ingrese el mes correcto")
	} else {
		fmt.Println(mesesDelAno[mes-1])
	}
}
