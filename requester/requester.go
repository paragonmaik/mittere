package requester

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "os"
	"time"

	"mittere/reader"
)

var client *http.Client

type Data struct {
	//field map[string]interface{}
	data map[string]json.RawMessage
}

func GetResp(url string) {
	resp, _ := http.Get(url)
	respData, _ := io.ReadAll(resp.Body)
	//var data map[string]interface{}
	var data Data

	err := json.Unmarshal(respData, &data.data)
	if err != nil {
		fmt.Print(err)
	}

	// fmt.Print(data.data)
	// fmt.Print(string(respData))
	fmt.Printf("{\n")
	for k, v := range data.data {
		fmt.Printf("\t%v %v, \n", k, string(v))
	}
	fmt.Printf("}\n")
	// fmt.Println(os.Args[1:])
}

func Request(httpMethod string, urlPath string) {
	// fmt.Println(httpMethod)
	client = &http.Client{Timeout: 10 * time.Second}

	// GetResp("https://jsonplaceholder.typicode.com/todos/1")
	// switch httpMethod {
	// case "get":
	// GetResp(urlPath)
	// }

	reader.Read()
}
