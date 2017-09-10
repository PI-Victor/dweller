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

func newDomainResource() *libvirtxml.Domain {
	memSpec := newDomainMemoryResource()
	return &libvirtxml.Domain{
		Type:       defaultDomainType,
		Name:       defaultDomainName,
		VCPU:       newDomainVCPUResource(),
		OS:         newDomainOSResource(),
		Memory:     memSpec.domain,
		Devices:    newDomainDevicesResource(false),
		OnReboot:   "restart",
		OnPoweroff: "destroy",
		OnCrash:    "destroy",
	}
}

func newBIOSBootloaderResource() *libvirtxml.DomainLoader {
	return &libvirtxml.DomainLoader{
		Readonly: "yes",
		Type:     "pflash",
		Secure:   "no",
		// NOTE: remember to check this out!???
		Path: "/usr/lib/libvirt",
	}
}

func newDomainDiskResource(master bool) []libvirtxml.DomainDisk {
	// NOTE: Disk creation should be handled by the controller and attached
	// according to node. Master or Worker.
	if !master {
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
			{
				Driver: &libvirtxml.DomainDiskDriver{
					Name: "qemu",
					Type: "raw",
				},
				Type:   "volume",
				Device: "disk",
				Source: &libvirtxml.DomainDiskSource{
					// TODO: only for testing - change this to a created pool
					Pool:   "test-pool",
					Volume: "test-volume",
				},
			},
		}
	}

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

func newDomainDevicesResource(isMaster bool) *libvirtxml.DomainDeviceList {
	return &libvirtxml.DomainDeviceList{
		Disks: newDomainDiskResource(isMaster),
		Graphics: []libvirtxml.DomainGraphic{
			{
				Type:     "spice",
				AutoPort: "yes",
			},
		},
	}
}

func newDomainOSResource() *libvirtxml.DomainOS {
	return &libvirtxml.DomainOS{
		Loader: newBIOSBootloaderResource(),
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

func newDomainVCPUResource() *libvirtxml.DomainVCPU {
	return &libvirtxml.DomainVCPU{
		Value:     defaultDomainVCPU,
		Placement: defaultPlacement,
	}
}

func newDomainMemoryResource() *domainMemory {
	return &domainMemory{
		domain: &libvirtxml.DomainMemory{
			Value: defaultCurrentMemory,
			Unit:  "MiB",
		},
	}
}
