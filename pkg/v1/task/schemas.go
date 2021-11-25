package task

import (
	"encoding/json"
	"time"
)

// Status represents custom type for various task statuses.
type Status string

const (
	StatusInProgress Status = "IN_PROGRESS"
	StatusDone       Status = "DONE"
	StatusError      Status = "ERROR"
	StatusUnknown    Status = "UNKNOWN"
)

func getSupportedStatuses() []Status {
	return []Status{StatusInProgress, StatusDone, StatusError, StatusUnknown}
}

func isSupportedStatus(s Status) bool {
	for _, v := range getSupportedStatuses() {
		if s == v {
			return true
		}
	}

	return false
}

// Type represents custom type for various task types.
type Type string

const (
	TypeCreateCluster               Type = "CREATE_CLUSTER"
	TypeDeleteCluster               Type = "DELETE_CLUSTER"
	TypeRotateCerts                 Type = "ROTATE_CERTS"      // Task to rotate PKI-tree.
	TypeNodeGroupResize             Type = "NODE_GROUP_RESIZE" // Task to resize nodes in a group.
	TypeNodeReinstall               Type = "NODE_REINSTALL"    // Task to reinstall a single node.
	TypeClusterResize               Type = "CLUSTER_RESIZE"    // Task to change amount of node-groups in a cluster.
	TypeUpgradePatchVersion         Type = "UPGRADE_PATCH_VERSION"
	TypeUpgradeMinorVersion         Type = "UPGRADE_MINOR_VERSION"
	TypeUpdateNodegroupLabels       Type = "UPDATE_NODEGROUP_LABELS"
	TypeUpgradeMastersConfiguration Type = "UPGRADE_MASTERS_CONFIGURATION"
	TypeUpgradeClusterConfiguration Type = "UPGRADE_CLUSTER_CONFIGURATION"
	TypeUnknown                     Type = "UNKNOWN"
)

func getSupportedTaskTypes() []Type {
	return []Type{
		TypeCreateCluster, TypeDeleteCluster, TypeRotateCerts, TypeNodeGroupResize,
		TypeNodeReinstall, TypeClusterResize, TypeUpgradePatchVersion, TypeUpgradeMinorVersion,
		TypeUpdateNodegroupLabels, TypeUpgradeMastersConfiguration, TypeUpgradeClusterConfiguration, TypeUnknown,
	}
}

func isTaskTypeSupported(t Type) bool {
	for _, v := range getSupportedTaskTypes() {
		if t == v {
			return true
		}
	}

	return false
}

// View represents an unmarshalled cluster task body from an API response.
type View struct {
	// ID is the identifier of the task.
	ID string `json:"id"`

	// StartedAt is the timestamp in UTC timezone of when the task has been started.
	StartedAt *time.Time `json:"started_at"`

	// UpdatedAt is the timestamp in UTC timezone of when the task has been updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// ClusterID contains cluster identifier.
	ClusterID string `json:"cluster_id"`

	// Status represents current status of the task.
	Status Status `json:"-"`

	// Task represents task's type.
	Type Type `json:"-"`
}

func (result *View) UnmarshalJSON(b []byte) error {
	type tmp View
	var s struct {
		tmp
		Status Status `json:"status"`
		Type   Type   `json:"type"`
	}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*result = View(s.tmp)

	// Check task status.
	if isSupportedStatus(s.Status) {
		result.Status = s.Status
	} else {
		result.Status = StatusUnknown
	}

	// Check task type.
	if isTaskTypeSupported(s.Type) {
		result.Type = s.Type
	} else {
		result.Type = TypeUnknown
	}

	return nil
}
