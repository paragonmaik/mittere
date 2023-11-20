package reader

import (
	"fmt"
	"mittere/customerror"
	"os"
	"path/filepath"
)

var (
	supportedExtensions []string
)

func Read(filePath string) string {
	fileExt := filepath.Ext(filePath)

	// fmt.Println(fileExt)

	if fileExt != ".json" {
		fmt.Printf("Unsupported file type")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		// fmt.Printf("Error: %d", err)
		fmt.Println("Unsupported file extension: ",
			customerror.ErrValidation)
		//TODO: pass file extension as argument
		os.Exit(1)
	}

	parsedData := string(data)

	return parsedData
}
