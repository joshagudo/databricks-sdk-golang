package httpmodels

import "github.com/gpompe/databricks-sdk-golang/azure/secrets/models"

type ListSecretScopesResp struct {
	Scopes *[]models.SecretScope `json:"scopes,omitempty" url:"scopes,omitempty"`
}
