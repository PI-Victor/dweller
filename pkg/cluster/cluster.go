package cluster

// Cluster contains information specific to a cloud/private cluster
type Cluster struct {
	Spec *Spec
}

// Spec contains a generic cluster specification information
type Spec struct {
	Master Node
	Nodes  []*Node
}

// Node holds specific information about a node in the cluster
type Node struct {
	Master bool
}
