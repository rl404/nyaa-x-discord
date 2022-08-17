package entity

// Message constant value.
const (
	BlueColor  int    = 4886754
	FooterName string = "nyaa-x-discord"

	IntroTitle         string = "Welcome"
	IntroContent       string = "\nIt looks like this is your first time using this bot.\n\nThis bot will help you keeping track of [Nyaa](%s) update according to your query/filter.\n\n**How to Start**\n1. Set filter.\n2. Set category.\n3. Set query.\n4. Turn on subscription.\n5. Wait a bit and I will notify you if there is a new update.\n\n**{{prefix}}help** to see all command list."
	HelpCmd            string = "{{prefix}}help"
	HelpContent        string = "Show all command list."
	FilterCmd          string = "{{prefix}}filter"
	FilterContent      string = "Get filter names and their id."
	FilterSetCmd       string = "{{prefix}}filter set <filter_id>"
	FilterSetContent   string = "Set filter type for your query. `filter_id` is from `{{prefix}}filter`."
	CategCmd           string = "{{prefix}}category"
	CategContent       string = "Get category names and their id."
	CategSetCmd        string = "{{prefix}}category set <category_id>"
	CategSetContent    string = "Set category type for your query. `category_id` is from `{{prefix}}category`."
	QueryCmd           string = "{{prefix}}query"
	QueryContent       string = "Get your query string list and their id."
	QueryAddCmd        string = "{{prefix}}query add <string> [string...]"
	QueryAddContent    string = "Add your query string. You can have more than 1 query. For example:\n\nContain `naruto` ```{{prefix}}query add naruto```\nContain `one`, `piece` and `horriblesubs` ```{{prefix}}query add one piece horriblesubs```\nContain `bleach` but not `720p` ```{{prefix}}query add bleach -720p ```"
	QueryDeleteCmd     string = "{{prefix}}query delete <query_id> [query_id...]"
	QueryDeleteContent string = "Delete 1 or more your query strings. `query_id` is from `{{prefix}}query`."
	SubsCmd            string = "{{prefix}}subscribe"
	SubsContent        string = "Get your subscription status."
	SubsSetCmd         string = "{{prefix}}subscribe <on|off>"
	SubsSetContent     string = "Turn on or off bot subscription."
	InvalidCmd         string = "Invalid command. See **{{prefix}}help** for more information."
)
