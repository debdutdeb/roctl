package filesystem

type FilesystemStore struct {
	Deployments []struct {
		ContainerName string `json:"container_name,omitempty"`
	} `json:"deployments,omitempty"`
}
