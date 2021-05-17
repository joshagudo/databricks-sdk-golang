package azure_test

import (
	"log"
	"os"
	"testing"

	databricks "github.com/polar-rams/databricks-sdk-golang"
	dbAzure "github.com/polar-rams/databricks-sdk-golang/azure"
	// "github.com/onsi/gomega"
)

func TestAzureListWorkspace(t *testing.T) {
	var o databricks.DBClientOption
	o.Host = os.Getenv("DATABRICKS_HOST")
	o.Token = os.Getenv("DATABRICKS_TOKEN")

	var c dbAzure.DBClient
	c.Init(o)

	wss, e := c.Workspace().List("/")

	if e != nil {
		log.Printf("Error: %s", e)
	}

	if wss[0].Path != "/Users" || wss[1].Path != "/Shared" {
		t.Error("Default root workspaces are not /Users and /Shared")
	}
}

func TestAzureImportWorkspace(t *testing.T) {
	var o databricks.DBClientOption
	o.Host = os.Getenv("DATABRICKS_HOST")
	o.Token = os.Getenv("DATABRICKS_TOKEN")

	var c dbAzure.DBClient
	c.Init(o)

	wss, e := c.Workspace().List("/")

	if e != nil {
		log.Printf("Error: %s", e)
	}

	if wss[0].Path != "/Users" || wss[1].Path != "/Shared" {
		t.Error("Default root workspaces are not /Users and /Shared")
	}
}
