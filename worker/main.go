package main

import (
	"log"

	"github.com/koolay/temporal-quickstart/pkg/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

var serverHost = "localhost:7233"

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	client, err := client.NewClient(client.Options{
		HostPort:  serverHost,
		Namespace: "default",
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	defer client.Close()

	w := worker.New(client, "cron", worker.Options{})

	w.RegisterWorkflow(workflows.CronWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
