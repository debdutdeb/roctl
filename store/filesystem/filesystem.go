package filesystem

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

func (s *FilesystemStore) AddDeployment(containerName string) error {
	// only reaches this if deployment is successful
	// no need to check for name collisions
	s.Deployments = append(s.Deployments, struct {
		ContainerName string `json:"container_name,omitempty"`
	}{ContainerName: containerName})
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
