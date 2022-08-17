package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) readyHandler() func(*discordgo.Session, *discordgo.Ready) {
	return func(ss *discordgo.Session, _ *discordgo.Ready) {
		ss.UpdateListeningStatus(fmt.Sprintf("%shelp for command list", b.prefix))
	}
}
