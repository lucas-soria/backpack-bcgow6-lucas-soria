package clase

import (
	"fmt"
	"os"
)

func Tests() {
	env := os.Getenv("HOME")
	fmt.Println(env)
	err := os.WriteFile("./testFile", []byte(env), 0644)
	fmt.Println(err)
	file, ok := os.ReadFile("./testFile")
	if ok != nil {
		fmt.Println(file)
	}
}
