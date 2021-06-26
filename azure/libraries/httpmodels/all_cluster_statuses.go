package httpmodels

import "github.com/gpompe/databricks-sdk-golang/azure/libraries/models"

type AllClusterStatusesResp struct {
	Statuses *[]models.ClusterLibraryStatuses `json:"statuses,omitempty" url:"statuses,omitempty"`
}
