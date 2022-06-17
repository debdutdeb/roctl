package models

type DeploymentHandler interface {
	CreateDeployment(_store Store) (Deployment, error)
	StartDeployment(_store Store) (string, error)
	UpgradeDeployment(version string, _store Store) error
	BackupDeployment() error
	ExportContext(name string) error
}

type Deployment struct {
	Id      string
	Handler DeploymentHandler
}
