package command

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
)

func RefreshDataCommand(chatID int64, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(chatID, "Refreshing data...")
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error sending refresh message: %v", err)
	}

}
