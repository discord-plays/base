package iface

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
)

type Component interface {
	Id() discord.ComponentID
	Description() string
	Execute(bot Module, e *gateway.InteractionCreateEvent, data discord.ComponentInteraction) api.InteractionResponse
}
