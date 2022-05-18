package utils

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

func EphemeralSuccessResponse(message string) api.InteractionResponse {
	return EphemeralQuickResponse("Success:", message, 0x55ff55)
}

func EphemeralErrorResponse(message string) api.InteractionResponse {
	return EphemeralQuickResponse("Error:", message, 0xff5555)
}

func EphemeralQuickResponse(title string, message string, color discord.Color) api.InteractionResponse {
	return api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Flags: api.EphemeralResponse,
			Embeds: &[]discord.Embed{
				{
					Title:       title,
					Description: message,
					Color:       color,
				},
			},
		},
	}
}
