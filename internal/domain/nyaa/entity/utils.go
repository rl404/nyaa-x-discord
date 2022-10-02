package entity

import (
	"net/url"
	"strings"
)

// GenerateURL to generate nyaa url.
func GenerateURL(filter, category string, queries []string, isRSS ...bool) string {
	nyaa := NyaaURL + "?"
	nyaa += "f=" + filter
	nyaa += "&c=" + category

	tmp := make([]string, len(queries))
	for i := range queries {
		tmp[i] = "(" + queries[i] + ")"
	}

	nyaa += "&q=" + url.QueryEscape(strings.Join(tmp, "|"))
	if len(isRSS) > 0 && isRSS[0] {
		nyaa += "&page=rss"
	}

	return nyaa
}
