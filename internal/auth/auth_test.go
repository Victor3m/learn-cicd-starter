package auth

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type returnVals struct {
		response string
		err      error
	}

	tests := []struct {
		name      string
		headerKey string
		headerVal string
		expected  returnVals
	}{
		{
			name:      "Correct Authorization Header and Correct ApiKey",
			headerKey: "Authorization",
			headerVal: "ApiKey fwof923fgh0f093ur2309urfeijf",
			expected: returnVals{
				response: "fwof923fgh0f093ur2309urfeijf",
				err:      nil,
			},
		},
		{
			name:      "Correct Authorization Header and Incorrect ApiKey",
			headerKey: "Authorization",
			headerVal: "",
			expected: returnVals{
				response: "",
				err:      fmt.Errorf("malformed authorization header"),
			},
		},
		{
			name:      "Incorrect Authorization Header and Incorrect ApiKey",
			headerKey: "Auth",
			headerVal: "fwef wefwifji",
			expected: returnVals{
				response: "",
				err:      fmt.Errorf("no authorization header included"),
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.headerKey, tc.headerVal)
			actual, err := GetAPIKey(header)
			if errors.Is(err, tc.expected.err) && err != nil {
				t.Errorf("Test %v - '%s' FAIL: \nexpected error: %v\nactual error: %v", i, tc.name, tc.expected.err, err)
			}
			if actual != tc.expected.response {
				t.Errorf("Test %v - '%s' FAIL: \nexpected response: %v\nactual response: %v", i, tc.name, tc.expected.response, actual)
			}
		})
	}
}
