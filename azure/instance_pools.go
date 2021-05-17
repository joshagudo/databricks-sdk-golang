package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/instance_pools/httpmodels"
)

// InstancePoolsAPI exposes the InstancePools API
type InstancePoolsAPI struct {
	Client DBClient
}

func (a InstancePoolsAPI) init(client DBClient) InstancePoolsAPI {
	a.Client = client
	return a
}

// Create creates an instance pool
func (a InstancePoolsAPI) Create(instancePool httpmodels.CreateReq) (httpmodels.CreateResp, error) {
	var createResp httpmodels.CreateResp

	resp, err := a.Client.performQuery(http.MethodPost, "/instance-pools/create", instancePool, nil)
	if err != nil {
		return createResp, err
	}

	err = json.Unmarshal(resp, &createResp)
	return createResp, err
}

// Edit modifies the configuration of an existing instance pool.
func (a InstancePoolsAPI) Edit(editReq httpmodels.EditReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/instance-pools/edit", editReq, nil)
	return err
}

// Delete permanently deletes the instance pool.
func (a InstancePoolsAPI) Delete(deleteReq httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/instance-pools/delete", deleteReq, nil)
	return err
}

// Get retrieves the information for an instance pool given its identifier.
func (a InstancePoolsAPI) Get(getReq httpmodels.GetReq) (httpmodels.GetResp, error) {
	var getResp httpmodels.GetResp

	resp, err := a.Client.performQuery(http.MethodGet, "/instance-pools/get", getReq, nil)
	if err != nil {
		return getResp, err
	}

	err = json.Unmarshal(resp, &getResp)
	return getResp, err
}

// List returns information for all instance pools.
func (a InstancePoolsAPI) List() (httpmodels.ListResp, error) {
	var listResp httpmodels.ListResp

	resp, err := a.Client.performQuery(http.MethodGet, "/instance-pools/list", nil, nil)
	if err != nil {
		return listResp, err
	}

	err = json.Unmarshal(resp, &listResp)
	return listResp, err
}
