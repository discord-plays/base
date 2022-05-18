package commands

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/discord-plays/base/iface"
	"strings"
)

type CreditsCmd struct {
}

func (u *CreditsCmd) Name() string {
	return "credits"
}

func (u *CreditsCmd) Description() string {
	return "Credits for bot"
}

func (u *CreditsCmd) Options() discord.CommandOptions {
	return discord.CommandOptions{}
}

func (u *CreditsCmd) CommandType() discord.CommandType {
	return discord.ChatInputCommand
}

func (u *CreditsCmd) Execute(bot iface.Module, e *gateway.InteractionCreateEvent, _ *discord.CommandInteraction) api.InteractionResponse {
	gameConfig := bot.GameConfig()
	creditsConfig := bot.CreditsConfig()
	return api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Flags: api.EphemeralResponse,
			Embeds: &[]discord.Embed{
				{
					Title:       "Credits",
					Description: strings.Join(creditsConfig.Description, "\n"),
					Color:       0x15d0ed,
					Author: &discord.EmbedAuthor{
						Name: gameConfig.ProjectName,
						Icon: gameConfig.LogoAddress,
					},
					Fields: []discord.EmbedField{
						{
							Name:   "Idea By:",
							Value:  strings.Join(creditsConfig.IdeaBy, "\n"),
							Inline: true,
						},
						{
							Name:   "Developers:",
							Value:  strings.Join(creditsConfig.Developers, "\n"),
							Inline: true,
						},
						{
							Name:   "Artists:",
							Value:  strings.Join(creditsConfig.Artists, "\n"),
							Inline: true,
						},
						{
							Name:  "Thanks to:",
							Value: strings.Join(creditsConfig.ThanksTo, "\n"),
						},
						{
							Name:  "Github:",
							Value: strings.Join(creditsConfig.Github, "\n"),
						},
					},
				},
			},
		},
	}
}
