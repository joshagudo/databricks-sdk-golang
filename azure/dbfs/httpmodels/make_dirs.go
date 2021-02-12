package httpmodels

type MakeDirsReq struct {
	Path string `json:"path,omitempty" url:"path,omitempty"`
}