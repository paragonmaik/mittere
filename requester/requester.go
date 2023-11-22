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
	//field map[string]interface{}
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

func postResp(url, data string) {
	// body passa a ser valor lido do arquivo
	body2 := strings.NewReader(data)
	// body := strings.NewReader(`
	// {
	// "title": "foo",
	// "body": "bar",
	// "userId": 1,
	// }
	// `,
	// )

	// fmt.Println(body2)

	res, err := client.Post(url, "application/json;",
		body2)
	if err != nil {
		fmt.Print(err)
	}

	defer res.Body.Close()

	// fmt.Println(body2)

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))

	// req, err := http.NewRequest(http.MethodPost, url, body)
	// if err != nil {
	// fmt.Print(err)
	// }

	// fmt.Println(req)

	// r, err := newClient().Do(req)
	// if err != nil {
	// fmt.Println(err)
	// }

	// defer r.Body.Close()
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

func ExecRequest(httpMethod string, urlPath string) {
	client = &http.Client{Timeout: 10 * time.Second}

	normalizedMethod := strings.ToUpper(httpMethod)

	data := reader.Read("test.json")
	url := handleUrl(urlPath, data.Url)

	switch normalizedMethod {
	case http.MethodGet:
		getResp(url)
	case http.MethodPost:
		postResp(url, string(data.Data))
	}
}
