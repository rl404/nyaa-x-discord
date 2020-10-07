package main

// ellipsis to truncate long string.
func ellipsis(str string, length int) string {
	if len(str) > length {
		return str[:length] + "..."
	}
	return str
}

// keyValueToMessage to convert key value model to string.
func keyValueToMessage(keyValue []KeyValue) (msg string) {
	for _, kv := range keyValue {
		msg += kv.Key + " : " + kv.Value + "\n"
	}
	return msg
}

// getValueFromKey to get value with key.
func getValueFromKey(keyValue []KeyValue, key string) string {
	for _, kv := range keyValue {
		if kv.Key == key {
			return kv.Value
		}
	}
	return ""
}

// inArray to check if value is in array.
func inArray(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}