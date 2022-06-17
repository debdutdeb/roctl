package filesystem

import "roctl/models"

type FilesystemStore struct {
	Deployments []models.DeploymentsStore `json:"deployments,omitempty"`
}
