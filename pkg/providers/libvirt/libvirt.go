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

package libvirt

import (
	"github.com/libvirt/libvirt-go"

	"github.com/cloudflavor/dweller/pkg/config"
	"github.com/cloudflavor/dweller/pkg/providers/controllers"
)

var (
	defaultSystemURI = "qemu:///system"
)

// LibvirtClient interfaces the libvirt client.
type LibvirtClient interface {
	DomainDefineXML(xmlConfig string) (*libvirt.Domain, error)
}

// LibvirtProvider contains the libvirt connection instance information.
type LibvirtProvider struct {
	Client     LibvirtClient
	Controller controllers.ProviderController
}

// NewInfra creates a new cloudflavor infrastructure on top of qemu containing
// 1xMaster and 2xWorker Nodes with attached.
func (lb *LibvirtProvider) NewInfra(config *config.Infra) error {
	if err := lb.Controller.CreateResources(); err != nil {
		return err
	}
	return nil
}

// RegisterInstances adds a new instance to the infrastructure.
func (lb *LibvirtProvider) RegisterInstances(config *config.Infra) error {
	return nil
}

// DestroyInstances destroys a specific instance from the infrastructure.
func (lb *LibvirtProvider) DestroyInstances(config *config.Infra) error {
	return nil
}

// HaltInfra halts the current running infrastructure. It can also pause the
// current running infra and even destroy it if --destroy flag was passed.
func (lb *LibvirtProvider) HaltInfra(config *config.Infra) error {
	return nil
}

// NewLibvirtProvider creates a new libvirt provider with an active connection.
func NewLibvirtProvider(config *config.Infra) (*LibvirtProvider, error) {
	if config.LibvirtURI == nil {
		config.LibvirtURI = &defaultSystemURI
	}

	conn, err := libvirt.NewConnect(*config.LibvirtURI)
	if err != nil {
		return nil, err
	}
	// NOTE: remember to refactor resource definition.
	return &LibvirtProvider{
		Client:     conn,
		Controller: newController(conn, newDomainResource()),
	}, nil
}
