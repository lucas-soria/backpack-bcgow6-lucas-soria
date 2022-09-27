package ejercicio5

import (
	"errors"
	"fmt"
)

const (
	DOG       = "dog"
	CAT       = "cat"
	HAMSTER   = "hamster"
	TARANTULA = "tarantula"
)

/*
Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales
que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que
retorne una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
*/

func dogFood(quantity int) float64 {
	return float64(quantity) * 10.0
}

func catFood(quantity int) float64 {
	return float64(quantity) * 5.0
}

func hamsterFood(quantity int) float64 {
	return float64(quantity) * 0.250
}

func tarantulaFood(quantity int) float64 {
	return float64(quantity) * 0.150
}

func animal(animalType string) (function func(quantity int) float64, err error) {
	switch animalType {
	case DOG:
		function = dogFood
	case CAT:
		function = catFood
	case HAMSTER:
		function = hamsterFood
	case TARANTULA:
		function = tarantulaFood
	default:
		err = errors.New("no existe ese animal")
	}
	return
}

func FoodCalculator() {
	var amountFood float64

	if animalDog, err := animal(DOG); err == nil {
		amountFood += animalDog(5)
	} else {
		fmt.Println("Error: ", err)
	}
	if animalCat, err := animal(CAT); err == nil {
		amountFood += animalCat(8)
	} else {
		fmt.Println("Error: ", err)
	}

	fmt.Println("La cantidad de comida que hay que comprar es: ", amountFood)
}
