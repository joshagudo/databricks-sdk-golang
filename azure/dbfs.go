package azure

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/xinsnake/databricks-sdk-golang/azure/dbfs/httpmodels"
	"github.com/xinsnake/databricks-sdk-golang/azure/dbfs/models"
)

// DbfsAPI exposes the DBFS API
type DbfsAPI struct {
	Client DBClient
}

func (a DbfsAPI) init(client DBClient) DbfsAPI {
	a.Client = client
	return a
}

// AddBlock appends a block of data to the stream specified by the input handle
func (a DbfsAPI) AddBlock(data httpmodels.AddBlockReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/add-block", data, nil)
	return err
}

// Close closes the stream specified by the input handle
func (a DbfsAPI) Close(data httpmodels.CloseReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/close", data, nil)
	return err
}


// Create opens a stream to write to a file and returns a handle to this stream
func (a DbfsAPI) Create(data httpmodels.createReq) (httpmodels.CreateResp, error) {
	var createResponse httpmodels.CreateResp

	resp, err := a.Client.performQuery(http.MethodPost, "/dbfs/create", data, nil)

	if err != nil {
		return createResponse, err
	}

	err = json.Unmarshal(resp, &createResponse)
	return createResponse, err
}

// Delete deletes the file or directory (optionally recursively delete all files in the directory)
func (a DbfsAPI) Delete(data httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/delete", data, nil)
	return err
}

// GetStatus gets the file information of a file or directory
func (a DbfsAPI) GetStatus(data httpmodels.GetStatusReq) (httpmodels.GetStatusResp, error) {
	var fileInfo httpmodels.GetStatusResp

	resp, err := a.Client.performQuery(http.MethodGet, "/dbfs/get-status", data, nil)

	if err != nil {
		return fileInfo, err
	}

	err = json.Unmarshal(resp, &fileInfo)
	return fileInfo, err
}


// List lists the contents of a directory, or details of the file
func (a DbfsAPI) List(data httpmodels.ListReq) (httpmodels.ListResp, error) {
	var listResponse httpmodels.ListResp

	resp, err := a.Client.performQuery(http.MethodGet, "/dbfs/list", data, nil)

	if err != nil {
		return listResponse, err
	}

	err = json.Unmarshal(resp, &listResponse)
	return listResponse, err
}

// Mkdirs creates the given directory and necessary parent directories if they do not exist
func (a DbfsAPI) Mkdirs(data httpmodels.MakeDirsReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/mkdirs", data, nil)
	return err
}

// Move moves a file from one location to another location within DBFS
func (a DbfsAPI) Move(data httpmodels.MoveReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/move", data, nil)
	return err
}

// Put uploads a file through the use of multipart form post
func (a DbfsAPI) Put(data httpmodels.PutReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/put", data, nil)
	return err
}

// Read returns the contents of a file
func (a DbfsAPI) Read(data httpmodels.ReadReq) (httpmodels.ReadResp, error) {
	var readResponseBase64 httpmodels.ReadRespBase64
	var readResponse DbfsReadResponse

	resp, err := a.Client.performQuery(http.MethodGet, "/dbfs/read", data, nil)

	if err != nil {
		return readResponse, err
	}

	err = json.Unmarshal(resp, &readResponseBase64)
	if err != nil {
		return readResponse, err
	}

	readResponse.BytesRead = readResponseBase64.BytesRead
	readResponse.Data, err = base64.StdEncoding.DecodeString(readResponseBase64.Data)
	return readResponse, err
}
