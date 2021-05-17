package httpmodels

import (
	instancepoolsModels "github.com/polar-rams/databricks-sdk-golang/azure/instance_pools/models"
)

//ListResp is to compose the response to list an array of instance pools' stats
type ListResp struct {
	InstancePools []instancepoolsModels.InstancePoolStats `json:"instance_pools,omitempty" url:"instance_pools,omitempty"`
}
