package config

// Infra contains generic configuration that is used by the infrastructure.
type Infra struct {
	// Libvirt provider configuration
	LibvirtURI   *string
	ProviderName string
	Workers      int
}
