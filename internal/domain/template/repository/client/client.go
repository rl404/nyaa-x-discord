package client

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	nyaaEntity "github.com/rl404/nyaa-x-discord/internal/domain/nyaa/entity"
	"github.com/rl404/nyaa-x-discord/internal/domain/template/entity"
	"github.com/rl404/nyaa-x-discord/internal/utils"
)

// Client is template client.
type Client struct {
	prefix string
}

// New to create new template client.
func New(prefix string) *Client {
	return &Client{
		prefix: prefix,
	}
}

func (c *Client) clean(str string) string {
	return strings.ReplaceAll(str, "{{prefix}}", c.prefix)
}

// GetFirstTime to get first time template.
func (c *Client) GetFirstTime(username string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       entity.IntroTitle + " " + username,
		Description: c.clean(fmt.Sprintf(entity.IntroContent, nyaaEntity.NyaaURL)),
		Color:       entity.BlueColor,
	}
}

// GetHelp to get help template.
func (c *Client) GetHelp() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: entity.BlueColor,
		Description: c.clean(fmt.Sprintf("**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s\n\n**%s**\n%s",
			entity.HelpCmd, entity.HelpContent,
			entity.FilterCmd, entity.FilterContent,
			entity.FilterSetCmd, entity.FilterSetContent,
			entity.CategCmd, entity.CategContent,
			entity.CategSetCmd, entity.CategSetContent,
			entity.QueryCmd, entity.QueryContent,
			entity.QueryAddCmd, entity.QueryAddContent,
			entity.QueryDeleteCmd, entity.QueryDeleteContent,
			entity.SubsCmd, entity.SubsContent,
			entity.SubsSetCmd, entity.SubsSetContent)),
	}
}

// GetFilters to get filters template.
func (c *Client) GetFilters(user entity.User) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: entity.BlueColor,
		Description: fmt.Sprintf("**Options**```%s```\nYour current filter: **%s - %s**\nYour filtered Nyaa will look like [this](%s).",
			c.keyValueToMessage(nyaaEntity.Filters),
			user.Filter,
			nyaaEntity.Filters.GetValueByKey(user.Filter),
			nyaaEntity.GenerateURL(user.Filter, user.Category, user.Queries),
		),
	}
}

func (c *Client) keyValueToMessage(keyValues nyaaEntity.KeyValues) (msg string) {
	for _, kv := range keyValues {
		msg += kv.Key + " : " + kv.Value + "\n"
	}
	return msg
}

// GetInvalid to get invalid template.
func (c *Client) GetInvalid() string {
	return c.clean(entity.InvalidCmd)
}

// GetUpdated to get updated template.
func (c *Client) GetUpdated(user entity.User) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:       entity.BlueColor,
		Title:       "Updated",
		Description: fmt.Sprintf("Your filtered Nyaa will look like [this](%s).", nyaaEntity.GenerateURL(user.Filter, user.Category, user.Queries)),
	}
}

// GetCategories to get categories template.
func (c *Client) GetCategories(user entity.User) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: entity.BlueColor,
		Description: fmt.Sprintf("**Options**```%s```\nYour current category: **%s - %s**\nYour filtered Nyaa will look like [this](%s).",
			c.keyValueToMessage(nyaaEntity.Categories),
			user.Category,
			nyaaEntity.Categories.GetValueByKey(user.Category),
			nyaaEntity.GenerateURL(user.Filter, user.Category, user.Queries),
		),
	}
}

// GetQueries to get queries template.
func (c *Client) GetQueries(user entity.User) *discordgo.MessageEmbed {
	if len(user.Queries) == 0 {
		return &discordgo.MessageEmbed{
			Color:       entity.BlueColor,
			Description: "**Queries**```Empty query. Go create a new one.```",
		}
	}

	queries := make([]string, len(user.Queries))
	for i, q := range user.Queries {
		queries[i] = fmt.Sprintf("%v : %s", i, q)
	}

	return &discordgo.MessageEmbed{
		Color: entity.BlueColor,
		Description: fmt.Sprintf("**Queries**```%s```\nYour filtered Nyaa will look like [this](%s).",
			strings.Join(queries, "\n"),
			nyaaEntity.GenerateURL(user.Filter, user.Category, user.Queries),
		),
	}
}

// GetSubscribe to get subscribe template.
func (c *Client) GetSubscribe(user entity.User) *discordgo.MessageEmbed {
	status := "off"
	if user.Subscribe {
		status = "on"
	}

	return &discordgo.MessageEmbed{
		Color: entity.BlueColor,
		Description: fmt.Sprintf("**Status**\nYour current subscription: **%s**\nYour filtered Nyaa will look like [this](%s).",
			status,
			nyaaEntity.GenerateURL(user.Filter, user.Category, user.Queries),
		),
	}
}

// GetNyaaUpdate to get nyaa update template.
func (c *Client) GetNyaaUpdate(feeds []entity.Feed) *discordgo.MessageEmbed {
	msg := &discordgo.MessageEmbed{
		Color: entity.BlueColor,
	}

	for _, f := range feeds {
		msg.Fields = append(msg.Fields, &discordgo.MessageEmbedField{
			Name:  utils.Ellipsis(f.Title, 100),
			Value: fmt.Sprintf("[link](%s)  •  `%s`", f.Link, f.Size),
		})
	}

	return msg
}
