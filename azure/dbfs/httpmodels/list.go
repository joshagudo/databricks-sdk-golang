package httpmodels

import (
	"github.com/xinsnake/databricks-sdk-golang/azure/dbfs/models"
)

type ListReq struct {
	Path string `json:"path,omitempty" url:"path,omitempty"`
}

type ListResp struct {
	Files []models.FileInfo `json:"files,omitempty" url:"files,omitempty"`
}