package helper

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// CheckError checks for error and panics
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// ExtractURLID returns id part of URL
func ExtractURLID(r *http.Request, endpoint string) string {
	return strings.TrimPrefix(r.URL.Path, fmt.Sprintf("/v1/api/%s/", endpoint))
}

// GeneratePasswordHash returns hash from password
func GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}
