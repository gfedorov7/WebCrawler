package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Result struct {
	Status int
	Url    string
	Error  error
	Body   []byte
}

func (r *Result) String() string {
	return fmt.Sprintf("Status: %v, Url: %v, Error: %v", r.Status, r.Url, r.Error)
}

func Fetch(url string, timeout time.Duration) Result {
	client := &http.Client{Timeout: timeout}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Result{500, url, err, nil}
	}
	req.Header.Set("User-Agent", "SimpleCrawler/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return Result{400, url, err, nil}
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return Result{resp.StatusCode, url,
			fmt.Errorf("status code: %d", resp.StatusCode), nil}
	}

	body, err := io.ReadAll(resp.Body)
	return Result{resp.StatusCode, url, err, body}
}
