package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/token/httpmodels"
)

// TokenAPI exposes the Token API
type TokenAPI struct {
	Client DBClient
}

func (a TokenAPI) init(client DBClient) TokenAPI {
	a.Client = client
	return a
}

// Create creates and return a token
func (a SecretsAPI) Create(request httpmodels.CreateReq) (httpmodels.CreateResp, error) {
	var createResponse httpmodels.CreateResp

	resp, err := a.Client.performQuery(http.MethodPost, "/token/create", request, nil)
	if err != nil {
		return createResponse, err
	}

	err = json.Unmarshal(resp, &createResponse)
	return createResponse, err
}

// List lists all the valid tokens for a user-workspace pair
func (a SecretsAPI) List() (httpmodels.ListResp, error) {
	var listResponse httpmodels.ListResp

	resp, err := a.Client.performQuery(http.MethodGet, "/token/list", nil, nil)
	if err != nil {
		return listResponse, err
	}

	err = json.Unmarshal(resp, &listResponse)
	return listResponse, err
}

// Revoke revokes an access token
func (a SecretsAPI) Revoke(request httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/token/delete", request, nil)
	return err
}
