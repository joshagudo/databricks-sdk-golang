package httpmodels

import "github.com/polar-rams/databricks-sdk-golang/aws/models"

type GetSecretAclReq struct {
	Scope     string `json:"scope,omitempty" url:"scope,omitempty"`
	Principal string `json:"principal,omitempty" url:"principal,omitempty"`
}

type GetSecretAclResp struct {
	Principal  string                `json:"principal,omitempty" url:"principal,omitempty"`
	Permission *models.AclPermission `json:"permission,omitempty" url:"permission,omitempty"`
}
