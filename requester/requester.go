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

// TODO: use http.NewRequest(method, url, body)

func PostResp(url string) {
	body := strings.NewReader(`
    {
    "title": "foo",
    "body": "bar",
    "userId": 1,
    }
    `,
	)

	res, err := client.Post(url, "application/json;",
		body)
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

func GetResp(url string) {
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

func Request(httpMethod string, urlPath string) {
	// fmt.Println(httpMethod)
	client = &http.Client{Timeout: 10 * time.Second}

	normalizedMethod := strings.ToUpper(httpMethod)

	reader.Read("main.go")

	switch normalizedMethod {
	case http.MethodGet:
		GetResp(urlPath)
	case http.MethodPost:
		PostResp("https://jsonplaceholder.typicode.com/posts")
	}

}
