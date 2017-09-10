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
	"sync"

	"github.com/Sirupsen/logrus"
)

// LibvirtController handles creation and deletion of libvirt resources.
type LibvirtController struct {
	Resources []resources
	client    LibvirtClient
}

// Resources is an interface for libvirt domain resources
type resources interface {
	Marshal() (string, error)
}

func newController(client LibvirtClient, res ...resources) *LibvirtController {
	return &LibvirtController{
		Resources: res,
		client:    client,
	}
}

// CreateResources tries to create a new instance of a resource.
func (lc *LibvirtController) CreateResources() error {
	logrus.Debugf("Got here: %#v", lc.Resources)
	errChan := make(chan error, len(lc.Resources))
	var wg sync.WaitGroup

	for _, resource := range lc.Resources {
		logrus.Info("Starting resource creation")
		wg.Add(1)
		go func(wg *sync.WaitGroup, res resources) {
			defer wg.Done()
			r, err := res.Marshal()
			if err != nil {
				errChan <- err
			}
			_, err = lc.client.DomainDefineXML(r)
			errChan <- err
		}(&wg, resource)
	}
	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			logrus.Warningf("Failed to instantiate resource: %#v", err)
			return err
		}
	}
	return nil
}

// DeleteResources selectively deletes resources either specific to domain
// definition or generic resources such as network and storage-pool
func (lc *LibvirtController) DeleteResources() error {
	return nil
}
