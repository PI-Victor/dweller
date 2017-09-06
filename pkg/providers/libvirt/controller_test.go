package libvirt

import (
	"testing"

	"github.com/libvirt/libvirt-go"

	"github.com/cloudflavor/dweller/pkg/providers/controllers"
)

func newMockLibvirtController(res ...controllers.Resources) *LibvirtController {
	return &LibvirtController{
		Resources: append([]controllers.Resources{}, res...),
	}
}

func TestNewControllerInstance(t *testing.T) {
	if c := newController(); c == nil {
		t.Fatalf("Expected a new controller, got nil")
	}
}

func TestNewInstance(t *testing.T) {
	nc := newMockLibvirtController(
		[]controllers.Resources{
			libvirt.Domain{},
		},
	)
	if n := nc.NewInstance(); n != nil {
		t.Errorf("Expected NewInstance to return nil, got: %#v", n)
	}
}
