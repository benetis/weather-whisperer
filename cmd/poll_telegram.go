package main

import (
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
	w.RegisterWorkflow(workflows.TelegramCommand)

	err = w.Start()
	if err != nil {
		log.Fatalf("Unable to start Worker: %v", err)
	}

	log.Println("Worker started successfully")

	telegram.Poll()
}
