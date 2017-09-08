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
			logrus.Warning("Failed to instantiate resource: %#v", err)
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
