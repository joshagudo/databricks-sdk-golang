package azure_test

import (
	"os"

	databricks "github.com/polar-rams/databricks-sdk-golang"
	dbAzure "github.com/polar-rams/databricks-sdk-golang/azure"
)

var testConfig map[string]string
var c dbAzure.DBClient

const (
	DATABRICKS_HOST_KEY  = "DATABRICKS_HOST"
	DATABRICKS_TOKEN_KEY = "DATABRICKS_TOKEN"
)

func init() {
	testConfig = make(map[string]string)

	var o databricks.DBClientOption
	o.Host = os.Getenv(DATABRICKS_HOST_KEY)
	o.Token = os.Getenv(DATABRICKS_TOKEN_KEY)
	c.Init(o)
}

func GetTestDBClient() *dbAzure.DBClient {
	return &c
}
