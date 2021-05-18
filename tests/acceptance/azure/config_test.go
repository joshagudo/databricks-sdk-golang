package azure_test

import (
	"os"

	databricks "github.com/polar-rams/databricks-sdk-golang"
	dbAzure "github.com/polar-rams/databricks-sdk-golang/azure"
)

var c *dbAzure.DBClient

const (
	DATABRICKS_HOST_KEY  = "DATABRICKS_HOST"
	DATABRICKS_TOKEN_KEY = "DATABRICKS_TOKEN"
)

func init() {

	opt := databricks.NewDBClientOption("", "", os.Getenv(DATABRICKS_HOST_KEY), os.Getenv(DATABRICKS_TOKEN_KEY), nil, false, 0)
	c = dbAzure.NewDBClient(opt)
}
