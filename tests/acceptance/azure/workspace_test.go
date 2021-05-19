package azure_test

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/polar-rams/databricks-sdk-golang/azure/workspace/httpmodels"
	"github.com/polar-rams/databricks-sdk-golang/azure/workspace/models"
	"github.com/stretchr/testify/assert"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func TestAzureWorkspaceList(t *testing.T) {
	// list and assert
	const rootPath = "/"
	listRequest := httpmodels.ListReq{
		Path: rootPath,
	}
	workSpaceList, e := c.Workspace().List(listRequest)
	assert.Nil(t, e, "could not list workspaces")
	assert.NotEqual(t, len(*workSpaceList.Objects), 0, "No workspaces found in root path")
}

func TestAzureWorkspaceImportAndDelete(t *testing.T) {
	// import and assert
	var scalaLanguage models.Language = models.LanguageScala
	var sourceFormat models.ExportFormat = models.ExportFormatSource
	var samplePath = fmt.Sprintf("/samplenotebook-%s", randSeq(5))
	importRequestScala := httpmodels.ImportReq{
		Path:      samplePath,
		Language:  scalaLanguage,
		Format:    sourceFormat,
		Content:   "MSsx\n",
		Overwrite: true,
	}
	assert.Nil(t, c.Workspace().Import(importRequestScala), fmt.Sprintf("could not import request: %s", importRequestScala.Language))

	// delete and assert
	deleteRequest := httpmodels.DeleteReq{
		Path:      samplePath,
		Recursive: true,
	}
	assert.Nil(t, c.Workspace().Delete(deleteRequest), fmt.Sprintf("could not delete the imported resquest: %s", importRequestScala.Language))
}

func TestAzureWorkspaceExport(t *testing.T) {
	// import and assert
	var pythonLanguage models.Language = models.LanguagePython
	var sourceFormat models.ExportFormat = models.ExportFormatSource
	var samplePath = fmt.Sprintf("/samplenotebook-%s", randSeq(5))
	importRequestPython := httpmodels.ImportReq{
		Path:      samplePath,
		Language:  pythonLanguage,
		Format:    sourceFormat,
		Content:   "MSsx\n",
		Overwrite: true,
	}
	assert.Nil(t, c.Workspace().Import(importRequestPython), fmt.Sprintf("could not import request for test: %s", importRequestPython.Language))

	// Export and assert
	exportRequest := httpmodels.ExportReq{
		Path:   samplePath,
		Format: sourceFormat,
	}
	exportResponse, e := c.Workspace().Export(exportRequest)
	assert.Nil(t, e, fmt.Sprintf("could not export notebook: %s", importRequestPython.Path))
	strImportReq, e := base64.StdEncoding.DecodeString(importRequestPython.Content)
	assert.Nil(t, e, fmt.Sprintf("could not decode import request content: %s", importRequestPython.Content))
	strExportRes, e := base64.StdEncoding.DecodeString(exportResponse.Content)
	assert.Nil(t, e, fmt.Sprintf("could not decode export response content: %s", exportResponse.Content))
	assert.Equal(t, true, strings.Contains(string(strExportRes), string(strImportReq)), fmt.Sprintf("Imported content: %s\t Exported content: %s. Are not equal.", importRequestPython.Content, exportResponse.Content))

	// delete and assert
	deleteRequest := httpmodels.DeleteReq{
		Path:      samplePath,
		Recursive: true,
	}
	assert.Nil(t, c.Workspace().Delete(deleteRequest), fmt.Sprintf("could not delete the imported resquest: %s", importRequestPython.Language))
}
