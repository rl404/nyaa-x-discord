package internal

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Handler contains all handler function.
type Handler interface {
	Handler() func(*discordgo.Session, *discordgo.MessageCreate)
	PrefixCheck(cmd string) bool
	CleanPrefix(cmd string) string
	HandleFirstTime(session *discordgo.Session, msg *discordgo.MessageCreate) error
	HandlePing(session *discordgo.Session, channelID string) error
	HandleHelp(session *discordgo.Session, user User) error
	HandleFilter(session *discordgo.Session, user User, args []string) error
	HandleCategory(session *discordgo.Session, user User, args []string) error
	HandleQuery(session *discordgo.Session, user User, args []string) error
	HandleSubscribe(session *discordgo.Session, user User, args []string) error
}

type handler struct {
	db     Database
	prefix string
}

// NewHandler to create new message handler.
func NewHandler(d Database, prefix string) Handler {
	return &handler{
		db:     d,
		prefix: prefix,
	}
}

// Handler to get handler function.
func (h *handler) Handler() func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore all messages created by the bot itself.
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Command and prefix check.
		if h.PrefixCheck(m.Content) {
			return
		}

		// Remove prefix.
		m.Content = h.CleanPrefix(m.Content)

		// Get user data.
		user, exist, err := h.db.GetUser(m.Author.ID)
		if err != nil {
			log.Println(err)
			return
		}

		// Handle first time.
		if !exist {
			if err = h.HandleFirstTime(s, m); err != nil {
				log.Println(err)
				return
			}
		}

		// Get arguments.
		r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`)
		args := r.FindAllString(m.Content, -1)

		switch args[0] {
		case "ping":
			err = h.HandlePing(s, user.ChannelID)
		case "help":
			err = h.HandleHelp(s, user)
		case "filter":
			err = h.HandleFilter(s, user, args)
		case "category":
			err = h.HandleCategory(s, user, args)
		case "query":
			err = h.HandleQuery(s, user, args)
		case "subscribe":
			err = h.HandleSubscribe(s, user, args)
		default:
			return
		}

		if err != nil {
			log.Println(err)
		}
	}
}

// PrefixCheck to check if prefix valid.
func (h *handler) PrefixCheck(cmd string) bool {
	return len(cmd) <= len(h.prefix) || cmd[:len(h.prefix)] != h.prefix
}

// CleanPrefix to remove prefix from command.
func (h *handler) CleanPrefix(cmd string) string {
	return cmd[len(h.prefix):]
}

func (h *handler) handleInvalid(s *discordgo.Session, chID string) error {
	_, err := s.ChannelMessageSend(chID, invalidCmd)
	return err
}

func (h *handler) handleSuccess(s *discordgo.Session, user User) error {
	msg := discordgo.MessageEmbed{
		Color:       blueColor,
		Title:       "Updated",
		Description: "Your filtered Nyaa will look like [this](" + getNyaaQuery(user) + ").",
	}
	_, err := s.ChannelMessageSendEmbed(user.ChannelID, &msg)
	return err
}

// HandleFirstTime to handle user first time.
func (h *handler) HandleFirstTime(s *discordgo.Session, m *discordgo.MessageCreate) error {
	// Get DM channel.
	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		return err
	}

	// Create new user.
	if err = h.db.CreateUser(m.Author.ID, c.ID); err != nil {
		return err
	}

	msg := discordgo.MessageEmbed{
		Title:       introTitle + " " + m.Author.Username,
		Description: introContent,
		Color:       blueColor,
	}

	_, err = s.ChannelMessageSendEmbed(c.ID, &msg)
	return err
}

// HandlePing to reply "ping" message.
func (h *handler) HandlePing(s *discordgo.Session, chID string) error {
	_, err := s.ChannelMessageSend(chID, "pong!")
	return err
}

// HandleHelp to handle "help" message.
func (h *handler) HandleHelp(s *discordgo.Session, user User) error {
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
	return err
}

// HandleFilter to handle "filter" message.
func (h *handler) HandleFilter(s *discordgo.Session, user User, args []string) error {
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
		return err
	}

	// Invalid command.
	if len(args) != 3 || args[1] != "set" || getValueFromKey(filters, args[2]) == "" {
		return h.handleInvalid(s, user.ChannelID)
	}

	user.Filter = args[2]
	if err := h.db.UpdateUser(user); err != nil {
		return err
	}

	return h.handleSuccess(s, user)
}

// HandleCategory to handle "category" message.
func (h *handler) HandleCategory(s *discordgo.Session, user User, args []string) error {
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
		return err
	}

	// Invalid command.
	if len(args) != 3 || args[1] != "set" || getValueFromKey(categories, args[2]) == "" {
		return h.handleInvalid(s, user.ChannelID)
	}

	user.Category = args[2]
	if err := h.db.UpdateUser(user); err != nil {
		return err
	}

	return h.handleSuccess(s, user)
}

// HandleQuery to handle "query" message.
func (h *handler) HandleQuery(s *discordgo.Session, user User, args []string) error {
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
		return err
	}

	// Invalid command.
	if len(args) < 3 || (args[1] != "add" && args[1] != "delete") {
		return h.handleInvalid(s, user.ChannelID)
	}

	if args[1] == "add" {
		user.Queries = append(user.Queries, strings.Join(args[2:], " "))
		if err := h.db.UpdateUser(user); err != nil {
			return err
		}
	} else {
		newQuery := []string{}
		for i, q := range user.Queries {
			if !inArray(args[2:], strconv.Itoa(i)) {
				newQuery = append(newQuery, q)
			}
		}

		user.Queries = newQuery
		if err := h.db.UpdateUser(user); err != nil {
			return err
		}
	}

	return h.handleSuccess(s, user)
}

// HandleSubscribe to handle "subscribe" message.
func (h *handler) HandleSubscribe(s *discordgo.Session, user User, args []string) error {
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
		return err
	}

	// Invalid command.
	if len(args) != 2 || (args[1] != "on" && args[1] != "off") {
		return h.handleInvalid(s, user.ChannelID)
	}

	user.Subscribe = args[1] == "on"
	if err := h.db.UpdateUser(user); err != nil {
		return err
	}

	return h.handleSuccess(s, user)
}
