package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		input    http.Header
		want     string
		errorMsg string
	}{
		{input: http.Header{}, want: "", errorMsg: "no authorization header included"},
		{input: http.Header{"Authorization": []string{"ApiKeyabcd"}}, want: "", errorMsg: "malformed authorization header"},
		{input: http.Header{"Authorization": []string{"ApiKe abcd"}}, want: "", errorMsg: "malformed authorization header"},
		{input: http.Header{"Authorization": []string{"ApiKey abcd"}}, want: "abcd", errorMsg: ""},
	}

	for _, tc := range tests {
		val, err := GetAPIKey(tc.input)
		if !(val == tc.want || err.Error() == tc.errorMsg) {
			t.Fatalf("expected value: %v, got: %v; expected error: %v, got: %v", tc.want, val, tc.errorMsg, err.Error())
		}
	}
}
