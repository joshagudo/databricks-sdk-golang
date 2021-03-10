package httpmodels

import (
	"github.com/xinsnake/databricks-sdk-golang/azure/jobs/models"
)

type JobsRunsGetOutputResp struct {
	NotebookOutput models.NotebookOutput `json:"notebook_output,omitempty" url:"notebook_output,omitempty"`
	Error          string                `json:"error,omitempty" url:"error,omitempty"`
	Metadata       models.Run            `json:"metadata,omitempty" url:"metadata,omitempty"`
}
