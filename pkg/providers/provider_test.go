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

package provider

import (
	"errors"
	"reflect"
	"testing"

	"github.com/cloudflavor/dweller/pkg/config"
)

var (
	ErrMockProvider = errors.New("this is a test")
)

type mockProvider struct {
	ProviderName string
	err          error
}

func (m *mockProvider) NewInfra(c *config.Infra) error          { return m.err }
func (m *mockProvider) HaltInfra(c *config.Infra) error         { return m.err }
func (m *mockProvider) RegisterInstances(c *config.Infra) error { return m.err }
func (m *mockProvider) DestroyInstances(c *config.Infra) error  { return m.err }
func (m *mockProvider) ListInstances(c *config.Infra) error     { return m.err }

func newMockConfig(prov string) *config.Infra {
	lvURI := "qemu:///sytem"
	return &config.Infra{
		LibvirtURI:   &lvURI,
		ProviderName: prov,
	}
}

func newMockProvider(conf *config.Infra, err error) Provider {
	return &mockProvider{
		ProviderName: conf.ProviderName,
		err:          err,
	}
}

func TestNewInvalidProvider(t *testing.T) {
	_, err := NewProvider(newMockConfig("invalidProvider"))
	if err != ErrUnknownProvider {
		t.Errorf("Expected error to be: %#v, got: %#v", ErrUnknownProvider, err)
	}
}

func TestNewInfra(t *testing.T) {
	conf := newMockConfig("libvirt")
	newProv := newMockProvider(conf, nil)
	ni := NewInfra(newProv, conf)

	if !reflect.DeepEqual(ni.Config, conf) {
		t.Errorf("Expected %#v, got: %#v", conf, ni.Config)
	}
	if !reflect.DeepEqual(ni.Provider, newProv) {
		t.Errorf("Expected %#v, got: %#v", newProv, ni.Provider)
	}
}

func TestInfraUp(t *testing.T) {
	conf := newMockConfig("libvirt")
	newProv := newMockProvider(conf, nil)
	newInfra := NewInfra(newProv, conf)
	err := newInfra.Up()
	if err != nil {
		t.Fatalf("Did not expect an error from newInfra.Up(): %#v", err)
	}
}

// NOTE: the repetitive instantiation should be eliminated once the other
// functions are implemented.
func TestInfraUpError(t *testing.T) {
	conf := newMockConfig("libvirt")
	newProv := newMockProvider(conf, ErrMockProvider)
	newInfra := NewInfra(newProv, conf)
	err := newInfra.Up()
	if err == nil {
		t.Fatalf("Expected error: %#v from newInfra.Up()", ErrMockProvider)
	}
}

func TestInfraHalt(t *testing.T) {
	conf := newMockConfig("libvirt")
	newProv := newMockProvider(conf, nil)
	newInfra := NewInfra(newProv, conf)
	err := newInfra.Halt(false, false)
	if err != nil {
		t.Fatalf("Did not expect error from newInfra.Halt(): %#v", err)
	}
}

func TestAddInfraNewInstances(t *testing.T) {
	conf := newMockConfig("libvirt")
	newProv := newMockProvider(conf, nil)
	newInfra := NewInfra(newProv, conf)
	err := newInfra.NewInstances()
	if err != nil {
		t.Fatalf("Did not expect error from newInfra.NewInstances(): %#v", err)
	}
}

func TestDeleteInfraInstances(t *testing.T) {
	conf := newMockConfig("libvirt")
	newProv := newMockProvider(conf, nil)
	newInfra := NewInfra(newProv, conf)
	err := newInfra.DeleteInstances("")
	if err != nil {
		t.Fatalf("Did not expect error from newInfra.DeleteInstances(): %#v", err)
	}
}
