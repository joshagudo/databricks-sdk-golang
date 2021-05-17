package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/secrets/httpmodels"
)

// SecretsAPI exposes the Secrets API
type SecretsAPI struct {
	Client DBClient
}

func (a SecretsAPI) init(client DBClient) SecretsAPI {
	a.Client = client
	return a
}

// CreateSecretScope create an Azure Key Vault-backed or Databricks-backed scope
func (a SecretsAPI) CreateSecretScope(createSecretScope httpmodels.CreateSecretScopeReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/scopes/create", createSecretScope, nil)
	return err
}

// DeleteSecretScope deletes a secret scope
func (a SecretsAPI) DeleteSecretScope(deleteSecretScopeReq httpmodels.DeleteSecretScopeReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/scopes/delete", deleteSecretScopeReq, nil)
	return err
}

// ListSecretScopes lists all secret scopes available in the workspace
func (a SecretsAPI) ListSecretScopes() (httpmodels.ListSecretScopesResp, error) {
	var listSecretScopesResp httpmodels.ListSecretScopesResp

	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/scopes/list", nil, nil)
	if err != nil {
		return listSecretScopesResp, err
	}

	err = json.Unmarshal(resp, &listSecretScopesResp)
	return listSecretScopesResp, err
}

// PutSecret creates or modifies a bytes secret depends on the type of scope backend with
func (a SecretsAPI) PutSecret(putSecretReq httpmodels.PutSecretReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/put", putSecretReq, nil)
	return err
}

// DeleteSecret deletes a secret depends on the type of scope backend
func (a SecretsAPI) DeleteSecret(deleteSecretReq httpmodels.DeleteSecretReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/delete", deleteSecretReq, nil)
	return err
}

// ListSecrets lists the secret keys that are stored at this scope
func (a SecretsAPI) ListSecrets(listSecretsReq httpmodels.ListSecretsReq) (httpmodels.ListSecretsResp, error) {
	var listSecretsResp httpmodels.ListSecretsResp

	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/list", listSecretsReq, nil)
	if err != nil {
		return listSecretsResp, err
	}

	err = json.Unmarshal(resp, &listSecretsResp)
	return listSecretsResp, err
}

// PutSecretACL creates or overwrites the ACL associated with the given principal (user or group) on the specified scope point
func (a SecretsAPI) PutSecretACL(putSecretACLReq httpmodels.PutSecretACLReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/acls/put", putSecretACLReq, nil)
	return err
}

// DeleteSecretACL deletes the given ACL on the given scope
func (a SecretsAPI) DeleteSecretACL(deleteSecretACLReq httpmodels.DeleteSecretACLReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/acls/delete", deleteSecretACLReq, nil)
	return err
}

// GetSecretACL describe the details about the given ACL, such as the group and permission
func (a SecretsAPI) GetSecretACL(getSecretACLReq httpmodels.GetSecretACLReq) (httpmodels.GetSecretACLResp, error) {
	var getSecretACLResp httpmodels.GetSecretACLResp

	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/acls/get", getSecretACLReq, nil)
	if err != nil {
		return getSecretACLResp, err
	}

	err = json.Unmarshal(resp, &getSecretACLResp)
	return getSecretACLResp, err
}

// ListSecretACLs lists the ACLs set on the given scope
func (a SecretsAPI) ListSecretACLs(listSecretACLsReq httpmodels.ListSecretACLsReq) (httpmodels.ListSecretACLsResp, error) {
	var listSecretACLsResp httpmodels.ListSecretACLsResp

	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/acls/list", listSecretACLsReq, nil)
	if err != nil {
		return listSecretACLsResp, err
	}

	err = json.Unmarshal(resp, &listSecretACLsResp)
	return listSecretACLsResp, err
}
