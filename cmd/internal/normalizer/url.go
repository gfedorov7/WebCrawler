package normalizer

import "strings"

const countProtocolSlash = 2

func ParseUrl(depth int, startUrl, body string) (url string) {
	if strings.Trim(body, " ") == "" || strings.Count(body, "#") > 0 {
		return ""
	}

	li := strings.LastIndex(body, "/")
	if li != len(body)-1 {
		body += "/"
	}

	if strings.HasPrefix(body, "http") {
		countSlash := strings.Count(body, "/")
		if countSlash-countProtocolSlash > depth {
			return ""
		}
		return body
	}

	li = strings.LastIndex(startUrl, "/")
	if li != len(startUrl)-1 {
		startUrl += "/"
	}
	i := strings.Index(body, "/")
	if i == 0 {
		body = body[1:]
	}

	return startUrl + body
}
