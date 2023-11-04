package requester

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var client *http.Client

func GetResp(url string) {
	resp, _ := http.Get(url)
	respData, _ := io.ReadAll(resp.Body)

	fmt.Print(string(respData))
}

func Request() {
	client = &http.Client{Timeout: 10 * time.Second}

	//GetResponse("https://pokeapi.co/api/v2/pokemon/ditto")
	GetResp("https://jsonplaceholder.typicode.com/todos/1")
}
