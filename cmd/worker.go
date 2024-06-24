package main

import (
	"github.com/benetis/weather-whisperer/internal/meteo"
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
	w.RegisterWorkflow(meteo.FetchForecastsWorkflow)
	w.RegisterActivity(meteo.FetchForecasts)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
