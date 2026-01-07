package main

import (
	"WebCrawler/cmd/internal/fetcher"
	"WebCrawler/cmd/internal/normalizer"
	"WebCrawler/cmd/internal/parser"
	"WebCrawler/cmd/internal/worker"
	"fmt"
	"sync"
)

func main() {
	var consoleFlag parser.ConsoleFlag
	err := parser.ParseFlag(&consoleFlag)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := fetcher.Fetch(consoleFlag.Url, consoleFlag.Timeout)
	if result.Error != nil {
		fmt.Println(result.Error, result.Status)
		return
	}

	hrefs := parser.HtmlHrefParser(result.Body)
	validHrefs := []string{consoleFlag.Url}
	for _, href := range hrefs {
		v := normalizer.ParseUrl(consoleFlag.Depth, consoleFlag.Url, href)
		if v == "" {
			continue
		}
		validHrefs = append(validHrefs, v)
	}

	jobs := make(chan string)
	results := make(chan worker.Stats)

	safer := &worker.SafeURLCollection{}

	var wg sync.WaitGroup
	var resultWg sync.WaitGroup

	resultWg.Add(1)
	go func() {
		defer resultWg.Done()
		worker.HandleResults(results)
	}()

	for i := 0; i < consoleFlag.Workers; i++ {
		wg.Add(1)
		go worker.StatsCollector(safer, jobs, results, &wg, &consoleFlag)
	}

	go func() {
		for _, url := range validHrefs {
			jobs <- url
		}
		close(jobs)
	}()

	wg.Wait()
	close(results)
	resultWg.Wait()
}
