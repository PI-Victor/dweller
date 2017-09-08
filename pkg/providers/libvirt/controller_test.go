package libvirt

import (
	"testing"
)

type fakeController struct {
	err error
}

func (f *fakeController) CreateResources() error {
	return f.err
}

func (f *fakeController) DeleteResources() error {
	return f.err
}

func newFakeController(err error) *fakeController {
	return &fakeController{
		err: err,
	}
}

type fakeResource struct {
	testString string
	err        error
}

func (f *fakeResource) Marshal() (string, error) {
	return f.testString, f.err
}

func newFakeResource() fakeResource {
	return fakeResource{}
}

func newMockLibvirtController(err error, res ...resources) *LibvirtController {
	return &LibvirtController{
		client:    newMockLibvirtClient(nil, err),
		Resources: append([]resources{}, res...),
	}
}

func TestNewControllerInstance(t *testing.T) {
	if c := newController(newMockLibvirtClient(nil, nil)); c == nil {
		t.Fatalf("Expected a new controller, got nil")
	}
}

func TestCreateResources(t *testing.T) {
	//controller := newMockLibvirtController(nil)
	//if
}

func TestDeleteResources(t *testing.T) {
	nc := newMockLibvirtController(nil)
	if err := nc.DeleteResources(); err != nil {
		t.Errorf("Expected a nil error, got: %#v", err)
	}
}
