package httpmodels

import (
	"github.com/gpompe/databricks-sdk-golang/azure/workspace/models"
)

type ListReq struct {
	Path string `json:"path,omitempty" url:"path,omitempty"`
}

type ListResp struct {
	Objects *[]models.ObjectInfo `json:"objects,omitempty" url:"objects,omitempty"`
}
