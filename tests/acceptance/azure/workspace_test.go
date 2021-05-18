package azure_test

import (
	"testing"

	databricks "github.com/polar-rams/databricks-sdk-golang"
	dbAzure "github.com/polar-rams/databricks-sdk-golang/azure"
	"github.com/polar-rams/databricks-sdk-golang/azure/workspace/httpmodels"
)

func TestAzureListWorkspace(t *testing.T) {
	// Initialize the test
	var o databricks.DBClientOption

	testConfig := GetTestConfig()

	o.Host = testConfig[DATABRICKS_HOST_KEY]
	o.Token = testConfig[DATABRICKS_TOKEN_KEY]

	var c dbAzure.DBClient
	c.Init(o)

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
