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

// Client interfaces the libvirt client.
type Client interface {
	DomainDefineXML(xmlConfig string) (*libvirt.Domain, error)
}

// Provider contains the libvirt connection instance information.
type Provider struct {
	Client     Client
	Controller controllers.ProviderController
}

// NewInfra creates a new cloudflavor infrastructure on top of qemu containing
// 1xMaster and 2xWorker Nodes with attached.
func (p *Provider) NewInfra(config *config.Infra) error {
	if err := p.Controller.CreateResources(); err != nil {
		return err
	}
	return nil
}

// RegisterInstances adds a new instance to the infrastructure.
func (p *Provider) RegisterInstances(config *config.Infra) error {
	return nil
}

// DestroyInstances destroys a specific instance from the infrastructure.
func (p *Provider) DestroyInstances(config *config.Infra) error {
	return nil
}

// HaltInfra halts the current running infrastructure. It can also pause the
// current running infra and even destroy it if --destroy flag was passed.
func (p *Provider) HaltInfra(config *config.Infra) error {
	return nil
}

// ListInstances lists all the available instances for the provider.
func (p *Provider) ListInstances(config *config.Infra) error {
	return nil
}

// NewLibvirtProvider creates a new libvirt provider with an active connection.
func NewLibvirtProvider(config *config.Infra) (*Provider, error) {
	if config.LibvirtURI == nil {
		config.LibvirtURI = &defaultSystemURI
	}

	conn, err := libvirt.NewConnect(*config.LibvirtURI)
	if err != nil {
		return nil, err
	}
	// We provision the infrastructure with the number of workers that the user
	// specified. If there was no number of workers specified, we use the default
	// (2).
	var res []resources
	for i := 0; i < config.Workers; i++ {
		res = append(res, newDomainResource())
	}

	// NOTE: remember to refactor resource definition.
	return &Provider{
		Client:     conn,
		Controller: newController(conn, res),
	}, nil
}
