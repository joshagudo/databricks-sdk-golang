package httpmodels

type RunsDeleteReq struct {
	RunID int64 `json:"run_id,omitempty" url:"run_id,omitempty"`
}
