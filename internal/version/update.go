package version

import (
	"encoding/json"
	"net/http"
	"time"
)

const repoAPI = "https://api.github.com/repos/cryingcatscloud/runway/releases/latest"

type release struct {
	TagName string `json:"tag_name"`
}

func CheckForUpdate(current string) (string, error) {
	if current == "" || current == "dev" {
		return "", nil
	}

	client := http.Client{
		Timeout: 800 * time.Millisecond,
	}

	resp, err := client.Get(repoAPI)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var r release
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", err
	}

	if r.TagName == "" || r.TagName == current {
		return "", nil
	}

	return r.TagName, nil
}

func CheckWithCache(current string) (string, error) {
	if current == "" || current == "dev" {
		return "", nil
	}

	c, err := readCache()
	if err == nil {
		if c.Latest != "" && c.Latest != current {
			return c.Latest, nil
		}

		if time.Since(c.CheckedAt) < 24*time.Hour {
			return "", nil
		}
	}

	latest, err := CheckForUpdate(current)
	if err != nil {
		return "", err
	}

	writeCache(latest)

	return latest, nil
}
