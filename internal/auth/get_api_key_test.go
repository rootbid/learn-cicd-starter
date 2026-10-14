package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("returns api key when auth header is valid", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey new-key")

		got, err := GetAPIKey(headers)
		if err != nil {
			t.Fatalf("GetAPIKey() returned unexpected error: %v", err)
		}
		if got != "new-key" {
			t.Fatalf("GetAPIKey() = %q, want %q", got, "new-key")
		}
	})

	t.Run("returns error when auth header is missing", func(t *testing.T) {
		headers := http.Header{}

		_, err := GetAPIKey(headers)
		if err != ErrNoAuthHeaderIncluded {
			t.Fatalf("GetAPIKey() error = %v, want %v", err, ErrNoAuthHeaderIncluded)
		}
	})

	t.Run("returns error when auth header format is malformed", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "Bearer test-key")

		_, err := GetAPIKey(headers)
		if err == nil {
			t.Fatal("GetAPIKey() error = nil, want non-nil")
		}
	})
}