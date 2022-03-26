package main

import (
	"context"
	"log"

	"github.com/koolay/temporal-quickstart/pkg/workflows"
	"go.temporal.io/sdk/client"
	clientSdk "go.temporal.io/sdk/client"
)

func main() {
	client, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer client.Close()

	options := clientSdk.StartWorkflowOptions{
		ID:           workflows.CronWorkflowID,
		TaskQueue:    workflows.CronQueueName,
		CronSchedule: "@every 10s",
	}

	running, err := client.ExecuteWorkflow(context.Background(), options, workflows.CronWorkflow)
	if err != nil {
		log.Fatalln("unable to start cron workflow", err)
	}

	log.Println("execute workflow:", running.GetID(), running.GetRunID())
}
