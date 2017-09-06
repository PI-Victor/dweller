package libvirt

import (
	"testing"
)

func TestNewDomainResource(t *testing.T) {
	domain := newDomainResource()
	if domain.Name != defaultDomainName {
		t.Errorf("Expected domain name to be: %#v, got :%#v",
			defaultDomainName,
			domain.Name,
		)
	}
}

func TestNewBIOSBootloaderResource(t *testing.T) {
	boot := newBIOSBootloaderResource()
	if boot.Readonly != "yes" {
		t.Errorf("Expected Readonly value to be yes, got: %s", boot.Readonly)
	}
}

func TestNewDomainDiskMasterResource(t *testing.T) {
	disks := newDomainDiskResource(true)
	if len(disks) > 1 {
		t.Errorf("Expected a single disk resource, got: %d", len(disks))
	}
}

func TestNewDomainVCPUResource(t *testing.T) {
	vcpu := newDomainVCPUResource()
	if vcpu.Value != defaultDomainVCPU {
		t.Errorf("Expected VCPU to be: %d got: %d", defaultDomainVCPU, vcpu.Value)
	}
}

func TestNewDomainMemoryResource(t *testing.T) {
	mem := newDomainMemoryResource()
	if mem.domain.Value != defaultCurrentMemory {
		t.Errorf("Expected memory: %d, got: %d", defaultCurrentMemory, mem.domain.Value)
	}
}
