package httpmodels

type CreateReq struct {
	InstancePoolName                   string `json:"instance_pool_name,omitempty" url:"instance_pool_name,omitempty"`
	MinIdleInstances                   string `json:"min_idle_instances,omitempty" url:"min_idle_instances,omitempty"`
	MaxCapacity                        string `json:"max_capacity,omitempty" url:"max_capacity,omitempty"`
	NodeTypeId                         string `json:"node_type_id,omitempty" url:"node_type_id,omitempty"`
	CustomTags                         string `json:"custom_tags,omitempty" url:"custom_tags,omitempty"`
	IdleInstanceAutoterminationMinutes string `json:"idle_instance_autotermination_minutes,omitempty" url:"idle_instance_autotermination_minutes,omitempty"`
	EnableElasticDisk                  string `json:"enable_elastic_disk,omitempty" url:"enable_elastic_disk,omitempty"`
	DiskSpec                           string `json:"disk_spec,omitempty" url:"disk_spec,omitempty"`
	PreloadedSparkVersions             string `json:"preloaded_spark_versions,omitempty" url:"preloaded_spark_versions,omitempty"`
	PreloadedDockerImages              string `json:"preloaded_docker_images,omitempty" url:"preloaded_docker_images,omitempty"`
	AzureAttributes                    string `json:"azure_attributes,omitempty" url:"azure_attributes,omitempty"`
}

type CreateResp struct {
	InstancePoolID string `json:"instance_pool_id,omitempty" url:"instance_pool_id,omitempty"`
}
