package main

import (
	"WebCrawler/cmd/internal/fetcher"
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
	fmt.Println(hrefs)
}
