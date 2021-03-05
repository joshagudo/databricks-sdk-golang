package httpmodels

import (
	"github.com/xinsnake/databricks-sdk-golang/azure/jobs/models"
)

type GenericRunReq struct {
	RunID int64 `json:"run_id,omitempty" url:"run_id,omitempty"`
}

type GenericJobReq struct {
	JobID int64 `json:"job_id,omitempty" url:"job_id,omitempty"`
}

type GenericJobRunsUpdateReq struct {
	JobID          int64              `json:"job_id,omitempty" url:"job_id,omitempty"`
	NewSettings    models.JobSettings `json:"new_settings,omitempty" url:"new_settings,omitempty"`
	FieldsToRemove []string           `json:"fields_to_remove,omitempty" url:"fields_to_remove,omitempty"`
}
