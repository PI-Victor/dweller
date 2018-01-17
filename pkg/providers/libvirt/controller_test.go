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

type mockController struct {
	err error
}

func (f *mockController) CreateResources() error {
	return f.err
}

func (f *mockController) DeleteResources() error {
	return f.err
}

func newmockController(err error) *mockController {
	return &mockController{
		err: err,
	}
}

type mockResource struct {
	testString string
	err        error
}

func (f *mockResource) Marshal() (string, error) {
	return f.testString, f.err
}

func newmockResource() mockResource {
	return mockResource{}
}

func newMockLibvirtController(err error, res ...resource) *Controller {
	return &Controller{
		client:    newMockLibvirtClient(nil, err),
		Resources: append([]resource{}, res...),
	}
}

func TestNewControllerInstance(t *testing.T) {
	if c := newController(newMockLibvirtClient(nil, nil), nil, nil); c == nil {
		t.Fatalf("Expected a new controller, got nil")
	}
}

func TestDeleteResources(t *testing.T) {
	nc := newMockLibvirtController(nil, nil, nil)
	if err := nc.DeleteResources(); err != nil {
		t.Errorf("Expected a nil error, got: %#v", err)
	}
}
