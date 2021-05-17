package httpmodels

type DeleteReq struct {
	Scope string `json:"scope,omitempty" url:"scope,omitempty"`
	Key   string `json:"key,omitempty" url:"key,omitempty"`
}
