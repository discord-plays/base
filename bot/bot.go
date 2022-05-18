package bot

import (
	"context"
	"fmt"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/discord-plays/base/commands"
	"github.com/discord-plays/base/config"
	"github.com/discord-plays/base/iface"
	"github.com/discord-plays/base/utils"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type DiscordPlaysBot struct {
	iface.Module
	session       *session.Session
	application   *discord.Application
	commands      map[string]iface.Command
	components    map[discord.ComponentID]iface.Component
	dataStructure *config.DataStructure
	quit          chan struct{}
}

func NewDiscordPlaysBot(token string, dataStructure *config.DataStructure) *DiscordPlaysBot {
	s := session.New("Bot " + token)
	app, err := s.CurrentApplication()
	if err != nil {
		log.Fatalln("Failed to get application ID:", err)
	}
	bot := DiscordPlaysBot{
		session:       s,
		application:   app,
		commands:      make(map[string]iface.Command),
		components:    make(map[discord.ComponentID]iface.Component),
		dataStructure: dataStructure,
		quit:          make(chan struct{}),
	}
	return bot.init()
}

func (bot *DiscordPlaysBot) Session() *session.Session {
	return bot.session
}

func (bot *DiscordPlaysBot) Application() *discord.Application {
	return bot.application
}

func (bot *DiscordPlaysBot) Commands() map[string]iface.Command {
	return bot.commands
}

func (bot *DiscordPlaysBot) CreditsConfig() *config.CreditsJson {
	return bot.dataStructure.Credits
}

func (bot *DiscordPlaysBot) GameConfig() *config.GameJson {
	return bot.dataStructure.Game
}

func (bot *DiscordPlaysBot) StatusConfig() *config.StatusJson {
	return bot.dataStructure.Status
}

func (bot *DiscordPlaysBot) Run() {
	go bot.connect()
}

func (bot *DiscordPlaysBot) End() {
	close(bot.quit)
}

func (bot *DiscordPlaysBot) Hang() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	bot.End()
	fmt.Println()
	log.Println("Quitting Discord Plays bot...")
}

func (bot *DiscordPlaysBot) AddCommandCallback(c iface.Command) {
	bot.commands[c.Name()] = c
}

func (bot *DiscordPlaysBot) AddComponentCallback(c iface.Component) {
	bot.components[c.Id()] = c
}

func (bot *DiscordPlaysBot) connect() {
	if err := bot.session.Open(context.Background()); err != nil {
		log.Fatalln("failed to open:", err)
	}
	defer func(Session *session.Session) {
		err := Session.Close()
		if err != nil {
			log.Println()
		}
	}(bot.session)

	log.Println("Gateway connected!")

	// Block forever.
	<-bot.quit

	_ = bot.session.Close()
}

func (bot *DiscordPlaysBot) init() *DiscordPlaysBot {
	bot.AddCommandCallback(&commands.UpdateGuildCommandsCmd{})
	bot.AddCommandCallback(&commands.CreditsCmd{})

	bot.session.AddIntents(gateway.IntentGuilds)
	bot.session.AddIntents(gateway.IntentGuildMessages)

	bot.session.AddHandler(func(e *gateway.GuildCreateEvent) {
		if bot.MissingGuildCommands(e.Guild.ID) {
			fmt.Printf("Joining guild %s but missing guild commands...\n", e.ID)
			err := bot.UpdateGuildCommands(e.Guild.ID)
			if err != nil {
				fmt.Println("Failed to register new guild commands:", err)
			}
		}
	})
	bot.session.AddHandler(func(e *gateway.InteractionCreateEvent) {
		var resp api.InteractionResponse
		switch data := e.Data.(type) {
		case *discord.CommandInteraction:
			if commandName, ok := bot.commands[data.Name]; ok {
				resp = commandName.Execute(bot, e, data)
			} else {
				resp = utils.EphemeralErrorResponse("Unknown command: " + data.Name)
			}
		case discord.ComponentInteraction:
			if componentName, ok := bot.components[data.ID()]; ok {
				resp = componentName.Execute(bot, e, data)
			} else {
				resp = utils.EphemeralErrorResponse("Unknown component: " + string(data.ID()))
			}
		default:
			log.Printf("unknown interaction type %T", e.Data)
			return
		}
		if err := bot.session.RespondInteraction(e.ID, e.Token, resp); err != nil {
			log.Println("failed to send interaction callback:", err)
		}
	})
	return bot
}

func (bot *DiscordPlaysBot) MissingGuildCommands(guildId discord.GuildID) bool {
	a, err := bot.session.GuildCommands(bot.application.ID, guildId)
	if err != nil {
		return true
	}
	return len(a) == 0
}

func (bot *DiscordPlaysBot) UpdateGuildCommands(guildId discord.GuildID) error {
	a := make([]api.CreateCommandData, len(bot.commands))
	i := 0
	for _, v := range bot.commands {
		a[i] = api.CreateCommandData{
			Name:        v.Name(),
			Description: v.Description(),
			Options:     v.Options(),
			Type:        v.CommandType(),
		}
		i++
	}
	_, err := bot.session.BulkOverwriteGuildCommands(bot.application.ID, guildId, a)
	return err
}
