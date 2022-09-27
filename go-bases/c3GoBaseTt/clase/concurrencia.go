package clase

import "fmt"

func haceAlgo(c chan int) {
	c <- 1
	c <- 2
}

func Concurrencia() {
	c := make(chan int)
	go haceAlgo(c)
	hola := <-c
	fmt.Println(hola)
}
