package http

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const tokenScheme = "Token "

// errors
var (
	ErrAuthHeaderMissing = errors.New("authorization Header is missing")
	ErrAuthBadScheme     = errors.New("authorization Header Scheme is invalid")
)

// GetToken will parse the token from http Authorization Header.
func GetToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", ErrAuthHeaderMissing
	}
	if !strings.HasPrefix(header, tokenScheme) {
		return "", ErrAuthBadScheme
	}
	return header[len(tokenScheme):], nil
}

// SetToken adds the token to the request.
func SetToken(token string, req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("%s%s", tokenScheme, token))
}
