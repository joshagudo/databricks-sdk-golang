# databricks-sdk-golang

This is a Golang SDK for [DataBricks REST API 2.0](https://docs.databricks.com/api/latest/index.html#) and [Azure DataBricks REST API 2.0](https://docs.azuredatabricks.net/api/latest/index.html).

**WARNING: The SDK is unstable and under development. More testing needed!**

## Usage

```go
import (
  databricks "github.com/gpompe/databricks-sdk-golang"
  dbAzure "github.com/gpompe/databricks-sdk-golang/azure"

)

opt := databricks.NewDBClientOption("", "", os.Getenv("DATABRICKS_HOST"), os.Getenv("DATABRICKS_TOKEN"))
c := dbAzure.NewDBClient(opt)

jobs, err := c.Jobs().List()
```

## Implementation Progress

| API  | AWS | Azure |
| :--- | :---: | :---: |
| Account API | ✗ | N/A |
| Clusters API | ✗ | ✔ |
| Cluster Policies API | ✗ | ✗ |
| DBFS API | ✗ | ✔ |
| Global Init Scripts API | ✗ | ✗ |
| Groups API | ✗ | ✔ |
| Instance Pools API | ✗ | ✔ |
| Instance Profiles API | ✗ | N/A |
| IP Access List API | ✗ | ✗ |
| Jobs API | ✗ | ✔ |
| Libraries API | ✗ | ✔ |
| MLflow** API | ✗ | ✗ |
| Permissions API | ✗ | ✗ |
| SCIM** API | ✗ | ✗ |
| Secrets API | ✗ | ✔ |
| Token API | ✗ | ✔ |
| Token Management API | ✗ | ✗ |
| Workspace API | ✗ | ✔ |

** SCIM and MLflow are separate systems that are planned differently.
