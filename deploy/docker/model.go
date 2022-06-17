package docker

import "github.com/docker/docker/client"

type DockerNetwork struct {
	id   string
	name string
}

type DockerDeploymentModel struct {
	client  *client.Client
	network DockerNetwork
}
