package httpmodels

import (
	"github.com/xinsnake/databricks-sdk-golang/azure/jobs/models"
)

type JobsRunNowReq struct {
	JobID         int64 `json:"job_id,omitempty" url:"job_id,omitempty"`
	RunParameters models.RunParameters
}
