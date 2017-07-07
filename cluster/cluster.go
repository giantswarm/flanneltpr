package cluster

// Cluster holds the configuration of the cluster for which flannel is
// configured.
type Cluster struct {
	// ID is the cluster's ID. It is used for logging purposes.
	ID string `json:"id" yaml:"id"`

	// Namespace is the namespace cluster's resources are running in.
	// Namespace is monitored to be destroyed before bridges cleanup.
	Namespace string `json:"namespace" yaml:"namespace"`
}
