package sequencer

import (
	"time"

	"golang.org/x/time/rate"
)

// ApiOption -
type ApiOption func(*API)

// WithRateLimit -
func WithRateLimit(requestPerSecond int) ApiOption {
	return func(api *API) {
		if requestPerSecond > 0 {
			api.rateLimit = rate.NewLimiter(rate.Every(time.Second/time.Duration(requestPerSecond)), requestPerSecond)
		}
	}
}

// WithCacheInFS -
func WithCacheInFS(dir string) ApiOption {
	return func(api *API) {
		if dir != "" {
			api.cacheDir = dir
		}
	}
}
