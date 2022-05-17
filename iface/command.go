package iface

import (
	"github.com/diamondburned/arikawa/v3/discord"
)

type Command interface {
	Name() string
	Description() string
	Options() discord.CommandOptions
	CommandType() discord.CommandType
	Execute()
}
