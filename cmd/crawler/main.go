package main

import (
	"WebCrawler/cmd/internal/fetcher"
	"WebCrawler/cmd/internal/normalizer"
	"WebCrawler/cmd/internal/parser"
	"fmt"
)

func main() {
	var consoleFlag parser.ConsoleFlag
	err := parser.ParseFlag(&consoleFlag)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := fetcher.Fetch(consoleFlag.Url, consoleFlag.Timeout)
	if err != nil {
		fmt.Println(err)
		return
	}

	hrefs := parser.HtmlHrefParser(body)
	var validHrefs []string
	for _, href := range hrefs {
		v := normalizer.ParseUrl(consoleFlag.Depth, consoleFlag.Url, href)
		if v == "" {
			continue
		}
		validHrefs = append(validHrefs, v)
	}

}
