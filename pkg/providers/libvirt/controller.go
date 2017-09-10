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

// Controller handles creation and deletion of libvirt resources.
type Controller struct {
	DomainResources  []resources
	GenericResources []resources
	client           Client
}

// Resources is an interface for libvirt domain resources
type resources interface {
	Marshal() (string, error)
}

func newController(client Client, domainRes []resources, genRes ...resources) *Controller {
	return &Controller{
		DomainResources: domainRes,
		client:          client,
	}
}

// CreateResources tries to create a new instance of a resource.
func (c *Controller) CreateResources() error {
	logrus.Debugf("Got here: %#v", c.DomainResources)
	errChan := make(chan error, len(c.DomainResources))
	var wg sync.WaitGroup

	for _, resource := range c.DomainResources {
		logrus.Info("Starting resource creation")
		wg.Add(1)
		go func(wg *sync.WaitGroup, res resources) {
			defer wg.Done()
			r, err := res.Marshal()
			if err != nil {
				errChan <- err
			}
			_, err = c.client.DomainDefineXML(r)
			errChan <- err
		}(&wg, resource)
	}
	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		// NOTE: even if we return an error here, the available resources were
		// already created. This is kinda pointless, we should just warn the user that
		// the creation of the resource failed.
		// When the connection pool is implemented it can have a retry with a different
		// connection.
		if err != nil {
			logrus.Warningf("Failed to instantiate resource: %#v", err)
		}
	}
	return nil
}

// CreateGenericResources creates generic libvirt resources such as network and
// storage pools.
func (c *Controller) CreateGenericResources() error {
	return nil
}

// DeleteResources selectively deletes resources either specific to domain
// definition or generic resources such as network and storage-pool
func (c *Controller) DeleteResources() error {
	return nil
}
