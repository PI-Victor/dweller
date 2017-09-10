/*
Copyright 2017 The Cloudflavor authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cloudflavor/dweller/pkg/cli"
)

// rootCmd is the main command for the CLI.
var rootCmd = &cobra.Command{
	Use:   "dw",
	Short: "dw - A CLI for provisioning a new Cloudflavor infrastructure",
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
