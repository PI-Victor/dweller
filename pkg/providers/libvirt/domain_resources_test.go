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
