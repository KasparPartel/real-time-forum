package helper

import (
	"fmt"
	"net/http"
	"strings"
)

// CheckError checks for error and panics
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// ExtractURLID returns id part of URL
func ExtractURLID(r *http.Request, endpoint string) string {
	return strings.TrimPrefix(r.URL.Path, fmt.Sprintf("/api/%s/", endpoint))
}
