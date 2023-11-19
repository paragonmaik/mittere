package reader

import (
	"fmt"
	"os"
)

func Read() {
	data, err := os.ReadFile("main.go")
	if err != nil {
		fmt.Printf("Error: %d", err)
	}

	parsedData := string(data)

	fmt.Println(parsedData)
}
