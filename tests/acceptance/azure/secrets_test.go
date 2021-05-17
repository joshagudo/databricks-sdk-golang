package azure_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	databricks "github.com/polar-rams/databricks-sdk-golang"
	dbAzure "github.com/polar-rams/databricks-sdk-golang/azure"
	// "testing"
	// "github.com/onsi/gomega"
)

func TestAzureCreateSecretsScope(t *testing.T) {
	var o databricks.DBClientOption
	o.Host = os.Getenv("DATABRICKS_HOST")
	o.Token = os.Getenv("DATABRICKS_TOKEN")

	var c dbAzure.DBClient
	c.Init(o)

	scopeName := "testSecretScope12456"

	e := c.Secrets().CreateSecretScope(scopeName, "users")

	if e != nil {
		log.Printf("Error : %v", e)
		t.Error("could not create scope")
	}

	sc, er := c.Secrets().ListSecretScopes()

	if er != nil {
		log.Printf("Error : %v", er)
		t.Error("could not list scope")
	}

	ssExists := false

	for _, sesc := range sc {
		if sesc.Name == scopeName {
			ssExists = true
			break
		}
		fmt.Printf("sesc: %s\n", sesc.Name)
	}

	if !ssExists {
		t.Errorf("expected secret scope %s to exist", scopeName)
	}

	e = c.Secrets().DeleteSecretScope(scopeName)
	if e != nil {
		log.Printf("Error : %v", e)
		t.Error("could not delete secret scope")
	}

}
