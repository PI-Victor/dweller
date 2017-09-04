package libvirt

import (
	"github.com/libvirt/libvirt-go-xml"
)

var (
	defaultDomainName      = "ArmetOS-Infra"
	defaultDomainType      = "qemu"
	defaultDomainOSType    = "hvm"
	defaultDomainOSMachine = "pc"
	defaultOSSourceFile    = "/var/lib/libvirt/images/armetOS-0-1.x86_64.qcow2"
	// NOTE: till ArmetOS qcow2 is ready, use a default image from fedora cloud
	tmpOSSOurceFile                 = "/var/lib/libvirt/images/fedora25.qcow2"
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
		Type:       defaultDomainType,
		Name:       defaultDomainName,
		VCPU:       newDomainVCPU(),
		OS:         newDomainOS(),
		Memory:     memSpec.domain,
		Devices:    newDomainDevices(),
		OnReboot:   "restart",
		OnPoweroff: "destroy",
		OnCrash:    "destroy",
	}
}

func newDomainDisk(master bool) []libvirtxml.DomainDisk {
	// NOTE: Disk creation should be handled by the controller and attached
	// according to node. Master or Worker.
	return []libvirtxml.DomainDisk{
		{
			Driver: &libvirtxml.DomainDiskDriver{
				Name: "qemu",
				Type: "qcow2",
			},
			Type:   "file",
			Device: "disk",
			Source: &libvirtxml.DomainDiskSource{
				// NOTE: remember to change this to a working armetOS qcow2 file.
				File: tmpOSSOurceFile,
			},
			Target: &libvirtxml.DomainDiskTarget{
				Dev: "vda",
				Bus: "virtio",
			},
		},
	}
}

func newDomainDevices() *libvirtxml.DomainDeviceList {
	return &libvirtxml.DomainDeviceList{
		Disks: newDomainDisk(false),
		Graphics: []libvirtxml.DomainGraphic{
			{
				Type:     "spice",
				AutoPort: "yes",
			},
		},
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

func newDomainMemory() *domainMemory {
	return &domainMemory{
		domain: &libvirtxml.DomainMemory{
			Value: defaultCurrentMemory,
			Unit:  "MiB",
		},
	}
}
