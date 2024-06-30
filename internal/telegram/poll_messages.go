package telegram

import (
	"context"
	"go.temporal.io/sdk/client"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
)

func Poll() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	var err error
	bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalf("Unable to create Temporal client: %v", err)
	}
	defer c.Close()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Error getting updates channel: %v", err)
	}

	for update := range updates {
		switch {
		case update.Message != nil:
			handleMessage(update, c)
		case update.CallbackQuery != nil:
			handleCallback(update, c)
		}
	}
}

func handleMessage(update tgbotapi.Update, c client.Client) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	workflowOptions := client.StartWorkflowOptions{
		ID:        "telegram-command",
		TaskQueue: "telegram-task-queue",
	}

	_, err := c.ExecuteWorkflow(context.Background(),
		workflowOptions, "TelegramCommand",
		update.Message.Text,
		update.Message.Chat.ID,
	)

	if err != nil {
		log.Fatalf("Unable to execute workflow: %v", err)
	}
}

func handleCallback(update tgbotapi.Update, c client.Client) {
	log.Printf("[%s] %s", update.CallbackQuery.From.UserName, update.CallbackQuery.Data)

	workflowOptions := client.StartWorkflowOptions{
		ID:        "telegram-callback",
		TaskQueue: "telegram-task-queue",
	}

	_, err := c.ExecuteWorkflow(context.Background(),
		workflowOptions, "TelegramCommand",
		update.CallbackQuery.Data,
		update.CallbackQuery.Message.Chat.ID,
	)

	if err != nil {
		log.Fatalf("Unable to execute workflow: %v", err)
	}

	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
	if _, err := bot.AnswerCallbackQuery(callback); err != nil {
		log.Fatalf("Unable to acknowledge callback: %v", err)
	}
}
