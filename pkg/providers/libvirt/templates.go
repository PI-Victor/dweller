package libvirt

import (
	"github.com/libvirt/libvirt-go-xml"
)

var (
	defaultDomainName          = "ArmetOS-Infra"
	defaultMemory              = "1024"
	defaultAttachedStorageSize = "40GB"
)

func newDomain() *libvirtxml.Domain {
	newOS := newDomainOS()
	newVCPU := newDomainVCPU()
	newCPU := newDomainCPU()

	return &libvirtxml.Domain{
		Name: defaultDomainName,
		CPU:  newCPU,
		VCPU: newVCPU,
		OS:   newOS,
	}
}

func newDomainDisk() *libvirtxml.DomainDisk {
	return &libvirtxml.DomainDisk{}
}

func newDomainOS() *libvirtxml.DomainOS {
	return &libvirtxml.DomainOS{
		// TODO: fix this!
		Type: &libvirtxml.DomainOSType{},
	}
}

func newDomainVCPU() *libvirtxml.DomainVCPU {
	return &libvirtxml.DomainVCPU{
		Value:     2,
		Placement: "Static",
	}
}

func newDomainCPU() *libvirtxml.DomainCPU {
	return &libvirtxml.DomainCPU{}
}
