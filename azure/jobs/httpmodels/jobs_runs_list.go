package httpmodels

import (
	"github.com/polar-rams/databricks-sdk-golang/azure/jobs/models"
)

type JobsRunsListReq struct {
	ActiveOnly    bool   `json:"active_only,omitempty" url:"active_only,omitempty"`
	CompletedOnly bool   `json:"completed_only,omitempty" url:"completed_only,omitempty"`
	JobID         int64  `json:"job_id,omitempty" url:"job_id,omitempty"`
	Offset        int32  `json:"offset,omitempty" url:"offset,omitempty"`
	Limit         int32  `json:"limit,omitempty" url:"limit,omitempty"`
	RunType       string `json:"run_type,omitempty" url:"run_type,omitempty"`
}

type JobsRunsListResp struct {
	Runs    []models.Run `json:"runs,omitempty" url:"runs,omitempty"`
	HasMore bool         `json:"has_more,omitempty" url:"has_more,omitempty"`
}
