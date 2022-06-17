package docker

import (
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func NewDockerDeploymentHandler() (*DockerDeploymentModel, error) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	network, err := client.NetworkCreate(context.Background(), "roctl_rocketchat_network", types.NetworkCreate{})
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &DockerDeploymentModel{
		client: client,
		network: DockerNetwork{
			id:   network.ID,
			name: "roctl_rocketchat_network",
		},
	}, nil
}
