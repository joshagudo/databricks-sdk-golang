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
func (a ClustersAPI) Start(startReq httpmodels.ClusterReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/start", startReq, nil)
	return err
}

// Restart restart a Spark cluster given its ID. If the cluster is not in a RUNNING state, nothing will happen.
func (a ClustersAPI) Restart(restartReq httpmodels.ClusterReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/restart", restartReq, nil)
	return err
}

// Resize resizes a cluster to have a desired number of workers. This will fail unless the cluster is in a RUNNING state.
func (a ClustersAPI) Resize(resizeReq httpmodels.ClusterResizeReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/resize", resizeReq, nil)
	return err
}

// Terminate terminates a Spark cluster given its ID
func (a ClustersAPI) Terminate(terminateReq httpmodels.ClusterReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/delete", terminateReq, nil)
	return err
}

// Delete is an alias of Terminate
func (a ClustersAPI) Delete(deleteReq httpmodels.ClusterReq) error {
	return a.Terminate(deleteReq)
}

// PermanentDelete permanently delete a cluster
func (a ClustersAPI) PermanentDelete(permDelReq httpmodels.ClusterReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/permanent-delete", permDelReq, nil)
	return err
}

// Get retrieves the information for a cluster given its identifier
func (a ClustersAPI) Get(getReq httpmodels.ClusterReq) (httpmodels.GetResp, error) {
	var clusterInfo httpmodels.GetResp

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/get", getReq, nil)
	if err != nil {
		return clusterInfo, err
	}

	err = json.Unmarshal(resp, &clusterInfo)
	return clusterInfo, err
}

// Pin ensure that an interactive cluster configuration is retained even after a cluster has been terminated for more than 30 days
func (a ClustersAPI) Pin(pinReq httpmodels.ClusterReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/pin", pinReq, nil)
	return err
}

// Unpin allows the cluster to eventually be removed from the list returned by the List API
func (a ClustersAPI) Unpin(unpinReq httpmodels.ClusterReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/unpin", unpinReq, nil)
	return err
}

// List return information about all pinned clusters, currently active clusters,
// up to 70 of the most recently terminated interactive clusters in the past 30 days,
// and up to 30 of the most recently terminated job clusters in the past 30 days
func (a ClustersAPI) List() ([]httpmodels.GetResp, error) {
	var clusterList httpmodels.GetRespList

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/list", nil, nil)
	if err != nil {
		return clusterList.Clusters, err
	}

	err = json.Unmarshal(resp, &clusterList)
	return clusterList.Clusters, err
}

// ListNodeTypes returns a list of supported Spark node types
func (a ClustersAPI) ListNodeTypes() ([]httpmodels.ListNodeTypesRespItem, error) {
	var nodeTypeList httpmodels.ListNodeTypesRespItemList

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/list-node-types", nil, nil)
	if err != nil {
		return nodeTypeList.NodeTypes, err
	}

	err = json.Unmarshal(resp, &nodeTypeList)
	return nodeTypeList.NodeTypes, err
}

// SparkVersions return the list of available Spark versions
func (a ClustersAPI) SparkVersions() (httpmodels.SparkVersionsRespItem, error) {
	var versionsList httpmodels.SparkVersionsRespItemList

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/spark-versions", nil, nil)
	if err != nil {
		return versionsList.Versions, err
	}

	err = json.Unmarshal(resp, &versionsList)
	return versionsList.Versions, err
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
