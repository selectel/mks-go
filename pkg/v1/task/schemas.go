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

// Type represents custom type for various task types.
type Type string

const (
	TypeCreateCluster       Type = "CREATE_CLUSTER"
	TypeDeleteCluster       Type = "DELETE_CLUSTER"
	TypeRotateCerts         Type = "ROTATE_CERTS"      // Task to rotate PKI-tree.
	TypeNodeGroupResize     Type = "NODE_GROUP_RESIZE" // Task to resize nodes in a group.
	TypeNodeReinstall       Type = "NODE_REINSTALL"    // Task to reinstall a single node.
	TypeClusterResize       Type = "CLUSTER_RESIZE"    // Task to change amount of node-groups in a cluster.
	TypeUpgradePatchVersion Type = "UPGRADE_PATCH_VERSION"
	TypeUnknown             Type = "UNKNOWN"
)

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
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*result = View(s.tmp)

	// Check task status.
	switch s.Status {
	case StatusDone:
		result.Status = StatusDone
	case StatusInProgress:
		result.Status = StatusInProgress
	case StatusError:
		result.Status = StatusError
	default:
		result.Status = StatusUnknown
	}

	// Check task type.
	switch s.Type {
	case TypeCreateCluster:
		result.Type = TypeCreateCluster
	case TypeDeleteCluster:
		result.Type = TypeDeleteCluster
	case TypeRotateCerts:
		result.Type = TypeRotateCerts
	case TypeNodeGroupResize:
		result.Type = TypeNodeGroupResize
	case TypeNodeReinstall:
		result.Type = TypeNodeReinstall
	case TypeClusterResize:
		result.Type = TypeClusterResize
	case TypeUpgradePatchVersion:
		result.Type = TypeUpgradePatchVersion
	default:
		result.Type = TypeUnknown
	}

	return err
}
