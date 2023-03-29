package sequencer

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/goccy/go-json"
)

func (api API) parseJSONWithCache(body io.ReadCloser, path, cacheFileName string, output any) error {
	if api.cacheDir == "" {
		return errors.New("empty cached directory")
	}
	if cacheFileName == "" {
		return errors.New("empty cache file name")
	}

	dirPath := filepath.Join(api.cacheDir, path)
	if err := os.MkdirAll(dirPath, 0777); err != nil {
		return err
	}
	fullPath := filepath.Join(dirPath, cacheFileName)
	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(io.TeeReader(body, f)).Decode(output)
	return err
}

func (api API) readJSONFromCache(path, cacheFileName string, output any) error {
	if api.cacheDir == "" {
		return errors.New("empty cached directory")
	}
	if cacheFileName == "" {
		return errors.New("empty cache file name")
	}

	fullPath := filepath.Join(api.cacheDir, path, cacheFileName)
	f, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(output)
	return err
}
