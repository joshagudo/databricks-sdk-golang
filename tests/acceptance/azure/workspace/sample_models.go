package workspace

import (
	"github.com/polar-rams/databricks-sdk-golang/azure/workspace/httpmodels"
	"github.com/polar-rams/databricks-sdk-golang/azure/workspace/models"
)

var importRequestSamples map[string]httpmodels.ImportReq

var scalaLanguage models.Language = models.LanguageScala
var pythonLanguage models.Language = models.LanguagePython
var sqlLanguage models.Language = models.LanguageSQL
var rLanguage models.Language = models.LanguageR
var sourceFormat models.ExportFormat = models.ExportFormatSource
var htmlFormat models.ExportFormat = models.ExportFormatHtml
var jupyterFormat models.ExportFormat = models.ExportFormatJupyter
var dbcFormat models.ExportFormat = models.ExportFormatDbc

var importRequestScala = httpmodels.ImportReq{
	Path:      "/",
	Language:  &scalaLanguage,
	Format:    &sourceFormat,
	Content:   "MSsx\n",
	Overwrite: true,
}

func LoadImportRequestSamples() map[string]httpmodels.ImportReq {
	importRequestSamples = make(map[string]httpmodels.ImportReq)
	importRequestSamples["scala"] = importRequestScala
	return importRequestSamples
}
