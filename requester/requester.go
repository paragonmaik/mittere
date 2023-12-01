package requester

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"mittere/errs"
	"mittere/reader"
	"mittere/writer"
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

func setHeaders(h http.Header, headers map[string]string) {
	for k, v := range headers {
		h.Set(k, v)
	}
}

func makeRequest(url, data, method string,
	headers map[string]string) *http.Response {
	body := strings.NewReader(data)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	setHeaders(req.Header, headers)

	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return res
}

func ExecRequest(httpMethod, urlPath,
	filepath, selectedColor string,
	colorize bool) string {
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

	res := makeRequest(url, data.Data, method, data.Headers)

	status, body := writer.Write(res, colorize, selectedColor)

	return status + body
}
