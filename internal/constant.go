package internal

var (
	blueColor  = 4886754
	footerName = "nyaa-x-discord"
	nyaaURL    = "https://nyaa.si/"

	introTitle         = "Welcome"
	introContent       = "\nIt looks like this is your first time using this bot.\n\nThis bot will help you keeping track of [Nyaa](" + nyaaURL + ") update according to your query/filter.\n\n**How to Start**\n1. Set filter.\n2. Set category.\n3. Set query.\n4. Turn on subscription.\n5. Wait a bit and I will notify you if there is a new update.\n\n**!help** to see all command list."
	helpCmd            = "!help"
	helpContent        = "Show all command list."
	filterCmd          = "!filter"
	filterContent      = "Get filter names and their id."
	filterSetCmd       = "!filter set <filter_id>"
	filterSetContent   = "Set filter type for your query. `filter_id` is from `!filter`."
	categCmd           = "!category"
	categContent       = "Get category names and their id."
	categSetCmd        = "!category set <category_id>"
	categSetContent    = "Set category type for your query. `category_id` is from `!category`."
	queryCmd           = "!query"
	queryContent       = "Get your query string list and their id."
	queryAddCmd        = "!query add <string> [string...]"
	queryAddContent    = "Add your query string. You can have more than 1 query. For example:\n\nContain `naruto` ```!query add naruto```\nContain `one`, `piece` and `horriblesubs` ```!query add one piece horriblesubs```\nContain `bleach` but not `720p` ```!query add bleach -720p ```"
	queryDeleteCmd     = "!query delete <query_id> [query_id...]"
	queryDeleteContent = "Delete 1 or more your query strings. `query_id` is from `!query`."
	subsCmd            = "!subscribe"
	subsContent        = "Get your subscription status."
	subsSetCmd         = "!subscribe <on|off>"
	subsSetContent     = "Turn on or off bot subscription."
	invalidCmd         = "Invalid command. See **!help** for more information."
)

type keyValue struct {
	Key   string
	Value string
}

var filters = []keyValue{
	{Key: "0", Value: "No filter"},
	{Key: "1", Value: "No remakes"},
	{Key: "2", Value: "Trusted only"},
}

var categories = []keyValue{
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
