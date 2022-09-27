package ejercicio2

import (
	"fmt"
	"os"
	"strings"
)

/*
Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado, para ello requiere que:
se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y
a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando
precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50
*/

func prettySprint(dataMatrix [][]string) (distilledText string) {
	for _, i := range dataMatrix {
		if len(i) == 3 {
			distilledText += fmt.Sprintf("%s\t%10s\t%10s\n", i[0], i[1], i[2])
		}
	}
	return
}

func splitText(raw string) [][]string {
	var dataMatrix [][]string
	rows := strings.Split(raw, "\n")
	for _, row := range rows {
		dataMatrix = append(dataMatrix, strings.SplitN(row, ", ", 3))
	}
	return dataMatrix
}

func readFile(filePath string) {
	if textB, err := os.ReadFile(filePath); err == nil {
		dataMatrix := splitText(string(textB))
		distilledText := prettySprint(dataMatrix)
		fmt.Println(distilledText)
		fmt.Println("OK")
	} else {
		fmt.Println("Hubo un error al leer el archivo")
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func ReadFileTool() {
	fmt.Println("Leyendo el archivo...")
	readFile("/Users/LSORIA/Documents/bootcamp/backpack-bcgow6-lucas-soria/go-bases/c3GoBaseTm/ejercicio1.csv")
}
