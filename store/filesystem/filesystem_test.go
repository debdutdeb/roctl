package filesystem

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"roctl/models"
	"testing"
)

func setupCorrectStorageFile() {
	store := models.FilesystemStore{Deployments: []struct {
		ContainerName string `json:"container_name,omitempty"`
	}{
		{
			ContainerName: "rocketchat_abcd",
		},
		{
			ContainerName: "mongodb_hytg",
		},
	},
	}

	_bytes, err := json.Marshal(store)
	if err != nil {
		panic(errors.New("failed to marshal test store contents"))
	}

	err = ioutil.WriteFile("/tmp/.roctl.json", _bytes, 0700)
	if err != nil {
		panic(errors.New("failed to write to test storage file /tmp/.roctl.json"))
	}
}

func setupIncorrectStorageFile() {
	store := struct{ randomField string }{
		randomField: "roctl",
	}

	_bytes, err := json.Marshal(store)
	if err != nil {
		panic(errors.New("failed to marshal test store contents"))
	}

	err = ioutil.WriteFile("/tmp/.roctl.json", _bytes, 0700)
	if err != nil {
		panic(errors.New("failed to write to test storage file /tmp/.roctl.json"))
	}
}

func clearStorageFile() {
	os.Remove("/tmp/.roctl.json")
}

func TestNewFilesystemStore(t *testing.T) {
	defer clearStorageFile()

	var store *models.FilesystemStore
	var err error

	setupCorrectStorageFile()

	store, err = NewFilesystemStore("/tmp/.roctl.json")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		t.Fail()
	}

	expected := models.FilesystemStore{Deployments: []struct {
		ContainerName string `json:"container_name,omitempty"`
	}{
		{
			ContainerName: "rocketchat_abcd",
		},
		{
			ContainerName: "mongodb_hytg",
		},
	},
	}

	if !reflect.DeepEqual(expected, *store) {
		fmt.Fprintf(os.Stderr, "expected: %v, got: %v\n", expected, store)
		t.Fail()
	}
	clearStorageFile()

	setupIncorrectStorageFile()

	store, err = NewFilesystemStore("/tmp/.roctl.json")
	if store == new(models.FilesystemStore) {
		t.Fail()
	}

}
