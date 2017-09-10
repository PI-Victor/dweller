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
	"errors"
	"testing"

	"github.com/libvirt/libvirt-go"

	"github.com/cloudflavor/dweller/pkg/config"
)

var (
	ErrTestLibvirt = errors.New("this is a test")
)

type mockLibvirtClient struct {
	domain *libvirt.Domain
	err    error
}

func (m *mockLibvirtClient) DomainDefineXML(xmlConfig string) (*libvirt.Domain, error) {
	return m.domain, m.err
}

func newMockLibvirtClient(domain *libvirt.Domain, err error) *mockLibvirtClient {
	return &mockLibvirtClient{
		domain: domain,
		err:    err,
	}
}

func newMockLibvirtProvider(domain *libvirt.Domain, err error) *Provider {
	return &Provider{
		Client: newMockLibvirtClient(domain, err),
	}
}

func TestLibvirtNewInfra(t *testing.T) {
	np := newMockLibvirtProvider(nil, nil)
	np.Controller = newFakeController(nil)
	err := np.NewInfra(&config.Infra{})
	if err != nil {
		t.Errorf("Did not expect an error, got: %#v", err)
	}
}

func TestLibvirtNewInfraError(t *testing.T) {
	np := newMockLibvirtProvider(nil, ErrTestLibvirt)
	np.Controller = newFakeController(ErrTestLibvirt)
	err := np.NewInfra(&config.Infra{})
	if err == nil {
		t.Errorf("Expected error to be: %#v got nil", ErrTestLibvirt)
	}
	if err != ErrTestLibvirt {
		t.Errorf("Expected error to match: %#v, got: %#v", ErrTestLibvirt, err)
	}
}

func TestRegisterInstances(t *testing.T) {
	np := newMockLibvirtProvider(nil, nil)
	err := np.RegisterInstances(&config.Infra{})
	if err != nil {
		t.Errorf("Did not expected an error, got: %#v", err)
	}
}

func TestDestroyInstances(t *testing.T) {
	np := newMockLibvirtProvider(nil, nil)
	err := np.DestroyInstances(&config.Infra{})
	if err != nil {
		t.Errorf("Did not expected an error, got: %#v", err)
	}
}

func TestHaltInfra(t *testing.T) {
	np := newMockLibvirtProvider(nil, nil)
	err := np.HaltInfra(&config.Infra{})
	if err != nil {
		t.Errorf("Did not expected an error, got: %#v", err)
	}
}

func TestListInstances(t *testing.T) {
	np := newMockLibvirtProvider(nil, nil)
	err := np.ListInstances(&config.Infra{})
	if err != nil {
		t.Errorf("Did not expected an error, got: %#v", err)
	}
}
