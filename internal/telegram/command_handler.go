package telegram

import (
	"github.com/benetis/weather-whisperer/internal/telegram/command"
	"go.temporal.io/sdk/client"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var bot *tgbotapi.BotAPI

func HandleCommand(commandStr string, chatID int64, temporal client.Client) {
	switch commandStr {
	case ".":
		command.MenuCommand(chatID, bot)
	case "/refresh":
		command.RefreshDataCommand(chatID, bot, temporal)
	}
}
