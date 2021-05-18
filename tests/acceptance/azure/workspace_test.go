package azure_test

import (
	"testing"

	"github.com/polar-rams/databricks-sdk-golang/azure/workspace/httpmodels"
)

func TestAzureWorkspaceList(t *testing.T) {
	c := GetTestDBClient()

	const rootPath = "/"

	// Listing default workspaces
	var listRequest httpmodels.ListReq
	listRequest.Path = rootPath

	workSpaceList, e := c.Workspace().List(listRequest)

	if e != nil {
		t.Error("could not list workspaces")
	}

	if workSpaceList.Objects[0].Path != "/Users" || workSpaceList.Objects[1].Path != "/Shared" {
		t.Error("Default root workspaces are not /Users and /Shared")
	}
}

func TestAzureWorkspaceImport(t *testing.T) {
	// Initialize the test
	c := GetTestDBClient()

	// Creating object to import to workspace
	importRequestSamples := LoadImportRequestSamples()

	c.Workspace().Import(importRequestSamples["scala"])
}
