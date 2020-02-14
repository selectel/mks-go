package cluster

import "time"

// View represents an unmarshalled cluster body from an API response.
type View struct {
	// ID is the identifier of the cluster.
	ID string `json:"id"`

	// CreatedAt is the timestamp in UTC timezone of when the cluster has been created.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt is the timestamp in UTC timezone of when the cluster has been updated.
	UpdatedAt time.Time `json:"updated_at"`

	// Name represents the name of the cluster.
	Name string `json:"name"`

	// Status represents current status of the cluster.
	Status string `json:"status"`

	// ProjectID contains reference to the project of the cluster.
	ProjectID string `json:"project_id"`

	// NetworkID contains reference to the network of the cluster.
	NetworkID string `json:"network_id"`

	// SubnetID contains reference to the subnet of the cluster.
	SubnetID string `json:"subnet_id"`

	// KubeAPIIP represents the IP of the Kubernetes API.
	KubeAPIIP string `json:"kube_api_ip"`

	// KubeVersion represents the current Kubernetes version of the cluster.
	KubeVersion string `json:"kube_version"`

	// Region represents the region of where the cluster is located.
	Region string `json:"region"`

	// AdditionalSoftware represents information about additional software installed in the cluster.
	AdditionalSoftware map[string]interface{} `json:"additional_software"`

	// PKITreeUpdatedAt represents the timestamp in UTC timezone of when the PKI-tree of the cluster
	// has been updated.
	PKITreeUpdatedAt time.Time `json:"pki_tree_updated_at"`

	// MaintenanceWindowStart represents UTC time in "hh:mm:ss" format of when the cluster will start its
	// maintenance tasks.
	MaintenanceWindowStart string `json:"maintenance_window_start"`

	// MaintenanceWindowEnd represents UTC time in "hh:mm:ss" format of when the cluster will end its
	// maintenance tasks.
	MaintenanceWindowEnd string `json:"maintenance_window_end"`

	// MaintenanceLastStart is the timestamp in UTC timezone of the last cluster maintenance start.
	MaintenanceLastStart time.Time `json:"maintenance_last_start"`

	// EnableAutorepair reflects if worker nodes are allowed to be reinstalled automatically
	// in case of their unavailability or unhealthiness.
	EnableAutorepair bool `json:"enable_autorepair"`
}
