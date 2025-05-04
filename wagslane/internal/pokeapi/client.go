package pokeapi

import (
	"net/http"
	"time"

	"github.com/GitJohnFis/Open-source-SmartBoard-/internal/pokecache"
)

// Client -
type Client struct {
	cache 	   pokecache.Cache
	httpClient http.Client
	// NextURL    *string
	// PreviousURL *string
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		// NextURL:    nil,
		// PreviousURL: nil,
	}
}
