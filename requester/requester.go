package requester

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	// "os"
	"time"

	"mittere/reader"
)

var client *http.Client

type Data struct {
	data map[string]json.RawMessage
}

func handleUrl(envUrl, fileUrl string) string {
	if envUrl == "" && fileUrl == "" {
		fmt.Println("URL is required")
		os.Exit(1)
	}

	if envUrl == "" {
		return fileUrl
	}

	return envUrl
}

func handleMethod(envMethod, fileMethod string) string {
	if envMethod == "" && fileMethod == "" {
		fmt.Println("Method is required")
		os.Exit(1)
	}

	if envMethod == "" {
		return fileMethod
	}

	return envMethod
}

func postResp(url, data string) {
	body2 := strings.NewReader(data)
	res, err := client.Post(url, "application/json;",
		body2)
	if err != nil {
		fmt.Print(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))

	defer res.Body.Close()
}

func getResp(url string) {
	//TODO: handle error
	resp, _ := client.Get(url)
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

func ExecRequest(httpMethod, urlPath, filepath string) {
	client = &http.Client{Timeout: 10 * time.Second}

	data, err := reader.Read(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(data)
	url := handleUrl(urlPath, data.Url)
	method := handleMethod(strings.ToUpper(httpMethod),
		strings.ToUpper(data.Method))

	switch method {
	case http.MethodGet:
		getResp(url)
	case http.MethodPost:
		postResp(url, string(data.Data))
	}
}
