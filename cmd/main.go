package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cloudflavor/dweller/pkg/cli"
)

// rootCmd is the main command for the CLI.
var rootCmd = &cobra.Command{
	Use:   "dw",
	Short: "dw - A CLI to provision new Coudflavor infrastructure",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(cli.UpCommand)
	rootCmd.AddCommand(cli.NewCommand)
	rootCmd.AddCommand(cli.HaltCommand)
	rootCmd.AddCommand(cli.DeleteCommand)

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("Failed to load command: %s", err)
	}
}

func init() {
	// TODO: make this configurable through env var.
	logrus.SetLevel(logrus.DebugLevel)
}
