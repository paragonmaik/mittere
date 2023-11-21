package requester

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

// define types in file
type Request struct {
	Body string `json:"body"`
}

// TODO: use http.NewRequest(method, url, body)

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

	fmt.Println(body2)

	res, err := client.Post(url, "application/json;",
		body2)
	if err != nil {
		fmt.Print(err)
	}

	defer res.Body.Close()

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
	// fmt.Println(httpMethod)
	client = &http.Client{Timeout: 10 * time.Second}

	normalizedMethod := strings.ToUpper(httpMethod)

	data := reader.Read("test.json")

	switch normalizedMethod {
	case http.MethodGet:
		getResp(urlPath)
	case http.MethodPost:
		postResp("https://jsonplaceholder.typicode.com/posts", data)
	}

}
