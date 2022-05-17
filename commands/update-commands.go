package commands

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/discord-plays/base/iface"
	"github.com/discord-plays/base/utils"
	"log"
)

type UpdateCommands struct {
}

func (u *UpdateCommands) Name() string {
	return "update-commands"
}

func (u *UpdateCommands) Description() string {
	return "Update guild commands for this bot"
}

func (u *UpdateCommands) Options() discord.CommandOptions {
	return discord.CommandOptions{}
}

func (u *UpdateCommands) CommandType() discord.CommandType {
	return discord.ChatInputCommand
}

func (u *UpdateCommands) Execute(bot iface.Module, e *gateway.InteractionCreateEvent, _ *discord.CommandInteraction) api.InteractionResponse {
	guild, err := bot.Session().Guild(e.GuildID)
	if err != nil {
		return utils.EphemeralQuickResponse("Failed to get current guild")
	}
	member, err := bot.Session().Member(e.GuildID, e.Member.User.ID)
	if err != nil {
		return utils.EphemeralQuickResponse("Failed to get user information")
	}
	member
	err = bot.UpdateGuildCommands(e.GuildID)
	if err != nil {
		log.Println("Failed to update guild commands:", err)
		return utils.EphemeralQuickResponse("Failed to update guild commands")
	}
	return utils.EphemeralQuickResponse("Updated guild commands")
}
