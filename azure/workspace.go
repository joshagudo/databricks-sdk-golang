package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/workspace/httpmodels"
)

// WorkspaceAPI exposes the Workspace API
type WorkspaceAPI struct {
	Client DBClient
}

func (a WorkspaceAPI) init(client DBClient) WorkspaceAPI {
	a.Client = client
	return a
}

// Delete an object or a directory (and optionally recursively deletes all objects in the directory)
func (a WorkspaceAPI) Delete(deleteReq httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/workspace/delete", deleteReq, nil)
	return err
}

// Export a notebook or contents of an entire directory
func (a WorkspaceAPI) Export(exportReq httpmodels.ExportReq) (httpmodels.ExportResp, error) {
	var exportResp httpmodels.ExportResp

	resp, err := a.Client.performQuery(http.MethodGet, "/workspace/export", exportReq, nil)
	if err != nil {
		return exportResp, err
	}

	err = json.Unmarshal(resp, &exportResp)
	return exportResp, err
}

// Gets the status of an object or a directory
func (a WorkspaceAPI) GetStatus(getStatusReq httpmodels.GetStatusReq) (httpmodels.GetStatusResp, error) {
	var getStatusResp httpmodels.GetStatusResp

	resp, err := a.Client.performQuery(http.MethodGet, "/workspace/get-status", getStatusReq, nil)
	if err != nil {
		return getStatusResp, err
	}

	err = json.Unmarshal(resp, &getStatusResp)
	return getStatusResp, err
}

// Import a notebook or the contents of an entire directory
func (a WorkspaceAPI) Import(importReq httpmodels.ImportReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/workspace/import", importReq, nil)
	return err
}

// List lists the contents of a directory, or the object if it is not a directory
func (a WorkspaceAPI) List(listReq httpmodels.ListReq) (httpmodels.ListResp, error) {
	var listResp httpmodels.ListResp

	resp, err := a.Client.performQuery(http.MethodGet, "/workspace/list", listReq, nil)
	if err != nil {
		return listResp, err
	}

	err = json.Unmarshal(resp, &listResp)
	return listResp, err
}

// Mkdirs creates the given directory and necessary parent directories if they do not exists
func (a WorkspaceAPI) Mkdirs(mkdirsReq httpmodels.MkdirsReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/workspace/mkdirs", mkdirsReq, nil)
	return err
}
