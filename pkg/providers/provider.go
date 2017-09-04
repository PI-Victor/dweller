package provider

import (
	"github.com/cloudflavor/dweller/pkg/providers/config"
	"github.com/cloudflavor/dweller/pkg/providers/libvirt"
)

// Provider is an interface that all providers must implement in order to
// provision new Cloudflavor infrastructure.
type Provider interface {
	NewInfra(*config.Infra) error
	HaltInfra(*config.Infra) error
	RegisterInstances(*config.Infra) error
	DestroyInstances(*config.Infra) error
}

// CloudInfra contains information about the provisioners that infra will be
// instantiated.
type CloudInfra struct {
	Provider string
	Config   *config.Infra
}

// NewInfra creates a new infrastructure instance that has information about
// the currently used provider.
func NewInfra(provider string) *CloudInfra {
	return &CloudInfra{
		Provider: provider,
	}
}

func newProvider(config *config.Infra, provider string) (Provider, error) {
	// NOTE: suffices for now, should account for future providers.
	return libvirt.NewLibvirtProvider(config)
}

// Up will bring up a new infrastructure.
func (cf *CloudInfra) Up(config *config.Infra) error {
	prov, err := newProvider(config, cf.Provider)
	if err != nil {
		return err
	}
	if err := prov.NewInfra(config); err != nil {
		return err
	}
	return nil
}

// Halt will halt the infrastructure with the options of pausing it or
// destroying it permanently.
func (cf *CloudInfra) Halt(config *config.Infra, delete, pause bool) error {
	return nil
}

// NewInstances will add new instances to the already running infrastructure.
func (cf *CloudInfra) NewInstances(config *config.Infra) error {
	return nil
}

// DeleteInstances deletes one or more running instances.
func (cf *CloudInfra) DeleteInstances(config *config.Infra, instance string) error {
	return nil
}
