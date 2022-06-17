package filesystem

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"roctl/models"
)

// store in ~/.roctl.json ?

func NewFilesystemStore(storageFile string) (*FilesystemStore, error) {
	store := FilesystemStore{}

	if storageFile == "" {
		storageFile = "~/.roctl.json"
	}

	file, err := ioutil.ReadFile(storageFile)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = json.Unmarshal(file, &store)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &store, nil
}

func (s *FilesystemStore) AddDeployment(name, _type string) error {
	// only reaches this if deployment is successful
	// no need to check for name collisions
	s.Deployments = append(s.Deployments, models.DeploymentsStore{Name: name, Type: _type})
	return s.refresh()
}

func (s FilesystemStore) refresh() error {
	_bytes, err := json.Marshal(s)
	if err != nil {
		return errors.New(err.Error())
	}

	err = ioutil.WriteFile("~/.roctl.json", _bytes, 0700)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (s *FilesystemStore) DeleteDeployment(name string) error {
	var newDeploymentList []models.DeploymentsStore
	for idx, deployment := range s.Deployments {
		if deployment.Name == name {
			newDeploymentList = s.Deployments[:idx-1]
			newDeploymentList = append(newDeploymentList, s.Deployments[idx+1:]...)
			break
		}
	}
	s.Deployments = newDeploymentList
	return s.refresh()
}
