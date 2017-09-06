package libvirt

import (
	"github.com/libvirt/libvirt-go-xml"
)

func newStoragePool() *libvirtxml.StoragePool {
	return &libvirtxml.StoragePool{
		Source: &libvirtxml.StoragePoolSource{},
	}
}

func newNetwork() *libvirtxml.NetworkDomain {
	return &libvirtxml.NetworkDomain{}
}
