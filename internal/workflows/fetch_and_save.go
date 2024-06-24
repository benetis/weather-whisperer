package workflows

import (
	"github.com/benetis/weather-whisperer/internal/meteo"
	"github.com/benetis/weather-whisperer/internal/storage"
	"go.temporal.io/sdk/workflow"
	"time"
)

func FetchAndSaveForecastsWorkflow(ctx workflow.Context, city string) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Fetch forecasts workflow started")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result meteo.ForecastsResponse
	err := workflow.ExecuteActivity(ctx, meteo.FetchForecasts, city).Get(ctx, &result)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, storage.SaveForecasts, result).Get(ctx, nil)

	return err
}
