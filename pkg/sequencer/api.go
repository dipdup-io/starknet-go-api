package sequencer

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/time/rate"
)

// API - wrapper of starknet node API.
type API struct {
	client           *http.Client
	gatewayUrl       string
	feederGatewayUrl string
	cacheDir         string
	rateLimit        *rate.Limiter
	rps              int
}

// NewAPI - constructor of API
func NewAPI(gatewayUrl, feederGatewayUrl string, opts ...ApiOption) API {
	api := API{
		gatewayUrl:       gatewayUrl,
		feederGatewayUrl: feederGatewayUrl,
	}

	for i := range opts {
		opts[i](&api)
	}

	t := http.DefaultTransport.(*http.Transport).Clone()
	if api.rps < 1 || api.rps > 100 {
		t.MaxIdleConns = 10
		t.MaxConnsPerHost = 10
		t.MaxIdleConnsPerHost = 10
	} else {
		t.MaxIdleConns = api.rps
		t.MaxConnsPerHost = api.rps
		t.MaxIdleConnsPerHost = api.rps
	}

	api.client = &http.Client{
		Transport: t,
	}

	return api
}

func (api API) getFromFeederGateway(ctx context.Context, path, cacheFileName string, args map[string]string, output any) error {
	return api.get(ctx, api.feederGatewayUrl, path, cacheFileName, args, output)
}

func (api API) get(ctx context.Context, baseURL, path, cacheFileName string, args map[string]string, output any) error {
	if api.cacheDir != "" && cacheFileName != "" {
		if err := api.readJSONFromCache(path, cacheFileName, output); err == nil {
			return nil
		}
	}

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
		return errors.Wrapf(ErrRequest, "invalid status code: %d %s", response.StatusCode, string(body))
	}

	if api.cacheDir != "" && cacheFileName != "" {
		return api.parseJSONWithCache(response.Body, path, cacheFileName, output)
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

	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, response.Body); err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		var e Error
		if err := json.NewDecoder(buffer).Decode(&e); err != nil {
			return errors.Wrap(ErrRequest, err.Error())
		}
		return errors.Wrap(ErrRequest, e.Error())
	}

	err = json.NewDecoder(buffer).Decode(output)
	return err
}
