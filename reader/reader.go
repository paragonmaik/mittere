package reader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type RequestJson struct {
	Url     string          `json:"url"`
	Method  string          `json:"method"`
	Headers json.RawMessage `json:"headers"`
	Data    json.RawMessage `json:"data"`
}

type RequestYml struct {
	Url     string            `yaml:"url"`
	Method  string            `yaml:"method"`
	Headers map[string]string `yaml:"headers"`
	Data    map[string]string `yaml:"data"`
}

type Request struct {
	Url     string
	Method  string
	Headers string
	Data    string
}

func mapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for k, v := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", k, v)
	}
	return b.String()
}

func unmarshalRequestJson(content []byte) RequestJson {
	request := RequestJson{}

	err := json.Unmarshal(content, &request)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return request
}

func unmarshalRequestYml(content []byte) RequestYml {
	request := RequestYml{}

	err := yaml.Unmarshal(content, &request)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return request
}

func Read(filePath string) (Request, error) {
	fileExt := filepath.Ext(filePath)

	var requestJson RequestJson
	var requestYml RequestYml
	var request Request

	if fileExt == ".json" {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return request, &readErr{step: "read file",
				msg: "read file failed", cause: err}
		}
		requestJson = unmarshalRequestJson(content)

		request.Url = requestJson.Url
		request.Data = string(requestJson.Data)
		request.Method = requestJson.Method
		request.Headers = string(requestJson.Headers)

	} else if fileExt == ".yml" || fileExt == ".yaml" {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return request, &readErr{step: "read file",
				msg: "read file failed", cause: err}
		}
		requestYml = unmarshalRequestYml(content)

		request.Url = requestYml.Url
		request.Data = mapToString(requestYml.Data)
		request.Method = requestYml.Method
		request.Headers = mapToString(requestYml.Headers)

	} else {
		return request, &readErr{step: "read file",
			msg: "read file failed", cause: ErrInvalidExt}
	}

	return request, nil
}
