package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/jobs/httpmodels"
)

// JobsAPI exposes Jobs API endpoints
type JobsAPI struct {
	Client DBClient
}

func (a JobsAPI) init(client DBClient) JobsAPI {
	a.Client = client
	return a
}

// Create creates a new job
func (a JobsAPI) Create(createRequest httpmodels.CreateReq) (httpmodels.CreateResp, error) {
	var createResp httpmodels.CreateResp

	resp, err := a.Client.performQuery(http.MethodPost, "/jobs/create", createRequest, nil)
	if err != nil {
		return createResp, err
	}

	err = json.Unmarshal(resp, &createResp)
	return createResp, err
}

// List lists all jobs
func (a JobsAPI) List() (httpmodels.ListResp, error) {
	var listResp httpmodels.ListResp

	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/list", nil, nil)
	if err != nil {
		return listResp, err
	}

	err = json.Unmarshal(resp, &listResp)
	return listResp, err
}

// Delete deletes a job by ID
func (a JobsAPI) Delete(deleteReq httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/delete", deleteReq, nil)
	return err
}

// Get gets a job by ID
func (a JobsAPI) Get(getReq httpmodels.GetReq) (httpmodels.GetResp, error) {
	var getResp httpmodels.GetResp

	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/get", getReq, nil)
	if err != nil {
		return getResp, err
	}

	err = json.Unmarshal(resp, &getResp)
	return getResp, err
}

// Reset overwrites job settings
func (a JobsAPI) Reset(resetReq httpmodels.ResetReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/reset", resetReq, nil)
	return err
}

// Update adds, changes, or removes specific settings of an existing job
func (a JobsAPI) Update(updateReq httpmodels.UpdateReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/update", updateReq, nil)
	return err
}

// RunNow runs a job now and return the run_id of the triggered run
func (a JobsAPI) RunNow(runNowReq httpmodels.RunNowReq) (httpmodels.RunNowResp, error) {
	var runNowResp httpmodels.RunNowResp

	resp, err := a.Client.performQuery(http.MethodPost, "/jobs/run-now", runNowReq, nil)
	if err != nil {
		return runNowResp, err
	}

	err = json.Unmarshal(resp, &runNowResp)
	return runNowResp, err
}

// RunsSubmit submit a one-time run
func (a JobsAPI) RunsSubmit(runsSubmitReq httpmodels.RunsSubmitReq) (httpmodels.RunsSubmitResp, error) {
	var runsSubmitResp httpmodels.RunsSubmitResp

	resp, err := a.Client.performQuery(http.MethodPost, "/jobs/runs/submit", runsSubmitReq, nil)
	if err != nil {
		return runsSubmitResp, err
	}

	err = json.Unmarshal(resp, &runsSubmitResp)
	return runsSubmitResp, err
}

// RunsList lists runs from most recently started to least
func (a JobsAPI) RunsList(runsListReq httpmodels.RunsListReq) (httpmodels.RunsListResp, error) {
	var runsListResp httpmodels.RunsListResp

	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/list", runsListReq, nil)
	if err != nil {
		return runsListResp, err
	}

	err = json.Unmarshal(resp, &runsListResp)
	return runsListResp, err
}

// RunsGet retrieve the metadata of a run
func (a JobsAPI) RunsGet(runsGetReq httpmodels.RunsGetReq) (httpmodels.RunsGetResp, error) {
	var runsGetResp httpmodels.RunsGetResp

	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/get", runsGetReq, nil)
	if err != nil {
		return runsGetResp, err
	}

	err = json.Unmarshal(resp, &runsGetResp)
	return runsGetResp, err
}

// RunsExport exports and retrieve the job run task
func (a JobsAPI) RunsExport(runsExportReq httpmodels.RunsExportReq) (httpmodels.RunsExportResp, error) {
	var runsExportResp httpmodels.RunsExportResp

	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/export", runsExportReq, nil)
	if err != nil {
		return runsExportResp, err
	}

	err = json.Unmarshal(resp, &runsExportResp)
	return runsExportResp, err
}

// RunsCancel cancels a run
func (a JobsAPI) RunsCancel(runsCancelReq httpmodels.RunsCancelReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/runs/cancel", runsCancelReq, nil)
	return err
}

// RunsGetOutput retrieves the output of a run
func (a JobsAPI) RunsGetOutput(runsGetOutputReq httpmodels.RunsGetOutputReq) (httpmodels.RunsGetOutputResp, error) {
	var runsGetOutputResp httpmodels.RunsGetOutputResp

	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/get-output", runsGetOutputReq, nil)
	if err != nil {
		return runsGetOutputResp, err
	}

	err = json.Unmarshal(resp, &runsGetOutputResp)
	return runsGetOutputResp, err
}

// RunsDelete deletes a non-active run. Returns an error if the run is active.
func (a JobsAPI) RunsDelete(runsDeleteReq httpmodels.RunsDeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/runs/delete", runsDeleteReq, nil)
	return err
}
