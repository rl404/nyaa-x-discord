package internal

import (
	"net/url"
	"strings"
)

func keyValueToMessage(keyValue []keyValue) (msg string) {
	for _, kv := range keyValue {
		msg += kv.Key + " : " + kv.Value + "\n"
	}
	return msg
}

func getValueFromKey(keyValue []keyValue, key string) string {
	for _, kv := range keyValue {
		if kv.Key == key {
			return kv.Value
		}
	}
	return ""
}

func getNyaaQuery(user User, rss ...bool) string {
	nyaa := nyaaURL + "?"
	nyaa += "f=" + user.Filter
	nyaa += "&c=" + user.Category
	for i := range user.Queries {
		user.Queries[i] = "(" + user.Queries[i] + ")"
	}
	nyaa += "&q=" + url.QueryEscape(strings.Join(user.Queries, "|"))
	if len(rss) > 0 && rss[0] {
		nyaa += "&page=rss"
	}
	return nyaa
}

func inArray(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func ellipsis(str string, length int) string {
	if len(str) > length {
		return str[:length] + "..."
	}
	return str
}
