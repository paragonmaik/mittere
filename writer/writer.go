package writer

import (
	"fmt"
	"io"
	"net/http"
)

// take colorize flag parameter
func Write(res *http.Response) {
	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))

	defer res.Body.Close()
}
