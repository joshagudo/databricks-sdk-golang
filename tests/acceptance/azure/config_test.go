package azure_test

import "os"

var testConfig map[string]string

const (
	DATABRICKS_HOST_KEY  = "DATABRICKS_HOST"
	DATABRICKS_TOKEN_KEY = "DATABRICKS_TOKEN"
)

func init() {
	testConfig = make(map[string]string)
	testConfig[DATABRICKS_HOST_KEY] = os.Getenv(DATABRICKS_HOST_KEY)
	testConfig[DATABRICKS_TOKEN_KEY] = os.Getenv(DATABRICKS_TOKEN_KEY)
}

func GetTestConfig() map[string]string {
	return testConfig
}
