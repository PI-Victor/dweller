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

// Resources is an interface that all resources need to implement in order for
// the controller to create them
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
	errChan := make(chan error, 1)
	finished := make(chan struct{}, 0)
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
		close(finished)
	}()

	err := <-errChan
	logrus.Infof("Finished resource creation with err: %#v", err)
	return err
}

// DeleteResources selectively deletes resources either specific to domain
// definition or generic resources such as network and storage-pool
func (lc *LibvirtController) DeleteResources() error {
	return nil
}
