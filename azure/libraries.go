package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/libraries/httpmodels"
)

// LibrariesAPI exposes the Libraries API
type LibrariesAPI struct {
	Client DBClient
}

func (a LibrariesAPI) init(client DBClient) LibrariesAPI {
	a.Client = client
	return a
}

// AllClusterStatuses gets the status of all libraries on all clusters
func (a LibrariesAPI) AllClusterStatuses() (httpmodels.AllClusterStatusesResp, error) {
	var allClusterStatusesResp httpmodels.AllClusterStatusesResp

	resp, err := a.Client.performQuery(http.MethodGet, "/libraries/all-cluster-statuses", nil, nil)
	if err != nil {
		return allClusterStatusesResp, err
	}

	err = json.Unmarshal(resp, &allClusterStatusesResp)
	return allClusterStatusesResp, err
}

// ClusterStatus get the status of libraries on a cluster
func (a LibrariesAPI) ClusterStatus(clusterStatusReq httpmodels.ClusterStatusReq) (httpmodels.ClusterStatusResp, error) {
	var clusterStatusResp httpmodels.ClusterStatusResp

	resp, err := a.Client.performQuery(http.MethodGet, "/libraries/cluster-status", clusterStatusReq, nil)
	if err != nil {
		return clusterStatusResp, err
	}

	err = json.Unmarshal(resp, &clusterStatusResp)
	return clusterStatusResp, err
}

// Install installs libraries on a cluster
func (a LibrariesAPI) Install(installReq httpmodels.InstallReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/libraries/install", installReq, nil)
	return err
}

// Uninstall sets libraries to be uninstalled on a cluster
func (a LibrariesAPI) Uninstall(uninstallReq httpmodels.UninstallReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/libraries/uninstall", uninstallReq, nil)
	return err
}
