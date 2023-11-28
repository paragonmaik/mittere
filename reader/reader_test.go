package reader

import (
	"errors"
	"mittere/errs"
	"testing"
)

type request struct {
	Url     string
	Method  string
	Data    string
	Headers string
}

func TestReader(t *testing.T) {
	testCases := []struct {
		name        string
		filepath    string
		expected    request
		expectedErr error
	}{{
		name:     "test 1",
		filepath: "../testdata/test2.json",
		expected: request{
			Url:    "https://jsonplaceholder.typicode.com/posts/1",
			Method: "get",
		}},
		{
			name:     "test 2",
			filepath: "../testdata/test2.yaml",
			expected: request{
				Url:    "https://jsonplaceholder.typicode.com/posts/1",
				Method: "get",
			}},
		{
			name:        "test 3",
			filepath:    "../testdata/testErr.txt",
			expectedErr: &errs.ReadErr{Step: "read file"},
		},
		{
			name:        "test 4",
			filepath:    "../testdata/nonExistantFile.json",
			expectedErr: &errs.ReadErr{Step: "read file"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req, err := Read(testCase.filepath)
			if testCase.expectedErr != nil {
				if err == nil {
					t.Errorf("Expected error; received none")
				}

				if !errors.Is(err, testCase.expectedErr) {
					t.Errorf("Expected error: %q, received: %q",
						testCase.expectedErr, err)
				}
			}

			if req.Url != testCase.expected.Url {
				t.Errorf("Expected out: %s, received %s", req.Url,
					testCase.expected.Url)
			}

			if req.Method != testCase.expected.Method {
				t.Errorf("Expected out: %s, received %s", req.Method,
					testCase.expected.Method)
			}
		})
	}
}
