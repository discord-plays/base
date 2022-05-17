package iface

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
)

type Command interface {
	Name() string
	Description() string
	Options() discord.CommandOptions
	CommandType() discord.CommandType
	Execute(bot Module, e *gateway.InteractionCreateEvent, data *discord.CommandInteraction) api.InteractionResponse
}
