package azure_test

import (
	"testing"
	"time"

	"github.com/polar-rams/databricks-sdk-golang/azure/jobs/httpmodels"
	"github.com/polar-rams/databricks-sdk-golang/azure/jobs/models"
)

func TestAzureJobsJob(t *testing.T) {
	j := c.Jobs()

	jobName := "my-test-job-name"

	// Create
	createReq := httpmodels.CreateReq{
		Name: jobName,
		NewCluster: &models.NewCluster{
			SparkVersion: "7.3.x-scala2.12",
			NodeTypeID:   "Standard_D3_v2",
			NumWorkers:   1,
		},
	}
	createResp, err := j.Create(createReq)
	if err != nil {
		t.Errorf("unable to create job %s", err.Error())
	}
	time.Sleep(3 * time.Second)

	// List
	listResp, err := j.List()
	if len(listResp.Jobs) < 1 {
		t.Errorf("unable to list jobs %s", err.Error())
	}
	time.Sleep(1 * time.Second)

	// Get
	jobID := createResp.JobID
	getReq := httpmodels.GetReq{
		JobID: jobID,
	}
	getResp, err := j.Get(getReq)
	if err != nil {
		t.Errorf("unable to get job %s", err.Error())
	}

	if getResp.Settings.Name != jobName {
		t.Errorf("unable to get valid job %s", err.Error())
	}
}

func TestAzureJobsRun(t *testing.T) {
}
