package controllers

// ProviderController is an interface that all provider controllers must
// implement in order to manipulate resources.
type ProviderController interface {
	CreateResources() error
	DeleteResources() error
}

// Resources is an interface that all resources need to implement in order for
// the controller to create them
type Resources interface{}
