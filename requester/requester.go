package requester

import (
	"fmt"
	"net/http"
	"net/url"
)

func Request() {
	resp, err := http.PostForm("http://localhost:3000/login",
		url.Values{"email": {"donkey@example.com"}, "password": {"12345678"}})

	fmt.Print(resp, err)
}
