package azure_test

import (
	"fmt"
	"math/rand"
	"testing"

	databricks "github.com/polar-rams/databricks-sdk-golang"
	dbAzure "github.com/polar-rams/databricks-sdk-golang/azure"
)

func TestAzureCreateSecretsScope(t *testing.T) {
	var o databricks.DBClientOption

	testConfig := GetTestConfig()
	o.Host = testConfig[DATABRICKS_HOST_KEY]
	o.Token = testConfig[DATABRICKS_TOKEN_KEY]

	var c dbAzure.DBClient
	c.Init(o)

	const (
		scpNmPre = "testSecrScpName"
		scpGrp   = "users"
	)

	scpNmSuf := fmt.Sprintf("%08d", rand.Intn(100000000))
	scpNm := fmt.Sprintf("%s%s", scpNmPre, scpNmSuf)

	if e := c.Secrets().CreateSecretScope(scpNm, scpGrp); e != nil {
		t.Error("could not create scope")
	}

	sc, e := c.Secrets().ListSecretScopes()
	if e != nil {
		t.Error("could not list scopes")
	}
	ssExists := false

	for _, sesc := range sc {
		if sesc.Name == scpNm {
			ssExists = true
			break
		}
	}

	if !ssExists {
		t.Errorf("expected secret scope %s to exist", scpNm)
	}

	e = c.Secrets().DeleteSecretScope(scpNm)
	if e != nil {
		t.Error("could not delete secret scope")
	}

}
