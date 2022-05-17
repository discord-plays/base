package commands

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/discord-plays/base"
	"log"
)

type UpdateCommands struct {
}

func (u *UpdateCommands) GetName() string {
	return "update-commands"
}

func (u *UpdateCommands) Execute(bot *base.DiscordPlaysBot) {
	a := make([]api.CreateCommandData, len(bot.Commands))
	i := 0
	for _, v := range bot.Commands {
		a[i] = api.CreateCommandData{
			Name:        v.Name(),
			Description: v.Description(),
			Options:     v.Options(),
			Type:        v.CommandType(),
		}
		i++
	}

	_, err := bot.Session.BulkOverwriteGuildCommands(bot.Application.ID, bot.Application.GuildID, a)
	if err != nil {
		log.Println("Failed to update guild commands:", err)
		return
	}
}
