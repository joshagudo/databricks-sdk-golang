package httpmodels

import (
	"github.com/xinsnake/databricks-sdk-golang/azure/jobs/models"
)

type JobsRunsSubmitReq struct {
	RunName          string `json:"run_name,omitempty" url:"run_name,omitempty"`
	ClusterSpec      models.ClusterSpec
	JobTask          models.JobTask
	TimeoutSeconds   int32  `json:"timeout_seconds,omitempty" url:"timeout_seconds,omitempty"`
	IdempotencyToken string `json:"idempotency_token,omitempty" url:"idempotency_token,omitempty"`
}
