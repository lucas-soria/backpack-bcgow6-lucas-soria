package planteadoEnClase

import (
	"fmt"
)

/*
Ejercicio planteado en clase:
Queremos resolver el conocido “FizzBuzz”. FizzBuzz es una consigna que nos plantea realizar un programa que
imprima los números del 1 al 100, reemplazando los múltiplos de 3 por la palabra “Fizz”, los múltiplos de 5
por la palabra “Buzz”, y los múltiplos de 3 y de 5 por la palabra “FizzBuzz”.
*/

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Printf("Fizz")
			fallthrough
		case i%5 == 0:
			fmt.Printf("Buzz")
		case i%3 == 0:
			fmt.Printf("Fizz")
		default:
			fmt.Printf("%v", i)
		}
		fmt.Printf("\n")
	}
}
