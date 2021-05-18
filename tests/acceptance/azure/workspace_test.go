package azure_test

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/polar-rams/databricks-sdk-golang/azure/workspace/httpmodels"
	"github.com/polar-rams/databricks-sdk-golang/azure/workspace/models"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func TestAzureWorkspaceList(t *testing.T) {
	const rootPath = "/"

	// Listing default workspaces
	listRequest := httpmodels.ListReq{
		Path: rootPath,
	}

	workSpaceList, e := c.Workspace().List(listRequest)

	if e != nil {
		t.Error("could not list workspaces")
	}

	if len(workSpaceList.Objects) == 0 {
		t.Error("could not retrieve default workspaces")
	}
}

func TestAzureWorkspaceImportAndDelete(t *testing.T) {
	// Creating object to import to workspace
	var scalaLanguage models.Language = models.LanguageScala
	var sourceFormat models.ExportFormat = models.ExportFormatSource
	var samplePath = fmt.Sprintf("/samplenotebook-%s", randSeq(5))
	importRequestScala := httpmodels.ImportReq{
		Path:      samplePath,
		Language:  &scalaLanguage,
		Format:    &sourceFormat,
		Content:   "MSsx\n",
		Overwrite: true,
	}

	if e := c.Workspace().Import(importRequestScala); e != nil {
		t.Errorf("could not import request: %s", *importRequestScala.Language)
	}

	deleteRequest := httpmodels.DeleteReq{
		Path:      samplePath,
		Recursive: true,
	}
	if e := c.Workspace().Delete(deleteRequest); e != nil {
		t.Errorf("could not delete the imported resquest: %s", *importRequestScala.Language)
	}
}

func TestAzureWorkspaceExport(t *testing.T) {
	// Creating object to import to workspace
	var pythonLanguage models.Language = models.LanguagePython
	var sourceFormat models.ExportFormat = models.ExportFormatSource
	var samplePath = fmt.Sprintf("/samplenotebook-%s", randSeq(5))
	importRequestPython := httpmodels.ImportReq{
		Path:      samplePath,
		Language:  &pythonLanguage,
		Format:    &sourceFormat,
		Content:   "MSsx\n",
		Overwrite: true,
	}

	if e := c.Workspace().Import(importRequestPython); e != nil {
		t.Errorf("could not import request for test: %s", *importRequestPython.Language)
	}

	exportRequest := httpmodels.ExportReq{
		Path:   samplePath,
		Format: &sourceFormat,
	}

	exportResponse, e := c.Workspace().Export(exportRequest)
	if e != nil {
		t.Errorf("could not export notebook: %s", importRequestPython.Path)
	}

	strImportReq, err := base64.StdEncoding.DecodeString(importRequestPython.Content)
	if err != nil {
		t.Errorf("could not decode import request content: %s", importRequestPython.Content)
	}
	strExportRes, err := base64.StdEncoding.DecodeString(exportResponse.Content)
	if err != nil {
		t.Errorf("could not decode export response content: %s", exportResponse.Content)
	}

	present := strings.Contains(string(strExportRes), string(strImportReq))

	if !present {
		t.Errorf("Imported content: %s\t Exported content: %s. Are not equal.", importRequestPython.Content, exportResponse.Content)
	}

	deleteRequest := httpmodels.DeleteReq{
		Path:      samplePath,
		Recursive: true,
	}
	if e := c.Workspace().Delete(deleteRequest); e != nil {
		t.Errorf("could not delete the imported resquest: %s", *importRequestPython.Language)
	}
}

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
