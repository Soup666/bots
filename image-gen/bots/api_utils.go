package bots

import (
	"fmt"
	"net/http"
)

type APIUtils struct{}

func NewAPIUtils() *APIUtils {
	return &APIUtils{}
}

// APIUtils provides utility functions for API operations
func (au *APIUtils) ValidateResponse(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}
	return nil
}
