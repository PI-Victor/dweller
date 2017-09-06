package libvirt

import (
	"github.com/cloudflavor/dweller/pkg/providers/controllers"
)

// LibvirtController handles creation and deletion of libvirt resources.
type LibvirtController struct {
	Resources []controllers.Resources
}

func newController(res ...controllers.Resources) *LibvirtController {
	return &LibvirtController{
		Resources: res,
	}
}

// NewInstance tries to create a new instance of a resource.
func (lc *LibvirtController) NewInstance() error {
	return nil
}
