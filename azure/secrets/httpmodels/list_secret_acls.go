package httpmodels

import "github.com/polar-rams/databricks-sdk-golang/aws/models"

type ListSecretAclsReq struct {
	Scope string `json:"scope,omitempty" url:"scope,omitempty"`
}

type ListSecretAclsResp struct {
	Items *models.AclItem `json:"items,omitempty" url:"items,omitempty"`
}
