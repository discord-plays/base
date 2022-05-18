package commands

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/discord-plays/base/iface"
	"github.com/discord-plays/base/utils"
	"log"
)

type UpdateGuildCommandsCmd struct {
}

func (u *UpdateGuildCommandsCmd) Name() string {
	return "update-guild-commands"
}

func (u *UpdateGuildCommandsCmd) Description() string {
	return "Update guild commands for this bot"
}

func (u *UpdateGuildCommandsCmd) Options() discord.CommandOptions {
	return discord.CommandOptions{}
}

func (u *UpdateGuildCommandsCmd) CommandType() discord.CommandType {
	return discord.ChatInputCommand
}

func (u *UpdateGuildCommandsCmd) Execute(bot iface.Module, e *gateway.InteractionCreateEvent, _ *discord.CommandInteraction) api.InteractionResponse {
	perms, err := utils.InteractionUserPermissions(bot.Session(), e)
	if err != nil {
		return utils.EphemeralErrorResponse(err.Error())
	}
	if !perms.Has(discord.PermissionManageGuild) {
		return utils.EphemeralErrorResponse("You need the manage guild permission to use this command")
	}
	err = bot.UpdateGuildCommands(e.GuildID)
	if err != nil {
		log.Println("Failed to update guild commands:", err)
		return utils.EphemeralErrorResponse("Failed to update guild commands")
	}
	return utils.EphemeralSuccessResponse("Updated guild commands")
}
