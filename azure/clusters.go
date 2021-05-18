package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/clusters/httpmodels"
)

// ClustersAPI exposes the Clusters API
type ClustersAPI struct {
	Client DBClient
}

func (a ClustersAPI) init(client DBClient) ClustersAPI {
	a.Client = client
	return a
}

// Create creates a new Spark cluster
func (a ClustersAPI) Create(cluster httpmodels.CreateReq) (httpmodels.CreateResp, error) {
	var createResp httpmodels.CreateResp

	resp, err := a.Client.performQuery(http.MethodPost, "/clusters/create", cluster, nil)
	if err != nil {
		return createResp, err
	}

	err = json.Unmarshal(resp, &createResp)
	return createResp, err
}

// Edit edits the configuration of a cluster to match the provided attributes and size
func (a ClustersAPI) Edit(editReq httpmodels.EditReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/edit", editReq, nil)
	return err
}

// Start starts a terminated Spark cluster given its ID
func (a ClustersAPI) Start(startReq httpmodels.StartReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/start", startReq, nil)
	return err
}

// Restart restart a Spark cluster given its ID. If the cluster is not in a RUNNING state, nothing will happen.
func (a ClustersAPI) Restart(restartReq httpmodels.RestartReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/restart", restartReq, nil)
	return err
}

// Resize resizes a cluster to have a desired number of workers. This will fail unless the cluster is in a RUNNING state.
func (a ClustersAPI) Resize(resizeReq httpmodels.ResizeReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/resize", resizeReq, nil)
	return err
}

// Delete terminates a Spark cluster given its ID
func (a ClustersAPI) Delete(deleteReq httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/delete", deleteReq, nil)
	return err
}

// Terminate is an alias of Delete
func (a ClustersAPI) Terminate(deleteReq httpmodels.DeleteReq) error {
	return a.Delete(deleteReq)
}

// PermanentDelete permanently delete a cluster
func (a ClustersAPI) PermanentDelete(permDelReq httpmodels.PermanentDeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/permanent-delete", permDelReq, nil)
	return err
}

// Get retrieves the information for a cluster given its identifier
func (a ClustersAPI) Get(getReq httpmodels.GetReq) (httpmodels.GetResp, error) {
	var clusterInfo httpmodels.GetResp

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/get", getReq, nil)
	if err != nil {
		return clusterInfo, err
	}

	err = json.Unmarshal(resp, &clusterInfo)
	return clusterInfo, err
}

// Pin ensure that an interactive cluster configuration is retained even after a cluster has been terminated for more than 30 days
func (a ClustersAPI) Pin(pinReq httpmodels.PinReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/pin", pinReq, nil)
	return err
}

// Unpin allows the cluster to eventually be removed from the list returned by the List API
func (a ClustersAPI) Unpin(unpinReq httpmodels.UnpinReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/unpin", unpinReq, nil)
	return err
}

// List return information about all pinned clusters, currently active clusters,
// up to 70 of the most recently terminated interactive clusters in the past 30 days,
// and up to 30 of the most recently terminated job clusters in the past 30 days
func (a ClustersAPI) List() (httpmodels.ListResp, error) {
	var listResp httpmodels.ListResp

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/list", nil, nil)
	if err != nil {
		return listResp, err
	}

	err = json.Unmarshal(resp, &listResp)
	return listResp, err
}

// ListNodeTypes returns a list of supported Spark node types
func (a ClustersAPI) ListNodeTypes() (httpmodels.ListNodeTypesResp, error) {
	var listNodeTypesResp httpmodels.ListNodeTypesResp

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/list-node-types", nil, nil)
	if err != nil {
		return listNodeTypesResp, err
	}

	err = json.Unmarshal(resp, &listNodeTypesResp)
	return listNodeTypesResp, err
}

// RuntimeVersions return the list of available Runtime versions
func (a ClustersAPI) RuntimeVersions() (httpmodels.RuntimeVersionsResp, error) {
	var runtimeVersionsResp httpmodels.RuntimeVersionsResp

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/spark-versions", nil, nil)
	if err != nil {
		return runtimeVersionsResp, err
	}

	err = json.Unmarshal(resp, &runtimeVersionsResp)
	return runtimeVersionsResp, err
}

// Events retrieves a list of events about the activity of a cluster
func (a ClustersAPI) Events(eventReq httpmodels.EventsReq) (httpmodels.EventsResp, error) {
	var eventsResponse httpmodels.EventsResp

	resp, err := a.Client.performQuery(http.MethodPost, "/clusters/events", eventReq, nil)
	if err != nil {
		return eventsResponse, err
	}

	err = json.Unmarshal(resp, &eventsResponse)
	return eventsResponse, err
}
