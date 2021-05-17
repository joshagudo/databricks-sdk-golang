package httpmodels

import "github.com/polar-rams/databricks-sdk-golang/aws/models"

type ListSecretACLsReq struct {
	Scope string `json:"scope,omitempty" url:"scope,omitempty"`
}

type ListSecretACLsResp struct {
	Items *models.ACLItem `json:"items,omitempty" url:"items,omitempty"`
}
