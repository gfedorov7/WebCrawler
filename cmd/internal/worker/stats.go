package worker

import (
	"WebCrawler/cmd/internal/fetcher"
	"WebCrawler/cmd/internal/parser"
	"fmt"
	"sync"
)

type Stats struct {
	Url     string
	Success bool
	Bad     bool
	Error   error
}

func StatsCollector(
	safer *SafeURLCollection,
	jobs <-chan string,
	results chan<- Stats,
	wg *sync.WaitGroup,
	c *parser.ConsoleFlag,
) {
	defer wg.Done()

	for url := range jobs {
		if !safer.AddIfNotExists(url) {
			continue
		}
		result := fetcher.Fetch(url, c.Timeout)
		if result.Status >= 200 && result.Status <= 299 {
			results <- Stats{Url: url, Success: true, Bad: false, Error: nil}
		} else {
			results <- Stats{Url: url, Success: false, Bad: true, Error: result.Error}
		}
	}
}

func HandleResults(results <-chan Stats) {
	for result := range results {
		fmt.Println("-------------------")
		fmt.Println(result.Url)
		if result.Bad {
			fmt.Printf("[!] %s\n", result.Error)
		} else {
			fmt.Printf("[+] %v\n", result.Success)
		}
	}
}
