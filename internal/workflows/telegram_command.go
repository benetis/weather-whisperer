package workflows

import (
	"github.com/benetis/weather-whisperer/internal/telegram"
	"go.temporal.io/sdk/workflow"
)

func TelegramCommand(ctx workflow.Context, data string, chatID int64) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("TelegramCommand started", "data", data, "chatID", chatID)

	telegram.HandleCommand(data, chatID)

	return nil
}
