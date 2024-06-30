package telegram

import (
	"go.temporal.io/sdk/client"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
)

func Poll(temporal client.Client) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	var err error
	bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Error getting updates channel: %v", err)
	}

	for update := range updates {
		switch {
		case update.Message != nil:
			handleMessage(update, temporal)
		case update.CallbackQuery != nil:
			handleCallback(update, temporal)
		}
	}
}

func handleMessage(update tgbotapi.Update, temporal client.Client) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	HandleCommand(update.Message.Text, update.Message.Chat.ID, temporal)
}

func handleCallback(update tgbotapi.Update, temporal client.Client) {
	log.Printf("[%s] %s", update.CallbackQuery.From.UserName, update.CallbackQuery.Data)

	HandleCommand(update.CallbackQuery.Data, update.CallbackQuery.Message.Chat.ID, temporal)

	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
	if _, err := bot.AnswerCallbackQuery(callback); err != nil {
		log.Fatalf("Unable to acknowledge callback: %v", err)
	}
}
