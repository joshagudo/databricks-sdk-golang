package azure_test

import (
	"testing"

	"github.com/polar-rams/databricks-sdk-golang/azure/dbfs/httpmodels"
)

func TestAzureDBFSList(t *testing.T) {
	listRequest := httpmodels.ListReq{
		Path: "/",
	}
	listResponse, e := c.Dbfs().List(listRequest)
	if e != nil {
		t.Error("could not list dbfs")
	}

	if len(listResponse.Files) == 0 {
		t.Error("could not get any files from dbfs")
	}
}
