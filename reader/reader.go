package reader

import (
	"encoding/json"
	"fmt"
	"mittere/customerror"
	"os"
	"path/filepath"
	// "strings"
)

// define types in file
type Request struct {
	Url     string          `json:"url"`
	Method  string          `json:"method"`
	Headers json.RawMessage `json:"headers"`
	Data    json.RawMessage `json:"data"`
}

func unmarshalRequest(content []byte) Request {
	request := Request{}

	err := json.Unmarshal(content, &request)
	if err != nil {
		fmt.Println(err)
		// TODO: add custom error
		os.Exit(1)
	}

	return request
}

func Read(filePath string) Request {
	fileExt := filepath.Ext(filePath)

	var request Request

	if fileExt == ".json" {
		content, err := os.ReadFile(filePath)

		if err != nil {
			fmt.Println("file error: ",
				customerror.ErrValidation)
			//TODO: pass file extension as argument
			os.Exit(1)
		}
		request = unmarshalRequest(content)

	} else if fileExt == ".yml" || fileExt == ".yaml" {
		return request

	} else {
		fmt.Printf("Unsupported file type")
		os.Exit(1)
	}

	return request
}
