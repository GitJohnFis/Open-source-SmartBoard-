package pokeapi

import (
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	// NextURL    *string
	// PreviousURL *string
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		// NextURL:    nil,
		// PreviousURL: nil,
	}
}
