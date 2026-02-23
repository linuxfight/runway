package version

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const cacheFileName = ".crew-update-check"

type cacheData struct {
	CheckedAt time.Time `json:"checked_at"`
	Latest    string    `json:"latest"`
}

func cachePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, cacheFileName), nil
}

func readCache() (*cacheData, error) {
	path, err := cachePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c cacheData
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

func writeCache(latest string) {
	path, err := cachePath()
	if err != nil {
		return
	}

	c := cacheData{
		CheckedAt: time.Now(),
		Latest:    latest,
	}

	data, _ := json.Marshal(c)
	_ = os.WriteFile(path, data, 0644)
}
