package httpmodels

import (
	"github.com/xinsnake/databricks-sdk-golang/azure/dbfs/models"
)

type ListResp struct {
	Files []models.FileInfo `json:"files,omitempty" url:"files,omitempty"`
}