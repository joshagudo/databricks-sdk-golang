package httpmodels

import (
	"github.com/xinsnake/databricks-sdk-golang/azure/dbfs/models"
)

type CreateResp struct {
	Handle int64 `json:"handle,omitempty" url:"handle,omitempty"`
}
