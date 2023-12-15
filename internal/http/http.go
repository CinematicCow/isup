package http

import (
	"fmt"
	"net/http"
)

func Get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %s", err)
	}

	return resp, nil
}
