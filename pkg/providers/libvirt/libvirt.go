package libvirt

import (
	"github.com/Sirupsen/logrus"
	"github.com/libvirt/libvirt-go"

	"github.com/cloudflavor/dweller/pkg/config"
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
	LibvirtClient
}

// NewInfra creates a new cloudflavor infrastructure on top of qemu containing
// 1xMaster and 2xWorker Nodes with attached.
func (lb *LibvirtProvider) NewInfra(config *config.Infra) error {
	domain, err := newDomainResource().Marshal()
	if err != nil {
		return nil
	}
	logrus.Debugf("Domain XML Schema: %s", domain)

	_, err = lb.DomainDefineXML(domain)
	if err != nil {
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

	return &LibvirtProvider{
		conn,
	}, nil
}
