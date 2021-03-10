package azure

import (
	"encoding/json"
	"net/http"

	"github.com/xinsnake/databricks-sdk-golang/azure/jobs/httpmodels"
	"github.com/xinsnake/databricks-sdk-golang/azure/jobs/models"
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
func (a JobsAPI) Create(jobSettings models.JobSettings) (models.Job, error) {
	var job models.Job

	resp, err := a.Client.performQuery(http.MethodPost, "/jobs/create", jobSettings, nil)
	if err != nil {
		return job, err
	}

	err = json.Unmarshal(resp, &job)
	return job, err
}

// List lists all jobs
func (a JobsAPI) List() ([]models.Job, error) {
	var jobsList httpmodels.JobsListResp

	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/list", nil, nil)
	if err != nil {
		return jobsList.Jobs, err
	}

	err = json.Unmarshal(resp, &jobsList)
	return jobsList.Jobs, err
}

// Delete deletes a job by ID
func (a JobsAPI) Delete(jobID int64) error {
	data := httpmodels.GenericJobReq{JobID: jobID}
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/delete", data, nil)
	return err
}

// Get gets a job by ID
func (a JobsAPI) Get(jobID int64) (models.Job, error) {
	var job models.Job

	data := httpmodels.GenericJobReq{JobID: jobID}
	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/get", data, nil)
	if err != nil {
		return job, err
	}

	err = json.Unmarshal(resp, &job)
	return job, err
}

// Reset overwrites job settings
func (a JobsAPI) Reset(jobID int64, jobSettings models.JobSettings) error {
	data := httpmodels.GenericJobRunsUpdateReq{
		JobID:       jobID,
		NewSettings: jobSettings,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/reset", data, nil)
	return err
}

// Update adds, changes, or removes specific settings of an existing job
func (a JobsAPI) Update(jobID int64, jobSettings models.JobSettings, fieldsToRemove []string) error {
	data := httpmodels.GenericJobRunsUpdateReq{
		JobID:          jobID,
		NewSettings:    jobSettings,
		FieldsToRemove: fieldsToRemove,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/update", data, nil)
	return err
}

// RunNow runs a job now and return the run_id of the triggered run
func (a JobsAPI) RunNow(jobID int64, runParameters models.RunParameters) (models.Run, error) {
	var run models.Run

	data := httpmodels.JobsRunNowReq{
		JobID:         jobID,
		RunParameters: runParameters,
	}
	resp, err := a.Client.performQuery(http.MethodPost, "/jobs/run-now", data, nil)
	if err != nil {
		return run, err
	}

	err = json.Unmarshal(resp, &run)
	return run, err
}

// RunsSubmit submit a one-time run
func (a JobsAPI) RunsSubmit(runName string, clusterSpec models.ClusterSpec, jobTask models.JobTask, timeoutSeconds int32, idempotencyToken string) (models.Run, error) {
	var run models.Run

	data := httpmodels.JobsRunsSubmitReq{
		RunName:          runName,
		ClusterSpec:      clusterSpec,
		JobTask:          jobTask,
		TimeoutSeconds:   timeoutSeconds,
		IdempotencyToken: idempotencyToken,
	}
	resp, err := a.Client.performQuery(http.MethodPost, "/jobs/runs/submit", data, nil)
	if err != nil {
		return run, err
	}

	err = json.Unmarshal(resp, &run)
	return run, err
}

// RunsList lists runs from most recently started to least
func (a JobsAPI) RunsList(activeOnly, completedOnly bool, jobID int64, offset, limit int32, runType string) (httpmodels.JobsRunsListResp, error) {
	var runlistResponse httpmodels.JobsRunsListResp

	data := httpmodels.JobsRunsListReq{
		ActiveOnly:    activeOnly,
		CompletedOnly: completedOnly,
		JobID:         jobID,
		Offset:        offset,
		Limit:         limit,
		RunType:       runType,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/list", data, nil)
	if err != nil {
		return runlistResponse, err
	}

	err = json.Unmarshal(resp, &runlistResponse)
	return runlistResponse, err
}

// RunsGet retrieve the metadata of a run
func (a JobsAPI) RunsGet(runID int64) (models.Run, error) {
	var run models.Run

	data := httpmodels.GenericRunReq{RunID: runID}
	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/get", data, nil)
	if err != nil {
		return run, err
	}

	err = json.Unmarshal(resp, &run)
	return run, err
}

// RunsExport exports and retrieve the job run task
func (a JobsAPI) RunsExport(runID int64, viewsToExport models.ViewsToExport) ([]models.ViewItem, error) {
	var viewItemsView = struct {
		Views []models.ViewItem `json:"views,omitempty" url:"views,omitempty"`
	}{}

	data := struct {
		RunID         int64                `json:"run_id,omitempty" url:"run_id,omitempty"`
		ViewsToExport models.ViewsToExport `json:"views_to_export,omitempty" url:"views_to_export,omitempty"`
	}{
		runID,
		viewsToExport,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/export", data, nil)
	if err != nil {
		return viewItemsView.Views, err
	}

	err = json.Unmarshal(resp, &viewItemsView)
	return viewItemsView.Views, err
}

// RunsCancel cancels a run
func (a JobsAPI) RunsCancel(runID int64) error {
	data := httpmodels.GenericRunReq{RunID: runID}
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/runs/cancel", data, nil)
	return err
}

// RunsGetOutput retrieves the output of a run
func (a JobsAPI) RunsGetOutput(runID int64) (httpmodels.JobsRunsGetOutputResp, error) {
	var runsGetOutputResponse httpmodels.JobsRunsGetOutputResp

	data := httpmodels.GenericRunReq{RunID: runID}
	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/get-output", data, nil)
	if err != nil {
		return runsGetOutputResponse, err
	}

	err = json.Unmarshal(resp, &runsGetOutputResponse)
	return runsGetOutputResponse, err
}

// RunsDelete deletes a non-active run. Returns an error if the run is active.
func (a JobsAPI) RunsDelete(runID int64) error {
	data := httpmodels.GenericRunReq{RunID: runID}
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/runs/delete", data, nil)
	return err
}
