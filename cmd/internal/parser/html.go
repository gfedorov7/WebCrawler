package parser

import (
	"bytes"
	"golang.org/x/net/html"
)

func HtmlHrefParser(body []byte) (hrefs []string) {
	tokenizer := html.NewTokenizer(bytes.NewReader(body))

	hrefs = []string{}
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}

		if tokenType != html.StartTagToken && tokenType != html.SelfClosingTagToken {
			continue
		}

		token := tokenizer.Token()
		if token.Data != "a" {
			continue
		}

		for _, attr := range token.Attr {
			if attr.Key == "href" && attr.Val != "" {
				hrefs = append(hrefs, attr.Val)
			}
		}
	}

	return
}
