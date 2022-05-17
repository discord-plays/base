package base

import (
	"context"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/discord-plays/base/iface"
	"log"
)

type DiscordPlaysBot struct {
	Session     *session.Session
	Application *discord.Application
	Commands    map[string]iface.Command
	quit        chan struct{}
}

func NewDiscordPlaysBot(token string) *DiscordPlaysBot {
	s := session.New(token)
	s.AddHandler(func(e *gateway.InteractionCreateEvent) {
		data := api.InteractionResponse{
			Type: api.MessageInteractionWithSource,
			Data: &api.InteractionResponseData{
				Content: option.NewNullableString("Pong!"),
			},
		}
		if err := s.RespondInteraction(e.ID, e.Token, data); err != nil {
			log.Println("failed to send interaction callback:", err)
		}
	})

	app, err := s.CurrentApplication()
	if err != nil {
		log.Fatalln("Failed to get application ID:", err)
	}
	return &DiscordPlaysBot{
		Session:     s,
		Application: app,
		Commands:    make(map[string]iface.Command),
		quit:        make(chan struct{}),
	}
}

func (bot *DiscordPlaysBot) Run() {
	go bot.connect()
}

func (bot *DiscordPlaysBot) End() {
	close(bot.quit)
}

func (bot *DiscordPlaysBot) connect() {
	if err := bot.Session.Open(context.Background()); err != nil {
		log.Fatalln("failed to open:", err)
	}
	defer func(Session *session.Session) {
		err := Session.Close()
		if err != nil {
			log.Println()
		}
	}(bot.Session)

	log.Println("Gateway connected. Getting all guild commands.")

	// Block forever.
	<-bot.quit
}

func (bot *DiscordPlaysBot) UpdateCommands() {
	newCommands := []api.CreateCommandData{
		{
			Name:        "ping",
			Description: "Basic ping command.",
		},
	}

	if _, err := bot.Session.BulkOverwriteGuildCommands(bot.Session.ID, guildID, newCommands); err != nil {
		log.Fatalln("failed to create guild command:", err)
	}
}
