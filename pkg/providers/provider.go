package provider

import (
	"fmt"

	"github.com/cloudflavor/dweller/pkg/providers/config"
	"github.com/cloudflavor/dweller/pkg/providers/libvirt"
)

// Provider is an interface that all service providers must implement in order
// order to provision new Cloudflavor infrastructure.
type Provider interface {
	NewInfra(*config.Infra) error
	HaltInfra(*config.Infra) error
	RegisterInstances(*config.Infra) error
	DestroyInstances(*config.ConfigInfra) error
}

// CloudInfra contains information about the provisioners that that infra will
// be instantiated .
type CloudInfra struct {
	Provider string
	Config   *config.ConfigInfra
}

func newProvider(provider string) (Provider, error) {
	// NOTE: suffices for now, should account for future providers.
	return libvirt.NewLibvirtProvider()
}

// Up will bring up a new infrastructure.
func (cf *CloudInfra) Up() error {
	provider, err := newProvider(cf.Provider)
	if err != nil {
		return err
	}
	if err := provider.NewInfra(); err != nil {
		return err
	}
	return nil
}

// Halt will halt the infrastructure with the options of pausing it or
// destroying it permanently.
func (cf *CloudInfra) Halt(delete, pause bool) error {
	return nil
}

func (cf *CloudInfra) NewInstances() error {
	return nil
}

func (cf *CloudInfra) DeleteInstances(instance string) error {
	return nil
}
