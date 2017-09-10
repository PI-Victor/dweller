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

package cli

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cloudflavor/dweller/pkg/config"
	"github.com/cloudflavor/dweller/pkg/providers"
)

var (
	infraProviderURI string
	destroyInfra     bool
	pauseInfra       bool
	upWorkers        int
	infraProvider    string
	masterIP         string
	virtProvisioner  string
)

// UpCommand provisines a basic cluster infrastructure.
var UpCommand = &cobra.Command{
	Use:   "up",
	Short: "Bring up a simple Cloudflavor infrastructure ",
	Example: `dw up - Starts a basic Cloudflavor infrastructure with the default
provider, libvirt. It will provision 3 machines - two workers and a master node.
`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := &config.Infra{
			LibvirtURI:   &infraProviderURI,
			ProviderName: infraProvider,
			Workers:      upWorkers,
		}
		prov, err := provider.NewProvider(conf)
		if err != nil {
			logrus.Fatalf("Error occured while creating a new provider: %#v", err)
		}
		infra := provider.NewInfra(prov, conf)
		if err = infra.Up(); err != nil {
			logrus.Fatalf("Error while bringing up infrastructure: %#v", err)
		}
	},
}

// NewCommand adds a new instance to an existing infrastructure.
var NewCommand = &cobra.Command{
	Use:   "new",
	Short: "Add a new instance to an already running cloudflavor infrastructure",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// DeleteCommand deletes a specified instance from the infrastructure.
var DeleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete a specific instance",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// HaltCommand halts the current running infrastructure. It can use --pause to
// pause it or --destroy to destroy it permanently.
var HaltCommand = &cobra.Command{
	Use:   "halt",
	Short: "Halts the currently running infrastructure",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	UpCommand.PersistentFlags().IntVar(
		&upWorkers,
		"workers",
		2,
		"Specify number of worker instances to provision",
	)

	UpCommand.PersistentFlags().StringVar(
		&infraProvider,
		"provider",
		"libvirt",
		"Specify the provider",
	)

	UpCommand.PersistentFlags().StringVar(
		&virtProvisioner,
		"provisioner",
		"qemu",
		"Specify the provisioner for libvirt (qemu or xen)",
	)

	UpCommand.PersistentFlags().StringVar(
		&infraProviderURI,
		"uri",
		"qemu:///system",
		"The URI that the provider should use to connect to",
	)

	NewCommand.PersistentFlags().StringVar(
		&masterIP,
		"master",
		"",
		"Specify the IP address of the master node",
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
