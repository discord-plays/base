package utils

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

func EphemeralQuickResponse(message string) api.InteractionResponse {
	return api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Flags:   api.EphemeralResponse,
			Content: option.NewNullableString(message),
		},
	}
}
