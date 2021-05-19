package azure_test

import (
	"testing"
	"time"

	"github.com/polar-rams/databricks-sdk-golang/azure/jobs/httpmodels"
	"github.com/polar-rams/databricks-sdk-golang/azure/jobs/models"
	workspaceHTTPModels "github.com/polar-rams/databricks-sdk-golang/azure/workspace/httpmodels"
	workspaceModels "github.com/polar-rams/databricks-sdk-golang/azure/workspace/models"
	"github.com/stretchr/testify/assert"
)

var testJobNotebookPath = "/ScalaExampleNotebook"

func beforeTestAzureJobsJob(t *testing.T) {
	importReq := workspaceHTTPModels.ImportReq{
		Content:   "MSsx\n",
		Path:      testJobNotebookPath,
		Language:  workspaceModels.LanguageScala,
		Overwrite: true,
		Format:    workspaceModels.ExportFormatSource,
	}
	err := c.Workspace().Import(importReq)
	assert.Nil(t, err)
}

func TestAzureJobsJobs(t *testing.T) {
	beforeTestAzureJobsJob(t)

	jobName := "my-test-job-name"

	// Create
	createReq := httpmodels.CreateReq{
		Name: jobName,
		NewCluster: &models.NewCluster{
			SparkVersion: "7.3.x-scala2.12",
			NodeTypeID:   "Standard_D3_v2",
			NumWorkers:   1,
		},
		NotebookTask: &models.NotebookTask{
			NotebookPath: testJobNotebookPath,
		},
		MaxRetries: 1,
	}
	createResp, err := c.Jobs().Create(createReq)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)

	// List
	listResp, err := c.Jobs().List()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, 1, len(*listResp.Jobs))

	// Get
	jobID := createResp.JobID
	getReq := httpmodels.GetReq{
		JobID: jobID,
	}
	getResp, err := c.Jobs().Get(getReq)
	assert.Nil(t, err)
	assert.Equal(t, jobName, getResp.Settings.Name)
	assert.Equal(t, int32(1), getResp.Settings.MaxRetries)

	// Update
	updateReq := httpmodels.UpdateReq{
		JobID: jobID,
		NewSettings: &models.JobSettings{
			MaxRetries: 2,
		},
	}
	err = c.Jobs().Update(updateReq)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)

	getResp, err = c.Jobs().Get(getReq)
	assert.Nil(t, err)
	assert.Equal(t, int32(2), getResp.Settings.MaxRetries)

	// Reset
	resetReq := httpmodels.ResetReq{
		JobID: jobID,
		NewSettings: &models.JobSettings{
			Name: jobName,
			NewCluster: &models.NewCluster{
				SparkVersion: "7.3.x-scala2.12",
				NodeTypeID:   "Standard_D3_v2",
				NumWorkers:   1,
			},
			NotebookTask: &models.NotebookTask{
				NotebookPath: testJobNotebookPath,
			},
			MaxRetries: 1,
		},
	}
	err = c.Jobs().Reset(resetReq)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)

	getResp, err = c.Jobs().Get(getReq)
	assert.Nil(t, err)
	assert.Equal(t, int32(1), getResp.Settings.MaxRetries)

	// Delete
	deleteReq := httpmodels.DeleteReq{
		JobID: jobID,
	}
	err = c.Jobs().Delete(deleteReq)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)

	_, err = c.Jobs().Get(getReq)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "does not exist")
}

func TestAzureJobsRuns(t *testing.T) {
}
