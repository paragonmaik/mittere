package reader

import (
	"encoding/json"
	"fmt"
	"mittere/customerror"
	"os"
	"path/filepath"
)

var (
	supportedExtensions []string
)

// define types in file
type Request struct {
	Url     string          `json:"url"`
	Method  string          `json:"method"`
	Headers json.RawMessage `json:"headers"`
	Data    json.RawMessage `json:"data"`
}

func unmarshalRequest(filepath string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}
	request := Request{}

	err2 := json.Unmarshal(content, &request)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Printf("%s \n", request)
}

func Read(filePath string) string {
	fileExt := filepath.Ext(filePath)

	// fmt.Println(fileExt)
	unmarshalRequest(filePath)

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
