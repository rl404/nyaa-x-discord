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
	for i := range queries {
		queries[i] = "(" + queries[i] + ")"
	}
	nyaa += "&q=" + url.QueryEscape(strings.Join(queries, "|"))
	if len(isRSS) > 0 && isRSS[0] {
		nyaa += "&page=rss"
	}
	return nyaa
}
