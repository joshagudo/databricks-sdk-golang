package httpmodels

type ReadResp struct {
	BytesRead int64  `json:"bytes_read,omitempty" url:"bytes_read,omitempty"`
	Data      []byte `json:"data,omitempty" url:"data,omitempty"`
}

type ReadRespBase64 struct {
	BytesRead int64  `json:"bytes_read,omitempty" url:"bytes_read,omitempty"`
	Data      []byte `json:"data,omitempty" url:"data,omitempty"`
}