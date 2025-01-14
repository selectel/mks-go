package nodegroup

import (
	"encoding/json"
	"time"

	"github.com/selectel/mks-go/pkg/v1/node"
)

// Status represents custom type for various nodegroup statuses.
type Status string

const (
	StatusActive               Status = "ACTIVE"
	StatusPendingCreate        Status = "PENDING_CREATE"
	StatusPendingUpdate        Status = "PENDING_UPDATE"
	StatusPendingDelete        Status = "PENDING_DELETE"
	StatusPendingScaleUp       Status = "PENDING_SCALE_UP"
	StatusPendingScaleDown     Status = "PENDING_SCALE_DOWN"
	StatusPendingNodeReinstall Status = "PENDING_NODE_REINSTALL"
	StatusUnknown              Status = "UNKNOWN"
	StatusError                Status = "ERROR"
)

func getSupportedStatuses() []Status {
	return []Status{
		StatusActive,
		StatusPendingCreate,
		StatusPendingUpdate,
		StatusPendingScaleUp,
		StatusPendingScaleDown,
		StatusPendingDelete,
		StatusPendingNodeReinstall,
		StatusError,
	}
}

func isStatusSupported(s Status) bool {
	for _, v := range getSupportedStatuses() {
		if s == v {
			return true
		}
	}

	return false
}

// BaseView represents a base struct of unmarshalled nodegroup body from an API response.
//
//nolint:maligned
type BaseView struct {
	// ID is the identifier of the nodegroup.
	ID string `json:"id"`

	// CreatedAt is the timestamp in UTC timezone of when the nodegroup has been created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt is the timestamp in UTC timezone of when the nodegroup has been updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// Status represents the current status of the nodegroup.
	Status Status `json:"-"`

	// ClusterID contains cluster identifier.
	ClusterID string `json:"cluster_id"`

	// FlavorID contains OpenStack flavor identifier for all nodes in the nodegroup.
	FlavorID string `json:"flavor_id"`

	// VolumeGB represents initial volume size in GB for each node.
	VolumeGB int `json:"volume_gb"`

	// VolumeType represents initial blockstorage volume type for each node.
	VolumeType string `json:"volume_type"`

	// LocalVolume represents if nodes use local volume.
	LocalVolume bool `json:"local_volume"`

	// AvailabilityZone represents OpenStack availability zone for all nodes in the nodegroup.
	AvailabilityZone string `json:"availability_zone"`

	// Nodes contains list of all nodes in the nodegroup.
	Nodes []*node.View `json:"nodes"`

	// Labels represents an object containing a set of Kubernetes labels that will be applied
	// for each node in the group. The keys must be user-defined.
	Labels map[string]string `json:"labels"`

	// Taints represents a list of nodegroup taints.
	Taints []Taint `json:"taints"`

	// EnableAutoscale reflects if the nodegroup is allowed to be scaled automatically.
	EnableAutoscale bool `json:"enable_autoscale"`

	// AutoscaleMinNodes represents minimum possible number of worker nodes in the nodegroup.
	AutoscaleMinNodes int `json:"autoscale_min_nodes"`

	// AutoscaleMaxNodes represents maximum possible number of worker nodes in the nodegroup.
	AutoscaleMaxNodes int `json:"autoscale_max_nodes"`

	// NodegroupType represents nodegroup type.
	NodegroupType string `json:"nodegroup_type"`

	// InstallNvidiaDevicePlugin indicates if nvidia device plugin installation was requested.
	InstallNvidiaDevicePlugin bool `json:"install_nvidia_device_plugin"`

	// Preemptible indicates if preemptible instances are used.
	Preemptible bool `json:"preemptible"`
}

// ListView represents an unmarshalled nodegroup body from the list API response.
type ListView struct {
	BaseView

	// AvailableAdditionalInfo provides additional information about nodegroup like userdata, etc.
	// Usually it's large volume of data and here we only show presence of this info.
	AvailableAdditionalInfo map[string]bool `json:"available_additional_info"`
}

// GetView represents an unmarshalled nodegroup body from the get API response.
type GetView struct {
	BaseView

	// UserData represents base64 data which is used to pass a script that worker nodes run on boot.
	UserData string `json:"user_data"`
}

func (result *GetView) UnmarshalJSON(b []byte) error {
	type tmp GetView
	var s struct {
		tmp
		Status Status `json:"status"`
	}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*result = GetView(s.tmp)

	// Check nodegroup status.
	if isStatusSupported(s.Status) {
		result.Status = s.Status
	} else {
		result.Status = StatusUnknown
	}

	return nil
}

// TaintEffect represents an effect of the node's taint.
type TaintEffect string

const (
	NoScheduleEffect       TaintEffect = "NoSchedule"
	NoExecuteEffect        TaintEffect = "NoExecute"
	PreferNoScheduleEffect TaintEffect = "PreferNoSchedule"
)

// Taint represents k8s node's taint.
type Taint struct {
	// Key is the key of the taint.
	Key string `json:"key"`

	// Value is the value of the taint.
	Value string `json:"value"`

	// Effect is the effect of the taint.
	Effect TaintEffect `json:"effect"`
}
