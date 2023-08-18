package api

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"
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

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return err
		}
	}

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
		return errors.Wrapf(ErrRequest, "request %d invalid status code: %d", output.ID, response.StatusCode)
	}

	if err := json.NewDecoder(buffer).Decode(output); err != nil {
		return err
	}

	if output.Error != nil {
		return errors.Wrapf(ErrRequest, "request %d error: %s", output.ID, output.Error.Error())
	}

	return nil
}
