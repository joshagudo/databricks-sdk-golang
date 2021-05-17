# databricks-sdk-golang

This is a Golang SDK for [DataBricks REST API 2.0](https://docs.databricks.com/api/latest/index.html#) and [Azure DataBricks REST API 2.0](https://docs.azuredatabricks.net/api/latest/index.html).

**WARNING: The SDK is unstable and under development. More testing needed!**

## Usage

```go
import (
  databricks "github.com/polar-rams/databricks-sdk-golang"
  dbAzure "github.com/polar-rams/databricks-sdk-golang/azure"
  // dbAws "github.com/polar-rams/databricks-sdk-golang/aws"
)

var o databricks.DBClientOption
o.Host = os.Getenv("DATABRICKS_HOST")
o.Token = os.Getenv("DATABRICKS_TOKEN")

var c dbAzure.DBClient
c.Init(o)

jobs, err := c.Jobs().List()
```

## Implementation Progress

| API  | AWS | Azure |
| :--- | :---: | :---: |
| Account API | ✗ | N/A |
| Clusters API | ✔ (Outdated) | ✔ |
| Cluster Policies API | ✗ | ✗ |
| DBFS API | ✔ (Outdated) | ✔ (Outdated) |
| Global Init Scripts API | ✗ | ✗ |
| Groups API | ✔ (Outdated) | ✔ |
| Instance Pools API | ✗ | ✗ |
| Instance Profiles API | ✔ (Outdated) | N/A |
| IP Access List API | ✗ | ✗ |
| Jobs API | ✔ (Outdated) | ✔ |
| Libraries API | ✔ (Outdated) | ✔ |
| MLflow** API | ✗ | ✗ |
| Permissions API | ✗ | ✗ |
| SCIM** API | ✗ | ✗ |
| Secrets API | ✔ (Outdated) | ✔ (Outdated) |
| Token API | ✔ (Outdated) | ✔ (Outdated) |
| Token Management API | ✗ | ✗ |
| Workspace API | ✔ (Outdated) | ✔ (Outdated) |

** SCIM and MLflow are separate systems that are planned differently.
