package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from the headers of am HTTP request
// example:
// Authorization: ApiKey {key}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no info was found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid authorization type")
	}

	return vals[1], nil
}
