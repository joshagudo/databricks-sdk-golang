package models

type InitScriptInfo struct {
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty" url:"dbfs,omitempty"`
	fs   *FileStorageInfo `json:"fs,omitempty" url:"fs,omitempty"`
}
