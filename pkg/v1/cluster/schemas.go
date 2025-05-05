package cluster

import (
	"encoding/json"
	"time"
)

// Status represents custom type for various cluster statuses.
type Status string

const (
	StatusActive                             Status = "ACTIVE"
	StatusPendingCreate                      Status = "PENDING_CREATE"
	StatusPendingUpdate                      Status = "PENDING_UPDATE"
	StatusPendingUpgrade                     Status = "PENDING_UPGRADE"
	StatusPendingRotateCerts                 Status = "PENDING_ROTATE_CERTS"
	StatusPendingDelete                      Status = "PENDING_DELETE"
	StatusPendingResize                      Status = "PENDING_RESIZE"
	StatusPendingNodeReinstall               Status = "PENDING_NODE_REINSTALL"
	StatusPendingUpgradePatchVersion         Status = "PENDING_UPGRADE_PATCH_VERSION"
	StatusPendingUpgradeMinorVersion         Status = "PENDING_UPGRADE_MINOR_VERSION"
	StatusPendingUpdateNodegroup             Status = "PENDING_UPDATE_NODEGROUP"
	StatusPendingUpgradeMastersConfiguration Status = "PENDING_UPGRADE_MASTERS_CONFIGURATION"
	StatusPendingUpgradeClusterConfiguration Status = "PENDING_UPGRADE_CLUSTER_CONFIGURATION"
	StatusMaintenance                        Status = "MAINTENANCE"
	StatusError                              Status = "ERROR"
	StatusUnknown                            Status = "UNKNOWN"
)

