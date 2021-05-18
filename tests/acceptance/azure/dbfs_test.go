package azure_test

import (
	"testing"

	"github.com/polar-rams/databricks-sdk-golang/azure/dbfs/httpmodels"
)

func TestDBFSList(t *testing.T) {
	const rootPath = "/"

	// Listing default workspaces
	var lReq httpmodels.ListReq
	listRequest.Path = rootPath

	workSpaceList, e := c.Workspace().List(listRequest)

	if e != nil {
		t.Error("TestAzureWorkspaceList: could not list workspaces")
	}

	if len(workSpaceList.Objects) == 0 {
		t.Error("TestAzureWorkspaceList: Could not retrieve default workspaces")
	}
}
