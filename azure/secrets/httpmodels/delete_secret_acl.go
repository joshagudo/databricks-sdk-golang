package httpmodels

type DeleteSecretAclReq struct {
	Scope     string `json:"scope,omitempty" url:"scope,omitempty"`
	Principal string `json:"principal,omitempty" url:"principal,omitempty"`
}
