package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

// discord is discord client.
var discord *discordgo.Session

// prefix is discord message prefix that the bot will read.
const prefix = '!'

// initDiscord to initialize new discord client.
func initDiscord() (err error) {
	discord, err = discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return err
	}

	// Handle received message.
	discord.AddHandler(messageHandler)

	// Login discord bot.
	err = discord.Open()
	if err != nil {
		return err
	}

	fmt.Println("discord bot is running...")
	return nil
}

// messageHandler will handle received message.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Prefix check.
	if m.Content[0] != prefix {
		return
	}

	// Remove prefix.
	m.Content = m.Content[1:]

	// Handle first time.
	user, exist, err := getUser(m.Author.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !exist {
		handleFirstTime(s, m)
		return
	}

	// Get arguments.
	r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`)
	msgSplit := r.FindAllString(m.Content, -1)

	// If the message is "ping" reply with "Pong!"
	if msgSplit[0] == "ping" {
		s.ChannelMessageSend(user.ChannelID, "Pong!")
		return
	}

	// Show all command list.
	if msgSplit[0] == "help" {
		handleHelp(s, user)
		return
	}

	// Filter command.
	if msgSplit[0] == "filter" {
		handleFilter(s, user, msgSplit)
		return
	}

	// Category command.
	if msgSplit[0] == "category" {
		handleCategory(s, user, msgSplit)
		return
	}

	// Query command.
	if msgSplit[0] == "query" {
		handleQuery(s, user, msgSplit)
		return
	}

	// Subscribe command.
	if msgSplit[0] == "subscribe" {
		handleSubs(s, user, msgSplit)
		return
	}
}

// sendFeed to request discord bot to send new feeds.
func sendFeed(feeds []Feed, user User) {
	msg := discordgo.MessageEmbed{
		Color: blueColor,
	}

	// Discord limit 25 fields per message.
	for len(feeds) > 0 {
		var fields []*discordgo.MessageEmbedField

		limit := 25
		if len(feeds)-25 < 25 {
			limit = len(feeds)
		}

		for _, feed := range feeds[:limit] {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name: ellipsis(feed.Title, 100),
				Value: fmt.Sprintf("[link](%s)  •  `%s`  •  %s",
					feed.Link,
					feed.Size,
					feed.Date.Format("2006-01-02 15:04:05")),
			})
		}

		msg.Fields = fields
		feeds = feeds[limit:]

		_, err := discord.ChannelMessageSendEmbed(user.ChannelID, &msg)
		if err != nil {
			fmt.Println(user.UserID, err)
			return
		}
	}
}

// handleFirstTime to handle first time user message.
func handleFirstTime(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Get DM channel.
	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create new user.
	err = createUser(m.Author.ID, c.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := discordgo.MessageEmbed{
		Title:       introTitle + " " + m.Author.Username,
		Description: introContent,
		Color:       blueColor,
	}

	_, err = s.ChannelMessageSendEmbed(c.ID, &msg)
	if err != nil {
		fmt.Println(err)
	}
}

// handleHelp to show all command list.
func handleHelp(s *discordgo.Session, user User) {
	msg := discordgo.MessageEmbed{
		Color: blueColor,
		Description: fmt.Sprintf("**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s",
			helpCmd, helpContent,
			filterCmd, filterContent,
			filterSetCmd, filterSetContent,
			categCmd, categContent,
			categSetCmd, categSetContent,
			queryCmd, queryContent,
			queryAddCmd, queryAddContent,
			queryDeleteCmd, queryDeleteContent,
			subsCmd, subsContent,
			subsSetCmd, subsSetContent),
	}

	_, err := s.ChannelMessageSendEmbed(user.ChannelID, &msg)
	if err != nil {
		fmt.Println(err)
	}
}

// handleFilter to handle filter commands.
func handleFilter(s *discordgo.Session, user User, args []string) {
	msg := discordgo.MessageEmbed{
		Color: blueColor,
	}

	// Show filter list.
	if len(args) == 1 {
		msg.Description = fmt.Sprintf("**Options**```%s```\nYour current filter: **%s - %s**\nYour filtered Nyaa will look like [this](%s).",
			keyValueToMessage(filters),
			user.Filter,
			getValueFromKey(filters, user.Filter),
			getNyaaQuery(user),
		)

		_, err := s.ChannelMessageSendEmbed(user.ChannelID, &msg)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Invalid command.
	if len(args) != 3 || args[1] != "set" || getValueFromKey(filters, args[2]) == "" {
		handleInvalidCmd(s, user)
		return
	}

	user.Filter = args[2]
	_, err := dbCollection.UpdateOne(dbCtx, bson.M{"_id": user.ID}, bson.D{{Key: "$set", Value: user}})
	if err != nil {
		fmt.Println(err)
		return
	}

	handleSuccess(s, user)
}

// handleCategory to handle category commands.
func handleCategory(s *discordgo.Session, user User, args []string) {
	msg := discordgo.MessageEmbed{
		Color: blueColor,
	}

	// Show category list.
	if len(args) == 1 {
		msg.Description = fmt.Sprintf("**Options**```%s```\nYour current category: **%s - %s**\nYour filtered Nyaa will look like [this](%s).",
			keyValueToMessage(categories),
			user.Category,
			getValueFromKey(categories, user.Category),
			getNyaaQuery(user),
		)

		_, err := s.ChannelMessageSendEmbed(user.ChannelID, &msg)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Invalid command.
	if len(args) != 3 || args[1] != "set" || getValueFromKey(categories, args[2]) == "" {
		handleInvalidCmd(s, user)
		return
	}

	user.Category = args[2]
	_, err := dbCollection.UpdateOne(dbCtx, bson.M{"_id": user.ID}, bson.D{{Key: "$set", Value: user}})
	if err != nil {
		fmt.Println(err)
		return
	}

	handleSuccess(s, user)
}

// handleQuery to handle query commands.
func handleQuery(s *discordgo.Session, user User, args []string) {
	msg := discordgo.MessageEmbed{
		Color: blueColor,
	}

	// Show category list.
	if len(args) == 1 {
		if len(user.Queries) == 0 {
			msg.Description = "**Queries**```Empty query. Go create a new one.```"
		} else {
			msg.Description = "**Queries**```"
			for i, q := range user.Queries {
				msg.Description += fmt.Sprintf("%v : %s", i, q)
				if i < len(user.Queries)-1 {
					msg.Description += "\n"
				}
			}
			msg.Description += "```\nYour filtered Nyaa will look like [this](" + getNyaaQuery(user) + ")."
		}

		_, err := s.ChannelMessageSendEmbed(user.ChannelID, &msg)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Invalid command.
	if len(args) < 3 || (args[1] != "add" && args[1] != "delete") {
		handleInvalidCmd(s, user)
		return
	}

	if args[1] == "add" {
		user.Queries = append(user.Queries, strings.Join(args[2:], " "))
		_, err := dbCollection.UpdateOne(dbCtx, bson.M{"_id": user.ID}, bson.D{{Key: "$set", Value: user}})
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		newQuery := []string{}
		for i, q := range user.Queries {
			if !inArray(args[2:], strconv.Itoa(i)) {
				newQuery = append(newQuery, q)
			}
		}

		user.Queries = newQuery
		_, err := dbCollection.UpdateOne(dbCtx, bson.M{"_id": user.ID}, bson.D{{Key: "$set", Value: user}})
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	handleSuccess(s, user)
}

