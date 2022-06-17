package cmd

import "github.com/spf13/cobra"

var rootCommand = cobra.Command{
	Use:   "roctl",
	Short: "Manage Rocket.Chat",
}

func Execute() {
	cobra.CheckErr(rootCommand.Execute())
}
