package httpmodels

import "github.com/gpompe/databricks-sdk-golang/azure/clusters/models"

type ListResp struct {
	Clusters []models.ClusterInfo
}
