package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"roctl/deploy/docker"
)

var version string

var deployCommand = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an instance",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := docker.NewDockerDeploymentHandler()
		if err != nil {
			panic(errors.New(err.Error()))
		}
		_, err = c.CreateDeployment("rocketchat", version)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	deployCommand.Flags().StringVarP(&version, "version", "", "latest", "Rocket.Chat version")
	rootCommand.AddCommand(deployCommand)
}
