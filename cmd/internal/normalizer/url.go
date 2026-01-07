package normalizer

import (
	"strings"
)

const countProtocolSlash = 2

func ParseUrl(depth int, startUrl, body string) (url string) {
	if depth == 0 {
		return startUrl
	}

	if strings.TrimSpace(body) == "" || strings.Count(body, "#") > 0 {
		return ""
	}

	if !strings.HasPrefix(body, "http") {
		li := strings.LastIndex(startUrl, "/")
		if li != len(startUrl)-1 {
			startUrl += "/"
		}
		if body[0] == '/' {
			body = body[1:]
		}

		url = startUrl + body
	} else {
		url = body
	}

	li := strings.LastIndex(url, "/")
	if li == len(url)-1 {
		url = url[:li]
	}

	if CountDepth(url) > depth {
		return ""
	}
	return
}

func CountDepth(url string) int {
	return strings.Count(url, "/") - countProtocolSlash
}
