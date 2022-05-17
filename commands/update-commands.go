package commands

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/discord-plays/base"
)

type UpdateCommands struct {
}

func (u *UpdateCommands) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (u *UpdateCommands) Execute(bot *base.DiscordPlaysBot) {
	a := make([]api.CreateCommandData)
	bot.Commands

	bot.Session.BulkOverwriteGuildCommands()
	bot.Application
	//TODO implement me
	panic("implement me")
}
