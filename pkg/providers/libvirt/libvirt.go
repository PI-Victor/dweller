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

// QemuClient interfaces the libvirt client.
type QemuClient interface {
	DomainDefineXML(xmlConfig string) (*libvirt.Domain, error)
	StoragePoolDefineXML(xmlConfig string, flags uint32) (*libvirt.StoragePool, error)
	NetworkDefineXML(xmlConfig string) (*libvirt.Network, error)
	ListAllNetworks(flags libvirt.ConnectListAllNetworksFlags) ([]libvirt.Network, error)
	ListAllStoragePools(flags libvirt.ConnectListAllStoragePoolsFlags) ([]libvirt.StoragePool, error)
}

// QemuProvider contains the libvirt connection instance information.
type QemuProvider struct {
	Client     QemuClient
	Controller controllers.ProviderController
}

// NewInfra creates a new cloudflavor infrastructure on top of qemu containing
// 1xMaster and 2xWorker Nodes with attached.
func (p *QemuProvider) NewInfra() error {
	if err := p.Controller.CreateResources(); err != nil {
		return err
	}
	return nil
}

// NewInstance adds a new instance to the infrastructure.
func (p *QemuProvider) NewInstance() error {
	return nil
}

// DestroyInstances destroys a specific instance from the infrastructure.
func (p *QemuProvider) DestroyInstances() error {
	if err := p.Controller.DeleteResources(); err != nil {
		return err
	}
	return nil
}

// HaltInfra halts the current running infrastructure. It can also pause the
// current running infra and even destroy it if --destroy flag was passed.
func (p *QemuProvider) HaltInfra() error {
	return nil
}

// ListInstances lists all the available instances for the provider.
func (p *QemuProvider) ListInstances() error {
	return nil
}

// RegisterInstances registers new instances into the cluster
func (p *QemuProvider) RegisterInstances() error {
	return nil
}

// NewLibvirtProvider creates a new libvirt provider with an active connection.
func NewLibvirtProvider(config *config.Infra) (*QemuProvider, error) {
	if config.LibvirtURI == nil {
		config.LibvirtURI = &defaultSystemURI
	}

	conn, err := libvirt.NewConnect(*config.LibvirtURI)
	if err != nil {
		return nil, err
	}
	// We also add the generic resources for the machine. Storage and Network.
	res := []resource{
		newStoragePool(),
		newNetwork(),
	}
	// We provision the infrastructure with the number of workers that the user
	// specified. If there was no number of workers specified, we use the default
	// (2).
	for i := 0; i < config.Workers; i++ {
		res = append(res, newDomainResource())
	}

	// NOTE: remember to refactor resource definition.
	return &QemuProvider{
		Client:     conn,
		Controller: newController(conn, res...),
	}, nil
}
