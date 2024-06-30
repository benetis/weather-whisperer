package command

import (
	"context"
	"github.com/benetis/weather-whisperer/internal/workflows"
	"go.temporal.io/sdk/client"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
)

func RefreshDataCommand(chatID int64, bot *tgbotapi.BotAPI, temporal client.Client) {
	msg := tgbotapi.NewMessage(chatID, "Atnaujinama...")
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error sending refresh message: %v", err)
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:        "refresh-data",
		TaskQueue: "telegram-task-queue",
	}

	we, err := temporal.ExecuteWorkflow(
		context.Background(),
		workflowOptions,
		workflows.DownloadForecastsWorkflow,
		"kaunas",
	)

	if err != nil {
		log.Fatalf("Unable to execute workflow: %v", err)
	}

	log.Printf("Workflow started with ID %s", we.GetID())
}
