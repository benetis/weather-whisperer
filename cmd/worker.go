package main

import (
	"github.com/benetis/weather-whisperer/internal/meteo"
	"github.com/benetis/weather-whisperer/internal/storage"
	"github.com/benetis/weather-whisperer/internal/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "weather-task-queue", worker.Options{})
	w.RegisterWorkflow(workflows.DownloadForecastsWorkflow)
	w.RegisterActivity(meteo.FetchForecasts)
	w.RegisterActivity(storage.SaveForecasts)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
