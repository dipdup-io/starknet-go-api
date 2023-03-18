package sequencer

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/time/rate"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// API - wrapper of starknet node API.
type API struct {
	client           *http.Client
	gatewayUrl       string
	feederGatewayUrl string
	rateLimit        *rate.Limiter
}

// NewAPI - constructor of API
func NewAPI(gatewayUrl, feederGatewayUrl string, opts ...ApiOption) API {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	client := &http.Client{
		Transport: t,
	}
	api := API{
		client:           client,
		gatewayUrl:       gatewayUrl,
		feederGatewayUrl: feederGatewayUrl,
	}

	for i := range opts {
		opts[i](&api)
	}

	return api
}

func (api API) getFromFeederGateway(ctx context.Context, path string, args map[string]string, output any) error {
	return api.get(ctx, api.feederGatewayUrl, path, args, output)
}

func (api API) get(ctx context.Context, baseURL, path string, args map[string]string, output any) error {
	u, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	u.Path, err = url.JoinPath(u.Path, path)
	if err != nil {
		return err
	}

	values := u.Query()
	for key, value := range args {
		values.Add(key, value)
	}
	u.RawQuery = values.Encode()

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return err
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	start := time.Now()
	response, err := api.client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	log.Trace().Msgf("[%d ms] %s", time.Since(start).Milliseconds(), u.String())

	if response.StatusCode != http.StatusOK {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return errors.Errorf("invalid status code: %d %s", response.StatusCode, string(body))
	}

	err = json.NewDecoder(response.Body).Decode(output)
	return err
}

func (api API) postToFeederGateway(ctx context.Context, path string, args map[string]string, body any, output any) error {
	return api.post(ctx, api.feederGatewayUrl, path, args, body, output)
}

func (api API) postToGateway(ctx context.Context, path string, args map[string]string, body any, output any) error {
	return api.post(ctx, api.gatewayUrl, path, args, body, output)
}

func (api API) post(ctx context.Context, baseURL, path string, args map[string]string, body any, output any) error {
	u, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	u.Path, err = url.JoinPath(u.Path, path)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return err
	}

	values := u.Query()
	for key, value := range args {
		values.Add(key, value)
	}
	u.RawQuery = values.Encode()

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return err
		}
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), buf)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := api.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var e data.Error
		if err := json.NewDecoder(io.TeeReader(response.Body, os.Stdout)).Decode(&e); err != nil {
			return err
		}
		return errors.Errorf("invalid status code: %d %s", response.StatusCode, e.Message)
	}

	err = json.NewDecoder(response.Body).Decode(output)
	return err
}
