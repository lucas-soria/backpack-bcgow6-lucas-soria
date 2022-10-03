package clase

import (
	"encoding/json"
	"fmt"
)

type product struct {
	Name      string  `json:"name" db:"product_name"`
	Price     float64 `json:"price" db:"price"`
	Published bool    `json:"published" db:"is_published"`
}

var (
	notebook = product{
		Name:      "MacBook Pro",
		Price:     1399.99,
		Published: true,
	}
	notebookBytes = []byte(`{"Name": "Samsung S22 ULTRA", "Price": 699.99, "Published": true}`)
)

func UseJSON() {
	if jsonData, err := json.Marshal(notebook); err != nil {
		panic("Ahhhhhh")
	} else {
		fmt.Println(string(jsonData))
	}
	var p product
	if err := json.Unmarshal(notebookBytes, &p); err != nil {
		panic("Ahhhhhhhhhhhhhh")
	} else {
		fmt.Printf("%+v", p)
	}
}
