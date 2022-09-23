package ejercicio2

import "fmt"

var (
	temperatura float32 = 17
	humedad     float32 = 24
	presion     uint16  = 1016
)

func Clima() {
	fmt.Printf("Temperatura actual: %v °C\nHumedad ambiente: %v\nPresion atmosférica: %vhPa\n", temperatura, humedad, presion)
}
