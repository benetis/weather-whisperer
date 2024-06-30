package main

import (
	"context"
	"github.com/benetis/weather-whisperer/internal/meteo"
	"github.com/benetis/weather-whisperer/internal/storage"
	"github.com/benetis/weather-whisperer/internal/telegram"
	"github.com/benetis/weather-whisperer/internal/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalf("Unable to create Temporal client: %v", err)
	}
	defer c.Close()

	w := worker.New(c, "telegram-task-queue", worker.Options{})
	w.RegisterWorkflow(workflows.DownloadForecastsWorkflow)
	w.RegisterActivity(meteo.FetchForecasts)
	w.RegisterActivity(storage.SaveForecasts)

	err = w.Start()
	if err != nil {
		log.Fatalf("Unable to start Worker: %v", err)
	}

	log.Println("Worker started successfully")

	startCronWorkflow(c)

	telegram.Poll(c)
}

func startCronWorkflow(c client.Client) {
	workflowOptions := client.StartWorkflowOptions{
		ID:           "refresh-data-scheduled",
		TaskQueue:    "telegram-task-queue",
		CronSchedule: "@every 30m",
	}

	we, err := c.ExecuteWorkflow(
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
