package ejercicio1

import "fmt"

var (
	palabra = "Esternocleidomastoideo"
)

/*
Ejercicio 1 - Letras de una palabra
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras
por separado para deletrearla.
Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
Luego imprimí cada una de las letras.
*/

func AnalizarPalabra() {
	fmt.Println("La cantidad de letras de la palabra es: ", len(palabra))
	for i := range palabra {
		fmt.Printf("%c ", palabra[i])
	}
	fmt.Printf("\n")
}
