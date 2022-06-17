package models

type Store interface {
	AddDeployment(name string) error
	DeleteDeployment(name string) error
}

type DeploymentsStore struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}
