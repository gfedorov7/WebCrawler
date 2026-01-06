package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Fetch(url string, timeout time.Duration) ([]byte, error) {
	client := &http.Client{Timeout: timeout}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "SimpleCrawler/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