func getSupportedStatuses() []Status {
	return []Status{
		StatusActive,
		StatusPendingCreate,
		StatusPendingUpdate,
		StatusPendingUpgrade,
		StatusPendingRotateCerts,
		StatusPendingDelete,
		StatusPendingResize,
		StatusPendingNodeReinstall,
		StatusPendingUpgradePatchVersion,
		StatusPendingUpgradeMinorVersion,
		StatusPendingUpdateNodegroup,
		StatusPendingUpgradeMastersConfiguration,
		StatusPendingUpgradeClusterConfiguration,
		StatusMaintenance,
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

// View represents an unmarshalled cluster body from an API response.
type View struct {
	// ID is the identifier of the cluster.
	ID string `json:"id"`

	// CreatedAt is the timestamp in UTC timezone of when the cluster has been created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt is the timestamp in UTC timezone of when the cluster has been updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// Name represents the name of the cluster.
	Name string `json:"name"`

	// Status represents current status of the cluster.
	Status Status `json:"-"`

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
	PKITreeUpdatedAt *time.Time `json:"pki_tree_updated_at"`

	// MaintenanceWindowStart represents UTC time in "hh:mm:ss" format of when the cluster will start its
	// maintenance tasks.
	MaintenanceWindowStart string `json:"maintenance_window_start"`

	// MaintenanceWindowEnd represents UTC time in "hh:mm:ss" format of when the cluster will end its
	// maintenance tasks.
	MaintenanceWindowEnd string `json:"maintenance_window_end"`

	// MaintenanceLastStart is the timestamp in UTC timezone of the last cluster maintenance start.
	MaintenanceLastStart *time.Time `json:"maintenance_last_start"`

	// EnableAutorepair reflects if worker nodes are allowed to be reinstalled automatically
	// in case of their unavailability or unhealthiness.
	EnableAutorepair bool `json:"enable_autorepair"`

	// EnablePatchVersionAutoUpgrade specifies if Kubernetes patch version of the cluster is allowed to be upgraded
	// automatically.
	EnablePatchVersionAutoUpgrade bool `json:"enable_patch_version_auto_upgrade"`

	// Zonal specifies that cluster has only a single master and that
	// control-plane is not in highly available mode.
	Zonal bool `json:"zonal"`

	// KubernetesOptions represents additional k8s options such as pod security policy,
	// feature gates, admission controllers, audit logs and oidc.
	KubernetesOptions *KubernetesOptions `json:"kubernetes_options,omitempty"`

	PrivateKubeAPI bool `json:"private_kube_api"`
}

func (result *View) UnmarshalJSON(b []byte) error {
	type tmp View
	var s struct {
		tmp
		Status Status `json:"status"`
	}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*result = View(s.tmp)

	// Check cluster status.
	if isStatusSupported(s.Status) {
		result.Status = s.Status
	} else {
		result.Status = StatusUnknown
	}

	return nil
}

// KubernetesOptions represents additional k8s options such as pod security policy,
// feature gates, admission controllers, audit logs and oidc.
type KubernetesOptions struct {
	// EnablePodSecurityPolicy indicates if PodSecurityPolicy admission controller
	// must be turned on/off.
	EnablePodSecurityPolicy bool `json:"enable_pod_security_policy"`

	// FeatureGates represents feature gates that should be enabled.
	FeatureGates []string `json:"feature_gates"`

	// AdmissionControllers represents admission controllers that should be enabled.
	AdmissionControllers []string `json:"admission_controllers"`

	// AuditLogs represents configuration of kubernetes audit logs in the cluster.
	// More: https://docs.selectel.ru/en/cloud/managed-kubernetes/clusters/logs/#configure-integration-with-external-system
	AuditLogs AuditLogs `json:"audit_logs"`

	// OIDC represents configuration to enable authorization via OpenID Connect in kubernetes cluster.
	// More: https://kubernetes.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens
	OIDC OIDC `json:"oidc"`
}

type AuditLogs struct {
	// Enabled indicates whether kubernetes audit logs should be collected
	// and pushed into SIEM system (e.g. logstash).
	// False by default.
	Enabled bool `json:"enabled"`

	// SecretName contains name of the kubernetes secret in namespace kube-system
	// with credentials of SIEM system where logs should be pushed.
	// Fields of the secret: host, port, username (optional), password (optional), ca.crt (optional).
	// This field is optional. By default, used "mks-audit-logs".
	// Secret name should be as a DNS subdomain name as defined in RFC 1123.
	// More: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names
	SecretName string `json:"secret_name"`
}

// OIDC represents parameters to connect client's OIDC provider with kubernetes.
type OIDC struct {
	// Enabled indicates whether OIDC should be turned on in the cluster.
	// False by default.
	Enabled bool `json:"enabled"`

	// ProviderName represents custom user defined name of the provider. It is not used in the cluster directly.
	// It is required when enabled = true.
	ProviderName string `json:"provider_name"`

	// IssuerURL represents URL of the provider that allows the API server to discover public signing keys.
	// Will be placed in `--oidc-issuer-url` flag.
	// It is required when enabled = true.
	IssuerURL string `json:"issuer_url"`

	// ClientID represents required client id that all tokens must be issued for.
	// Will be placed in `--oidc-client-id` flag.
	// It is required when enabled = true.
	ClientID string `json:"client_id"`

	// UsernameClaim represents optional JWT claim to use as the username. By default, `sub`.
	// Will be placed in `--oidc-username-claim` flag.
	UsernameClaim string `json:"username_claim"`

	// GroupsClaim represents optional JWT claim to use as the user's group. By default, `groups`.
	// Will be placed in `--oidc-groups-claim` flag.
	GroupsClaim string `json:"groups_claim"`

	// CACerts represent optional custom CA certs chain in X509 PEM format of provider's SSL certificate.
	// Will be written to file on masters which will be passed in `--oidc-ca-file` kube-apiserver flag.
	CACerts string `json:"ca_certs,omitempty"`
}

// KubeconfigFields is a struct that contains Kubeconfigs parsed fields and raw kubeconfig.
type KubeconfigFields struct {
	ClusterCA     string
	Server        string
	ClientCert    string
	ClientKey     string
	KubeconfigRaw string
}
