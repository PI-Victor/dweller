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

package provider

import (
	"errors"

	"github.com/cloudflavor/dweller/pkg/config"
	"github.com/cloudflavor/dweller/pkg/providers/libvirt"
)

var (
	defaultProvider = "libvirt"
	// ErrUnknownProvider is used for unsuported providers.
	ErrUnknownProvider = errors.New("Unkown provider")
)

// Provider is an interface that all providers must implement in order to
// provision new Cloudflavor infrastructure.
type Provider interface {
	NewInfra(*config.Infra) error
	HaltInfra(*config.Infra) error
	RegisterInstances(*config.Infra) error
	DestroyInstances(*config.Infra) error
	ListInstances(*config.Infra) error
}

// CloudInfra contains information about the provisioners that infra will be
// instantiated and their related configuration.
type CloudInfra struct {
	Config   *config.Infra
	Provider Provider
}

// NewInfra creates a new infrastructure instance that has information about
// the currently used provider.
func NewInfra(provider Provider, config *config.Infra) *CloudInfra {
	return &CloudInfra{
		Provider: provider,
		Config:   config,
	}
}

// NewProvider returns a new instance of a specified provider
func NewProvider(config *config.Infra) (Provider, error) {
	// NOTE: suffices for now, should account for future providers.
	if config.ProviderName == defaultProvider {
		return libvirt.NewLibvirtProvider(config)
	}
	return nil, ErrUnknownProvider
}

// Up will bring up a new infrastructure.
func (cf *CloudInfra) Up() error {
	if err := cf.Provider.NewInfra(cf.Config); err != nil {
		return err
	}
	return nil
}

// Halt will halt the infrastructure with the options of pausing it or
// destroying it permanently.
func (cf *CloudInfra) Halt(delete, pause bool) error {
	return nil
}

// NewInstances will add new instances to the already running infrastructure.
func (cf *CloudInfra) NewInstances() error {
	return nil
}

// DeleteInstances deletes one or more running instances.
func (cf *CloudInfra) DeleteInstances(instance string) error {
	return nil
}
