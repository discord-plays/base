package iface

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/session"
)

type Module interface {
	Session() *session.Session
	Application() *discord.Application
	Commands() map[string]Command
	AddCommand(command Command)
	UpdateGuildCommands(guildId discord.GuildID) error
}
