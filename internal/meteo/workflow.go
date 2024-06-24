package meteo

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

func FetchForecastsWorkflow(ctx workflow.Context, city string) (ForecastsResponse, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Fetch forecasts workflow started")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result ForecastsResponse
	err := workflow.ExecuteActivity(ctx, FetchForecasts, city).Get(ctx, &result)
	if err != nil {
		return ForecastsResponse{}, err
	}

	return result, nil
}
