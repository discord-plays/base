package utils

import (
	"errors"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

var (
	ErrFetchGuild   = errors.New("failed to get current guild")
	ErrFetchChannel = errors.New("failed to get current channel")
)

func InteractionUserPermissions(session *session.Session, e *gateway.InteractionCreateEvent) (discord.Permissions, error) {
	guild, err := session.Guild(e.GuildID)
	if err != nil {
		return 0, ErrFetchGuild
	}
	channel, err := session.Channel(e.ChannelID)
	if err != nil {
		return 0, ErrFetchChannel
	}
	return discord.CalcOverwrites(*guild, *channel, *e.Member), nil
}
