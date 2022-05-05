package helper

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/mail"
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
	return strings.TrimPrefix(r.URL.Path, fmt.Sprintf("/v1/api/%s/", endpoint))
}

// IsValidEmail validates email
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// GeneratePasswordHash returns hash from password
func GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
