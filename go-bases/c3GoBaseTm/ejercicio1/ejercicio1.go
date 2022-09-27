package ejercicio1

import (
	"fmt"
	"os"
)

/*
Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados,
separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

type product struct {
	id       int
	price    float64
	quantity int
}

func saveFile(filePath string, products []product) {
	fmt.Println("Escribiendo en el archivo...")
	for _, p := range products {
		text := fmt.Sprintf("%d, %.2f, %d\n", p.id, p.price, p.quantity)
		var fullText string
		if previousTextB, err := os.ReadFile(filePath); err == nil {
			fullText = string(previousTextB) + text
		} else {
			fullText = "ID, Price, Quantity\n" + text
		}
		if err := os.WriteFile(filePath, []byte(fullText), 0777); err != nil {
			fmt.Println("\tHubo un error en la escritura del archivo")
			fmt.Println("\t\tError:", err)
			os.Exit(1)
		}
	}
	fmt.Println("OK")
}

func SaveFileTool() {
	products := []product{{1, 2.5, 5}, {2, 5.5, 9}, {3, 9.5, 18}}
	saveFile(
		"/Users/LSORIA/Documents/bootcamp/backpack-bcgow6-lucas-soria/go-bases/c3GoBaseTm/ejercicio1.csv",
		products)
}
