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
		var resp api.InteractionResponse
		switch data := e.Data.(type) {
		case *discord.CommandInteraction:
			if data.Name != "buttons" {
				resp = api.InteractionResponse{
					Type: api.MessageInteractionWithSource,
					Data: &api.InteractionResponseData{
						Content: option.NewNullableString("Unknown command: " + data.Name),
					},
				}
				break
			}
			// Send a message with a button back on slash commands.
			resp = api.InteractionResponse{
				Type: api.MessageInteractionWithSource,
				Data: &api.InteractionResponseData{
					Content: option.NewNullableString("This is a message with a button!"),
					Components: discord.ComponentsPtr(
						&discord.ActionRowComponent{
							&discord.ButtonComponent{
								Label:    "Hello World!",
								CustomID: "first_button",
								Emoji:    &discord.ComponentEmoji{Name: "👋"},
								Style:    discord.PrimaryButtonStyle(),
							},
							&discord.ButtonComponent{
								Label:    "Secondary",
								CustomID: "second_button",
								Style:    discord.SecondaryButtonStyle(),
							},
							&discord.ButtonComponent{
								Label:    "Success",
								CustomID: "success_button",
								Style:    discord.SuccessButtonStyle(),
							},
							&discord.ButtonComponent{
								Label:    "Danger",
								CustomID: "danger_button",
								Style:    discord.DangerButtonStyle(),
							},
						},
						// This is automatically put into its own row.
						&discord.ButtonComponent{
							Label: "Link",
							Style: discord.LinkButtonStyle("https://google.com"),
						},
					),
				},
			}
		case discord.ComponentInteraction:
			resp = api.InteractionResponse{
				Type: api.UpdateMessage,
				Data: &api.InteractionResponseData{
					Content: option.NewNullableString("Custom ID: " + string(data.ID())),
				},
			}
		default:
			log.Printf("unknown interaction type %T", e.Data)
			return
		}
		if err := s.RespondInteraction(e.ID, e.Token, resp); err != nil {
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
