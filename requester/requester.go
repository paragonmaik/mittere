package requester

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"mittere/errs"
	"mittere/reader"
	"time"
)

var client *http.Client

func handleUrl(envUrl, fileUrl string) (string, error) {
	if envUrl == "" && fileUrl == "" {
		return "", &errs.ReadErr{Step: "set up request",
			Msg:   "URL is required",
			Cause: errs.ErrInvalidUrl}
	}

	if envUrl == "" {
		return fileUrl, nil
	}

	return envUrl, nil
}

func handleMethod(envMethod, fileMethod string) (string, error) {
	if envMethod == "" && fileMethod == "" {
		return "", &errs.ReadErr{Step: "set up request",
			Msg:   "request method is required",
			Cause: errs.ErrInvalidMethod}
	}

	if envMethod == "" {
		return fileMethod, nil
	}

	return envMethod, nil
}

func makeRequest(url, data, method string, headers map[string]string) {
	body := strings.NewReader(data)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	req.Header = http.Header{
		"Content-Type":  {headers["Content-Type"]},
		"Authorization": {headers["Authorization"]},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
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

	url, err := handleUrl(urlPath, data.Url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	method, err := handleMethod(strings.ToUpper(httpMethod),
		strings.ToUpper(data.Method))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// return request values to pass to the writer package
	makeRequest(url, data.Data, method, data.Headers)
}
