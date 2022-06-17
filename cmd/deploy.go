package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var deployCommand = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an instance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deploying....")
	},
}

func init() {
	rootCommand.AddCommand(deployCommand)
}
