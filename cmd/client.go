package main

import (
	"context"
	"fmt"
	"github.com/benetis/weather-whisperer/internal/meteo"
	"go.temporal.io/sdk/client"
	"log"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "weather_workflow",
		TaskQueue: "weather-task-queue",
	}

	we, err := c.ExecuteWorkflow(
		context.Background(),
		workflowOptions,
		meteo.FetchForecastsWorkflow,
		"kaunas",
	)

	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	var result meteo.ForecastsResponse
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

	fmt.Printf("Weather in %s: %.2f°C\n", result.Forecasts[0].ForecastTimeUtc, result.Forecasts[0].AirTemperature)
}
