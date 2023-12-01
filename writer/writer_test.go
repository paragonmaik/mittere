package writer

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

var (
	input = *&http.Response{
		Status: "200 ok",
		Body: io.NopCloser(
			bytes.NewReader([]byte(`{"foo":"bar"}`))),
	}
	input2 = *&http.Response{
		Status: "201 created",
		Body: io.NopCloser(
			bytes.NewReader([]byte(`{"bar":"foo"}`))),
	}
)

func TestWriter(t *testing.T) {
	testCases := []struct {
		name           string
		input          *http.Response
		expectedStatus string
		expectedBody   string
		expectedErr    error
	}{
		{name: "test 1",
			input:          &input,
			expectedStatus: "200 ok",
			expectedBody:   `{"foo":"bar"}`,
		},
		{name: "test 2",
			input:          &input2,
			expectedStatus: "201 created",
			expectedBody:   `{"bar":"foo"}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			status, body := Write(testCase.input, false, "blue")
			if status != testCase.expectedStatus {
				t.Errorf("Expected out: %s, received %s", status,
					testCase.expectedStatus)
			}

			if body != testCase.expectedBody {
				t.Errorf("Expected out: %s, received %s", body,
					testCase.expectedBody)
			}
		})
	}
}
