package api

import (
	"bytes"
	"context"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

func post[T any](ctx context.Context, client *http.Client, url string, req Request, output *Response[T]) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return err
	}

	requestCtx, cancel := context.WithTimeout(ctx, req.timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(requestCtx, http.MethodPost, url, buf)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.Errorf("invalid status code: %d", response.StatusCode)
	}

	if err := json.NewDecoder(response.Body).Decode(output); err != nil {
		return err
	}

	if output.Error != nil {
		return errors.Errorf("request %d error: %s (code %d)", output.ID, output.Error.Message, output.Error.Code)
	}

	return nil
}
