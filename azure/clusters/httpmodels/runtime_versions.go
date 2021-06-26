package httpmodels

import "github.com/gpompe/databricks-sdk-golang/azure/clusters/models"

type RuntimeVersionsResp struct {
	Versions []models.SparkVersion `json:"versions,omitempty" url:"versions,omitempty"`
}
