package libvirt

import (
	"github.com/libvirt/libvirt-go"

	"github.com/cloudflavor/dweller/pkg/providers/config"
)

var (
	defaultSystemURI = "qemu:///system"
)

type LibvirtProvider struct {
	client *libvirt.Connect
}

// NewInfra creates a new cloudflavor infrastructure on top of qemu formed out
// of 1xMaster and 2xWorker Nodes with an attached storage.
func (lb *LibvirtProvider) NewInfra(config *config.Infra) error {
	// domain := newDomain()
	// libvirt
	return nil
}

func (lp *LibvirtProvider) RegisterInstances(config *config.Infra) error {
	return nil
}

func (lb *LibvirtProvider) DestroyInstances(config *config.Infra) error {
	return nil
}

func (lb *LibvirtProvider) HaltInfra(config *config.Infra) error {
	return nil
}

func NewLibvirtProvider(config *config.Infra) (*LibvirtProvider, error) {

	if config.LibvirtURI == nil {
		config.LibvirtURI = &defaultSystemURI
	}

	conn, err := libvirt.NewConnect(*config.LibvirtURI)
	if err != nil {
		return nil, err
	}

	return &LibvirtProvider{
		client: conn,
	}, nil
}
