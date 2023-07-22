package entity

// NyaaURL is nyaa domain.
const NyaaURL string = "https://nyaa.si/"

// Filters is nyaa filters.
var Filters KeyValues = KeyValues{
	{Key: "0", Value: "No filter"},
	{Key: "1", Value: "No remakes"},
	{Key: "2", Value: "Trusted only"},
}

// Categories is nyaa categories.
var Categories KeyValues = KeyValues{
	{Key: "0_0", Value: "All categories"},
	{Key: "1_0", Value: "Anime"},
	{Key: "1_1", Value: "Anime Music Video"},
	{Key: "1_2", Value: "English-translated"},
	{Key: "1_3", Value: "Non-English-translated"},
	{Key: "1_4", Value: "Raw"},
	{Key: "2_0", Value: "Audio"},
	{Key: "2_1", Value: "Lossless"},
	{Key: "2_2", Value: "Lossy"},
	{Key: "3_0", Value: "Literature"},
	{Key: "3_1", Value: "English-translated"},
	{Key: "3_2", Value: "Non-English-translated"},
	{Key: "3_3", Value: "Raw"},
	{Key: "4_0", Value: "Live Action"},
	{Key: "4_1", Value: "English-translated"},
	{Key: "4_2", Value: "Idol/Promotional Video"},
	{Key: "4_3", Value: "Non-English-translated"},
	{Key: "4_4", Value: "Raw"},
	{Key: "5_0", Value: "Pictures"},
	{Key: "5_1", Value: "Graphics"},
	{Key: "5_2", Value: "Photos"},
	{Key: "6_0", Value: "Software"},
	{Key: "6_1", Value: "Applications"},
	{Key: "6_2", Value: "Games"},
}

// KeyValues is array of key values.
type KeyValues []keyValue

type keyValue struct {
	Key   string
	Value string
}

// GetValueByKey to get value by key.
func (kv KeyValues) GetValueByKey(key string) string {
	for _, v := range kv {
		if v.Key == key {
			return v.Value
		}
	}
	return ""
}
