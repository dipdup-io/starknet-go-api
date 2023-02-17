package api

import (
	"time"

	"golang.org/x/time/rate"
)

// RequestOption -
type RequestOption func(*Request)

// WithCustomID - set custom request id. Request id increment automatically per request.
func WithCustomID(id uint64) RequestOption {
	return func(req *Request) {
		req.ID = id
	}
}

// WithJsonRpcVersion - change json rpc version. Default: 2.0
func WithJsonRpcVersion(version string) RequestOption {
	return func(req *Request) {
		req.Version = version
	}
}

// WithTimeout - set custom timeout in seconds. Default: 10 seconds
func WithTimeout(timeout uint64) RequestOption {
	return func(req *Request) {
		if timeout > 0 {
			req.timeout = time.Second * time.Duration(timeout)
		}
	}
}

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