// handleSubs to handle subscription.
func handleSubs(s *discordgo.Session, user User, args []string) {
	msg := discordgo.MessageEmbed{
		Color: blueColor,
	}

	// Show subscription status.
	if len(args) == 1 {
		status := "off"
		if user.Subscribe {
			status = "on"
		}

		msg.Description = fmt.Sprintf("**Status**\nYour current subscription: **%s**\nYour filtered Nyaa will look like [this](%s).",
			status,
			getNyaaQuery(user),
		)

		_, err := s.ChannelMessageSendEmbed(user.ChannelID, &msg)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Invalid command.
	if len(args) != 2 || (args[1] != "on" && args[1] != "off") {
		handleInvalidCmd(s, user)
		return
	}

	user.Subscribe = args[1] == "on"
	_, err := dbCollection.UpdateOne(dbCtx, bson.M{"_id": user.ID}, bson.D{{Key: "$set", Value: user}})
	if err != nil {
		fmt.Println(err)
		return
	}

	handleSuccess(s, user)
}

// getNyaaQuery to generate full Nyaa url with params.
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

// handleInvalidCmd to handle invalid command.
func handleInvalidCmd(s *discordgo.Session, user User) {
	s.ChannelMessageSend(user.ChannelID, invalidCmd)
}

// handleSuccess to notify updated successfully.
func handleSuccess(s *discordgo.Session, user User) {
	msg := discordgo.MessageEmbed{
		Color:       blueColor,
		Title:       "Updated",
		Description: "Your filtered Nyaa will look like [this](" + getNyaaQuery(user) + ").",
	}
	_, err := s.ChannelMessageSendEmbed(user.ChannelID, &msg)
	if err != nil {
		fmt.Println(err)
	}
}
