package base

import (
	"os"
	"testing"
	"time"
)

func TestBaseBot(t *testing.T) {
	token := os.Getenv("TOKEN")
	bot := NewDiscordPlaysBot(token)
	bot.Run()

	time.Sleep(30 * time.Second)
	bot.End()
}
