package ejercicio4

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Ejercicio 4 - Ordenamiento
Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
un arreglo de números enteros con 100 valores
un arreglo de números enteros con 1000 valores
un arreglo de números enteros con 10000 valores

Para instanciar las variables utilizar rand
package main

import (
   "math/rand"
)
func main() {
   variable1 := rand.Perm(100)
   variable2 := rand.Perm(1000)
   variable3 := rand.Perm(10000)
}

Se debe realizar el ordenamiento de cada una por:
Ordenamiento por inserción
Ordenamiento por burbuja
Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento
fue mejor para cada arreglo
*/

type SortingAlgorithm struct {
	Name       string
	Speed100   int64
	Speed1000  int64
	Speed10000 int64
}

func register(lenSlice int, totTime int64, algorithm *SortingAlgorithm) {
	switch lenSlice {
	case 100:
		algorithm.Speed100 = totTime
	case 1_000:
		algorithm.Speed1000 = totTime
	case 10_000:
		algorithm.Speed10000 = totTime
	}
}

func bubble(slice []int, algorithm *SortingAlgorithm, cLock chan bool) {
	now := time.Now()
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	totTime := time.Now().Sub(now).Nanoseconds()
	register(len(slice), totTime, algorithm)
	cLock <- true
}

func insertion(slice []int, algorithm *SortingAlgorithm, cLock chan bool) {
	now := time.Now()
	var n = len(slice)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
			j = j - 1
		}
	}
	totTime := time.Now().Sub(now).Nanoseconds()
	register(len(slice), totTime, algorithm)
	cLock <- true
}

func selection(slice []int, algorithm *SortingAlgorithm, cLock chan bool) {
	now := time.Now()
	var n = len(slice)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
	totTime := time.Now().Sub(now).Nanoseconds()
	register(len(slice), totTime, algorithm)
	cLock <- true
}

func Ordenamiento() {
	cLock := make(chan bool)
	slice100 := rand.Perm(100)
	slice1_000 := rand.Perm(1_000)
	slice10_000 := rand.Perm(10_000)
	bubbleSort := SortingAlgorithm{Name: "Bubble"}
	insertionSort := SortingAlgorithm{Name: "Insertion"}
	selectionSort := SortingAlgorithm{Name: "Selection"}
	go bubble(slice100, &bubbleSort, cLock)
	go bubble(slice1_000, &bubbleSort, cLock)
	go bubble(slice10_000, &bubbleSort, cLock)
	for i := 0; i < 3; i++ {
		<-cLock
	}
	go insertion(slice100, &insertionSort, cLock)
	go insertion(slice1_000, &insertionSort, cLock)
	go insertion(slice10_000, &insertionSort, cLock)
	for i := 0; i < 3; i++ {
		<-cLock
	}
	go selection(slice100, &selectionSort, cLock)
	go selection(slice1_000, &selectionSort, cLock)
	go selection(slice10_000, &selectionSort, cLock)
	for i := 0; i < 3; i++ {
		<-cLock
	}
	fmt.Println("Velocidades en ns:")
	fmt.Printf("%+v\n", bubbleSort)
	fmt.Printf("%+v\n", insertionSort)
	fmt.Printf("%+v\n", selectionSort)
	fmt.Printf("\n!! Parece que siempre es mas rápido el insertion\n")
	fmt.Printf("Disclaimer:\n\tLos algoritmos 3 algoritmos están sacados de internet")
}
