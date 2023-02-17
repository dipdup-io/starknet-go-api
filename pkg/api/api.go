package api

import (
	"context"
	"net/http"
	"sync/atomic"
	"time"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/time/rate"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// API - wrapper of starknet node API.
type API struct {
	client    *http.Client
	baseURL   string
	id        *atomic.Uint64
	rateLimit *rate.Limiter
}

// NewAPI - constructor of API
func NewAPI(baseURL string, opts ...ApiOption) API {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	client := &http.Client{
		Transport: t,
	}
	api := API{
		client:  client,
		baseURL: baseURL,
		id:      new(atomic.Uint64),
	}

	for i := range opts {
		opts[i](&api)
	}

	return api
}

func (api API) prepareRequest(ctx context.Context, method string, params []any, opts ...RequestOption) *Request {
	req := Request{
		Version: defaultJSONRPC,
		ID:      api.id.Add(1),
		Method:  method,
		Params:  params,

		timeout: 10 * time.Second,
	}

	for i := range opts {
		opts[i](&req)
	}

	return &req
}
