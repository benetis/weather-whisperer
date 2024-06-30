package command

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
)

func MenuCommand(chatID int64, bot *tgbotapi.BotAPI) bool {
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Kada lis?", "/rain"),
			tgbotapi.NewInlineKeyboardButtonData("Atnaujink duomenis", "/refresh"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Pasirink:")
	msg.ReplyMarkup = buttons
	send, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
		return true
	}

	log.Printf("Sent message with ID %d", send.MessageID)
	return false
}
