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

func (m *mockProvider) NewInfra(c *config.Infra) error {
	return m.err
}

func (m *mockProvider) HaltInfra(c *config.Infra) error {
	return m.err
}

func (m *mockProvider) RegisterInstances(c *config.Infra) error {
	return m.err
}

func (m *mockProvider) DestroyInstances(c *config.Infra) error {
	return m.err
}

func newMockConfig(prov string) *config.Infra {
	lvURI := "qemu:///sytem"
	return &config.Infra{
		LibvirtURI:   &lvURI,
		ProviderName: &prov,
	}
}

func newMockProvider(conf *config.Infra, err error) Provider {
	return &mockProvider{
		ProviderName: *conf.ProviderName,
		err:          err,
	}
}

func TestNewInvalidProvider(t *testing.T) {
	_, err := NewProvider(newMockConfig("invalidProvider"))
	if err != ErrUnkownProvider {
		t.Errorf("Expected error to be: %#v, got: %#v", ErrUnkownProvider, err)
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
