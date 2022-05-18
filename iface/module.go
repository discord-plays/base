package iface

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/session"
	config2 "github.com/discord-plays/base/config"
)

type Module interface {
	Session() *session.Session
	Application() *discord.Application
	Commands() map[string]Command
	AddCommandCallback(command Command)
	UpdateGuildCommands(guildId discord.GuildID) error
	CreditsConfig() *config2.CreditsJson
	GameConfig() *config2.GameJson
	StatusConfig() *config2.StatusJson
}
