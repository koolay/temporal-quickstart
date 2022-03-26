package workflows

import (
	"go.temporal.io/sdk/workflow"
)

const (
	CronWorkflowID = "CronWorkFlow"
	CronQueueName  = "cron"
)

func CronWorkflow(ctx workflow.Context) error {
	workflow.GetLogger(ctx).Info("Cron workflow started.", "StartTime", workflow.Now(ctx))
	return nil
}
