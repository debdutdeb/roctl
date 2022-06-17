package docker

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

func (d *DockerDeploymentModel) CreateDeployment(name, version string) (string, error) {
	fmt.Println("here")
	body, err := d.client.ContainerCreate(context.Background(), &container.Config{
		AttachStderr:    false,
		AttachStdin:     false,
		AttachStdout:    false,
		Image:           fmt.Sprintf("registry.rocket.chat/rocketchat/rocket.chat:%s", version),
		NetworkDisabled: false,
	}, nil, &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{d.network.name: &network.EndpointSettings{NetworkID: d.network.id}},
	}, nil, name)
	if err != nil {
		return "", errors.New(err.Error())
	}
	fmt.Printf("%v\n", body)
	return "", nil
}
