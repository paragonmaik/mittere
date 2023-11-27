package requester

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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

func postResp(url, data string, headers map[string]string) {
	body := strings.NewReader(data)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// TODO: add error
		fmt.Print(err)
	}

	req.Header = http.Header{
		"Content-Type":  {headers["Content-Type"]},
		"Authorization": {headers["Authorization"]},
	}

	res, err := client.Do(req)
	if err != nil {
		// TODO: add error
		fmt.Print(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))

	defer res.Body.Close()
}

func getResp(url string, headers map[string]string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// TODO: add error
		fmt.Print(err)
	}

	req.Header = http.Header{
		"Content-Type":  {headers["Content-Type"]},
		"Authorization": {headers["Authorization"]},
	}

	res, err := client.Do(req)
	if err != nil {
		// TODO: add error
		fmt.Print(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))

	defer res.Body.Close()
}

func ExecRequest(httpMethod, urlPath, filepath string) {
	client = &http.Client{Timeout: 10 * time.Second}

	data, err := reader.Read(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// fmt.Println(data.Headers)
	url := handleUrl(urlPath, data.Url)
	method := handleMethod(strings.ToUpper(httpMethod),
		strings.ToUpper(data.Method))

	switch method {
	case http.MethodGet:
		getResp(url, data.Headers)
	case http.MethodPost:
		postResp(url, data.Data, data.Headers)
	}
}
