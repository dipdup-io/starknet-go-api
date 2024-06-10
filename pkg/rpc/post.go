package api

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func post[T any](ctx context.Context, api API, req Request, output *Response[T]) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, api.baseURL, buf)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	if api.apiKey != "" && api.headerApiKey != "" {
		request.Header.Add(api.headerApiKey, api.apiKey)
	}

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return err
		}
	}

	response, err := api.client.Do(request)
	if err != nil {
		return err
	}
	defer closeWithLogError(response.Body)

	if response.StatusCode != http.StatusOK {
		return errors.Wrapf(ErrRequest, "request %d invalid status code: %d", output.ID, response.StatusCode)
	}

	if err := json.NewDecoder(response.Body).Decode(output); err != nil {
		return err
	}

	if output.Error != nil {
		return output.Error
	}

	return nil
}

func closeWithLogError(stream io.ReadCloser) {
	if _, err := io.Copy(io.Discard, stream); err != nil {
		log.Err(err).Msg("api copy body response to discard")
	}
	if err := stream.Close(); err != nil {
		log.Err(err).Msg("api close body request")
	}
}
