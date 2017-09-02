package cli

import (
	"github.com/spf13/cobra"
)

var (
	destroyInfra  bool
	pauseInfra    bool
	upInstances   int
	infraProvider string
)

// UpCommand provisines a basic cluster infrastructure.
var UpCommand = &cobra.Command{
	Use:   "up",
	Short: "Spin up a simple Cloudflavor infrastructure ",
	Example: `dw up - Starts a basic Cloudflavor infrastructure with the default provider, libvirt.
It will provision 3 machines - two workers and a master node.
`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// NewCommand adds a new instance to an existing infrastructure.
var NewCommand = &cobra.Command{
	Use:   "new",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// DeleteCommand deletes a specified instance from the infrastructure.
var DeleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var HaltCommand = &cobra.Command{
	Use:   "halt",
	Short: "Halt the infrastructure",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	UpCommand.PersistentFlags().IntVar(
		&upInstances,
		"instances",
		3,
		"Override the default number (1xMaster & 2xWorkers) of instances to provision",
	)

	UpCommand.PersistentFlags().
		StringVar(
			&infraProvider,
			"provider",
			"libvirt",
			"Override the default (libvirt) provider",
		)

	HaltCommand.PersistentFlags().BoolVar(
		&destroyInfra,
		"destroy",
		false,
		"Destroy the infrastructure",
	)

	HaltCommand.PersistentFlags().BoolVar(
		&pauseInfra,
		"pause",
		false,
		`Pause the infrastructure instead of halting it. Works only with the libvirt
provider`,
	)
}
