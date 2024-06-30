package main

import (
	"context"
	"github.com/benetis/weather-whisperer/internal/meteo"
	"github.com/benetis/weather-whisperer/internal/workflows"
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
		workflows.DownloadForecastsWorkflow,
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

	log.Println("Workflow completed with result", result)
}
