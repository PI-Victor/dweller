package controllers

// ProviderController is an interface that all provider controllers must
// implement in order to manipulate resources.
type ProviderController interface {
	CreateResources() error
	DeleteResources() error
}
