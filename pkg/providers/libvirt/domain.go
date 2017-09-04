package libvirt

import (
	"github.com/libvirt/libvirt-go-xml"
)

var (
	defaultDomainName               = "ArmetOS-Infra"
	defaultDomainType               = "qemu"
	defaultDomainOSType             = "hvm"
	defaultDomainOSMachine          = "pc"
	defaultOSSourceFile             = "/var/lib/libvirt/images/armetOS-0-1.x86_64.qcow2"
	defaultCurrentMemory       uint = 1024
	defaultMaxMemory           uint = 4096
	defaultAttachedStorageSize      = "40GB"
	defaultPlacement                = "static"
	defaultDomainArch               = "x86_64"
	defaultEnabledBootMenu          = "no"
	defaultTimeoutBootMenu          = "1"
	defaultDomainVCPU               = 2
)

type domainMemory struct {
	domain  *libvirtxml.DomainMemory
	current *libvirtxml.DomainMemory
}

func newDomain() *libvirtxml.Domain {
	memSpec := newDomainMemory()
	return &libvirtxml.Domain{
		Type:   defaultDomainType,
		Name:   defaultDomainName,
		CPU:    newDomainCPU(),
		VCPU:   newDomainVCPU(),
		OS:     newDomainOS(),
		Memory: memSpec.domain,
	}
}

func newDomainDisk() []libvirtxml.DomainDisk {
	return []libvirtxml.DomainDisk{
		{
			Driver: &libvirtxml.DomainDiskDriver{
				Name: "qemu",
				Type: "qcow2",
			},
			Type:   "file",
			Device: "disk",
		},
	}
}

func newDomainDevices() *libvirtxml.DomainDeviceList {
	return &libvirtxml.DomainDeviceList{
		Disks: newDomainDisk(),
	}
}

func newDomainOS() *libvirtxml.DomainOS {
	return &libvirtxml.DomainOS{
		Type: &libvirtxml.DomainOSType{
			Type:    defaultDomainOSType,
			Arch:    defaultDomainArch,
			Machine: defaultDomainOSMachine,
		},
		BootMenu: &libvirtxml.DomainBootMenu{
			Enabled: defaultEnabledBootMenu,
			Timeout: defaultTimeoutBootMenu,
		},
	}
}

func newDomainVCPU() *libvirtxml.DomainVCPU {
	return &libvirtxml.DomainVCPU{
		Value:     defaultDomainVCPU,
		Placement: defaultPlacement,
	}
}

func newDomainCPU() *libvirtxml.DomainCPU {
	return &libvirtxml.DomainCPU{}
}

func newDomainMemory() *domainMemory {
	return &domainMemory{
		domain: &libvirtxml.DomainMemory{
			Value: defaultCurrentMemory,
		},
		current: &libvirtxml.DomainMemory{
			Value: defaultCurrentMemory,
		},
	}
}
