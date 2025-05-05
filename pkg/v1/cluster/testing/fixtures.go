package testing

import (
	"time"

	"github.com/selectel/mks-go/pkg/testutils"
	"github.com/selectel/mks-go/pkg/v1/cluster"
	"github.com/selectel/mks-go/pkg/v1/nodegroup"
)

// testGetClusterResponseRaw represents a raw response from the Get request.
const testGetClusterResponseRaw = `
{
    "cluster": {
        "additional_software": null,
        "created_at": "2020-02-13T09:18:32.05753Z",
        "enable_autorepair": true,
        "id": "dbe7559b-55d8-4f65-9230-6a22b985ff73",
        "kube_api_ip": "203.0.113.101",
        "kube_version": "1.15.7",
        "maintenance_last_start": "2020-02-13T09:18:32.05753Z",
        "maintenance_window_end": "03:00:00",
        "maintenance_window_start": "01:00:00",
        "name": "test-cluster",
        "network_id": "74a591be-6be7-4abc-d30f-1614c0f9721c",
        "pki_tree_updated_at": "2020-02-13T09:18:32.05753Z",
        "project_id": "65044a03bede4fd0a77e5a4c882e3059",
        "region": "ru-1",
        "status": "ACTIVE",
        "subnet_id": "c872541d-2d83-419f-841d-8288201b8fb9",
        "updated_at": "2020-02-13T09:18:32.05753Z",
        "enable_patch_version_auto_upgrade": true,
        "zonal": false,
        "kubernetes_options": {
            "enable_pod_security_policy": true,
            "feature_gates": [
                "TTLAfterFinished",
                "CSIMigrationOpenStack"
			],
            "admission_controllers": [
                "NamespaceLifecycle",
                "LimitRanger"
            ],
			"audit_logs": {
				"enabled": true,
				"secret_name": "mks-audit-logs"
			},
	        "oidc": {
				"enabled": true,
				"provider_name": "keycloak",
				"client_id": "kubernetes",
				"groups_claim": "groups",
				"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
				"issuer_url": "https://example.com/",
				"username_claim": "email"
	        }
        }
    }
}
`

var clusterResponseTimestamp, _ = time.Parse(time.RFC3339, "2020-02-13T09:18:32.05753Z")

// expectedGetClusterResponse represents an unmarshalled testGetClusterResponseRaw.
var expectedGetClusterResponse = &cluster.View{
	ID:                            "dbe7559b-55d8-4f65-9230-6a22b985ff73",
	CreatedAt:                     &clusterResponseTimestamp,
	UpdatedAt:                     &clusterResponseTimestamp,
	Name:                          "test-cluster",
	Status:                        cluster.StatusActive,
	ProjectID:                     "65044a03bede4fd0a77e5a4c882e3059",
	NetworkID:                     "74a591be-6be7-4abc-d30f-1614c0f9721c",
	SubnetID:                      "c872541d-2d83-419f-841d-8288201b8fb9",
	KubeAPIIP:                     "203.0.113.101",
	KubeVersion:                   "1.15.7",
	Region:                        "ru-1",
	AdditionalSoftware:            nil,
	PKITreeUpdatedAt:              &clusterResponseTimestamp,
	MaintenanceWindowStart:        "01:00:00",
	MaintenanceWindowEnd:          "03:00:00",
	MaintenanceLastStart:          &clusterResponseTimestamp,
	EnableAutorepair:              true,
	EnablePatchVersionAutoUpgrade: true,
	KubernetesOptions: &cluster.KubernetesOptions{
		EnablePodSecurityPolicy: true,
		FeatureGates: []string{
			"TTLAfterFinished",
			"CSIMigrationOpenStack",
		},
		AdmissionControllers: []string{
			"NamespaceLifecycle",
			"LimitRanger",
		},
		AuditLogs: cluster.AuditLogs{
			Enabled:    true,
			SecretName: "mks-audit-logs",
		},
		OIDC: cluster.OIDC{
			Enabled:       true,
			ProviderName:  "keycloak",
			IssuerURL:     "https://example.com/",
			ClientID:      "kubernetes",
			UsernameClaim: "email",
			GroupsClaim:   "groups",
			CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
		},
	},
}

// testGetZonalClusterResponseRaw represents a raw zonal cluster response from the Get request.
const testGetZonalClusterResponseRaw = `
{
    "cluster": {
        "additional_software": null,
        "created_at": "2020-02-13T09:18:32.05753Z",
        "enable_autorepair": true,
        "id": "763f1f5f-951e-48b2-abae-815a75ee747c",
        "kube_api_ip": "203.0.113.101",
        "kube_version": "1.16.9",
        "maintenance_last_start": "2020-02-13T09:18:32.05753Z",
        "maintenance_window_end": "07:00:00",
        "maintenance_window_start": "04:00:00",
        "name": "test-zonal-cluster",
        "network_id": "74a591be-6be7-4abc-d30f-1614c0f9721c",
        "pki_tree_updated_at": "2020-02-13T09:18:32.05753Z",
        "project_id": "65044a03bede4fd0a77e5a4c882e3059",
        "region": "ru-3",
        "status": "PENDING_UPGRADE_MINOR_VERSION",
        "subnet_id": "c872541d-2d83-419f-841d-8288201b8fb9",
        "updated_at": "2020-02-13T09:18:32.05753Z",
        "enable_patch_version_auto_upgrade": false,
        "zonal": true,
        "kubernetes_options": {
            "enable_pod_security_policy": true,
            "feature_gates": [
                "TTLAfterFinished",
                "CSIMigrationOpenStack"
            ],
            "admission_controllers": [
                "NamespaceLifecycle",
                "LimitRanger"
            ],
			"audit_logs": {
				"enabled": true,
				"secret_name": "mks-audit-logs"
			},
	        "oidc": {
				"enabled": true,
				"provider_name": "keycloak",
				"client_id": "kubernetes",
				"groups_claim": "groups",
				"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
				"issuer_url": "https://example.com/",
				"username_claim": "email"
	        }
        }
    }
}
`

// expectedGetZonalClusterResponse represents an unmarshalled testGetZonalClusterResponseRaw.
var expectedGetZonalClusterResponse = &cluster.View{
	ID:                            "763f1f5f-951e-48b2-abae-815a75ee747c",
	CreatedAt:                     &clusterResponseTimestamp,
	UpdatedAt:                     &clusterResponseTimestamp,
	Name:                          "test-zonal-cluster",
	Status:                        cluster.StatusPendingUpgradeMinorVersion,
	ProjectID:                     "65044a03bede4fd0a77e5a4c882e3059",
	NetworkID:                     "74a591be-6be7-4abc-d30f-1614c0f9721c",
	SubnetID:                      "c872541d-2d83-419f-841d-8288201b8fb9",
	KubeAPIIP:                     "203.0.113.101",
	KubeVersion:                   "1.16.9",
	Region:                        "ru-3",
	PKITreeUpdatedAt:              &clusterResponseTimestamp,
	MaintenanceWindowStart:        "04:00:00",
	MaintenanceWindowEnd:          "07:00:00",
	MaintenanceLastStart:          &clusterResponseTimestamp,
	EnableAutorepair:              true,
	EnablePatchVersionAutoUpgrade: false,
	Zonal:                         true,
	KubernetesOptions: &cluster.KubernetesOptions{
		EnablePodSecurityPolicy: true,
		FeatureGates: []string{
			"TTLAfterFinished",
			"CSIMigrationOpenStack",
		},
		AdmissionControllers: []string{
			"NamespaceLifecycle",
			"LimitRanger",
		},
		AuditLogs: cluster.AuditLogs{
			Enabled:    true,
			SecretName: "mks-audit-logs",
		},
		OIDC: cluster.OIDC{
			Enabled:       true,
			ProviderName:  "keycloak",
			IssuerURL:     "https://example.com/",
			ClientID:      "kubernetes",
			UsernameClaim: "email",
			GroupsClaim:   "groups",
			CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
		},
	},
}

// testListClustersResponseRaw represents a raw response from the List request.
const testListClustersResponseRaw = `
{
    "clusters": [
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "dbe7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.101",
            "kube_version": "1.15.7",
            "name": "test-cluster",
            "network_id": "74a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "65044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "ACTIVE",
            "subnet_id": "c872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": true,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "2be7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.102",
            "kube_version": "1.15.7",
            "name": "test-cluster-2",
            "network_id": "24a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "25044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "PENDING_CREATE",
            "subnet_id": "2872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": false,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "3be7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.103",
            "kube_version": "1.15.7",
            "name": "test-cluster-3",
            "network_id": "34a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "35044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "PENDING_UPDATE",
            "subnet_id": "3872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": true,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "4be7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.104",
            "kube_version": "1.15.7",
            "name": "test-cluster-4",
            "network_id": "44a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "45044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "PENDING_UPGRADE",
            "subnet_id": "4872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": false,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "5be7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.105",
            "kube_version": "1.15.7",
            "name": "test-cluster-5",
            "network_id": "54a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "55044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "PENDING_ROTATE_CERTS",
            "subnet_id": "5872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": true,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "6be7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.106",
            "kube_version": "1.15.7",
            "name": "test-cluster-6",
            "network_id": "64a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "65044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "PENDING_DELETE",
            "subnet_id": "6872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": false,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "7be7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.107",
            "kube_version": "1.15.7",
            "name": "test-cluster-7",
            "network_id": "74a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "76044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "PENDING_RESIZE",
            "subnet_id": "7872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": true,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "8be7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.108",
            "kube_version": "1.15.7",
            "name": "test-cluster-8",
            "network_id": "87a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "85044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "PENDING_NODE_REINSTALL",
            "subnet_id": "8872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": false,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "9be7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.109",
            "kube_version": "1.15.7",
            "name": "test-cluster-9",
            "network_id": "94a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "95044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "PENDING_UPGRADE_PATCH_VERSION",
            "subnet_id": "9872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": true,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "9be7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.109",
            "kube_version": "1.15.7",
            "name": "test-cluster-9",
            "network_id": "94a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "95044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "PENDING_UPGRADE_MINOR_VERSION",
            "subnet_id": "9872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": true,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "9be7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.119",
            "kube_version": "1.15.7",
            "name": "test-cluster",
            "network_id": "94a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "95044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "PENDING_UPDATE_NODEGROUP",
            "subnet_id": "9872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": true,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "9be7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.120",
            "kube_version": "1.15.7",
            "name": "test-cluster",
            "network_id": "94a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "95044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "PENDING_UPGRADE_MASTERS_CONFIGURATION",
            "subnet_id": "9872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": true,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "9be7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.120",
            "kube_version": "1.15.7",
            "name": "test-cluster",
            "network_id": "94a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "95044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "PENDING_UPGRADE_CLUSTER_CONFIGURATION",
            "subnet_id": "9872541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": true,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "10e7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.110",
            "kube_version": "1.15.7",
            "name": "test-cluster-10",
            "network_id": "10a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "10044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "MAINTENANCE",
            "subnet_id": "1072541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": false,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "117559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.111",
            "kube_version": "1.15.7",
            "name": "test-cluster-11",
            "network_id": "11a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "11044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "ERROR",
            "subnet_id": "1172541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": true,
            "zonal": false,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        },
        {
            "additional_software": null,
            "created_at": "2020-02-13T09:18:32.05753Z",
            "enable_autorepair": true,
            "id": "12e7559b-55d8-4f65-9230-6a22b985ff73",
            "kube_api_ip": "203.0.113.112",
            "kube_version": "1.15.7",
            "name": "test-cluster-12",
            "network_id": "12a591be-6be7-4abc-d30f-1614c0f9721c",
            "project_id": "12044a03bede4fd0a77e5a4c882e3059",
            "region": "ru-1",
            "status": "UNKNOWN",
            "subnet_id": "1272541d-2d83-419f-841d-8288201b8fb9",
            "updated_at": "2020-02-13T09:18:32.05753Z",
            "enable_patch_version_auto_upgrade": false,
            "zonal": true,
            "kubernetes_options": {
                "enable_pod_security_policy": false,
                "feature_gates": [
                    "CSIMigrationOpenStack"
                ],
                "admission_controllers": [
                    "LimitRanger"
                ],
				"audit_logs": {
					"enabled": true,
					"secret_name": "mks-audit-logs"
				},
				"oidc": {
					"enabled": true,
					"provider_name": "keycloak",
					"client_id": "kubernetes",
					"groups_claim": "groups",
					"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
					"issuer_url": "https://example.com/",
					"username_claim": "email"
				}
            }
        }
    ]
}
`

// expectedListClustersResponse represents an unmarshalled testListClustersResponseRaw.
var expectedListClustersResponse = []*cluster.View{
	{
		ID:                            "dbe7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster",
		Status:                        cluster.StatusActive,
		ProjectID:                     "65044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "74a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "c872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.101",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: true,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "2be7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster-2",
		Status:                        cluster.StatusPendingCreate,
		ProjectID:                     "25044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "24a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "2872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.102",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: false,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "3be7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster-3",
		Status:                        cluster.StatusPendingUpdate,
		ProjectID:                     "35044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "34a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "3872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.103",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: true,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "4be7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster-4",
		Status:                        cluster.StatusPendingUpgrade,
		ProjectID:                     "45044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "44a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "4872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.104",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: false,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "5be7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster-5",
		Status:                        cluster.StatusPendingRotateCerts,
		ProjectID:                     "55044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "54a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "5872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.105",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: true,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "6be7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster-6",
		Status:                        cluster.StatusPendingDelete,
		ProjectID:                     "65044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "64a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "6872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.106",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: false,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "7be7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster-7",
		Status:                        cluster.StatusPendingResize,
		ProjectID:                     "76044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "74a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "7872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.107",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: true,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "8be7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster-8",
		Status:                        cluster.StatusPendingNodeReinstall,
		ProjectID:                     "85044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "87a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "8872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.108",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: false,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "9be7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster-9",
		Status:                        cluster.StatusPendingUpgradePatchVersion,
		ProjectID:                     "95044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "94a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "9872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.109",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: true,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "9be7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster-9",
		Status:                        cluster.StatusPendingUpgradeMinorVersion,
		ProjectID:                     "95044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "94a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "9872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.109",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: true,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "9be7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster",
		Status:                        cluster.StatusPendingUpdateNodegroup,
		ProjectID:                     "95044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "94a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "9872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.119",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: true,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "9be7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster",
		Status:                        cluster.StatusPendingUpgradeMastersConfiguration,
		ProjectID:                     "95044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "94a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "9872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.120",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: true,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "9be7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster",
		Status:                        cluster.StatusPendingUpgradeClusterConfiguration,
		ProjectID:                     "95044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "94a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "9872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.120",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: true,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "10e7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster-10",
		Status:                        cluster.StatusMaintenance,
		ProjectID:                     "10044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "10a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "1072541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.110",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: false,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "117559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster-11",
		Status:                        cluster.StatusError,
		ProjectID:                     "11044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "11a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "1172541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.111",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: true,
		Zonal:                         false,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
	{
		ID:                            "12e7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:                     &clusterResponseTimestamp,
		UpdatedAt:                     &clusterResponseTimestamp,
		Name:                          "test-cluster-12",
		Status:                        cluster.StatusUnknown,
		ProjectID:                     "12044a03bede4fd0a77e5a4c882e3059",
		NetworkID:                     "12a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:                      "1272541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:                     "203.0.113.112",
		KubeVersion:                   "1.15.7",
		Region:                        "ru-1",
		AdditionalSoftware:            nil,
		PKITreeUpdatedAt:              nil,
		MaintenanceWindowStart:        "",
		MaintenanceWindowEnd:          "",
		MaintenanceLastStart:          nil,
		EnableAutorepair:              true,
		EnablePatchVersionAutoUpgrade: false,
		Zonal:                         true,
		KubernetesOptions: &cluster.KubernetesOptions{
			EnablePodSecurityPolicy: false,
			FeatureGates:            []string{"CSIMigrationOpenStack"},
			AdmissionControllers:    []string{"LimitRanger"},
			AuditLogs: cluster.AuditLogs{
				Enabled:    true,
				SecretName: "mks-audit-logs",
			},
			OIDC: cluster.OIDC{
				Enabled:       true,
				ProviderName:  "keycloak",
				IssuerURL:     "https://example.com/",
				ClientID:      "kubernetes",
				UsernameClaim: "email",
				GroupsClaim:   "groups",
				CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
			},
		},
	},
}

// testCreateClusterOptsRaw represents marshalled options for the Create request.
const testCreateClusterOptsRaw = `
{
    "cluster": {
        "name": "test-cluster-0",
        "kube_version": "1.15.7",
        "region": "ru-1",
        "nodegroups": [
            {
                "count": 1,
                "cpus": 1,
                "ram_mb": 2048,
                "volume_gb": 10,
                "volume_type": "fast.ru-1b",
                "keypair_name": "ssh-key",
                "availability_zone": "ru-1b",
                "labels": {
                  "test-label-key": "test-label-value"
                },
                "taints": [
                  {
                    "key": "test-key-0",
                    "value": "test-value-0",
                    "effect": "NoSchedule"
                  }
                ]
            }
        ],
        "kubernetes_options": {
            "enable_pod_security_policy": false,
            "feature_gates": [
                "CSIMigrationOpenStack"
            ],
            "admission_controllers": [
                "LimitRanger"
            ],
			"audit_logs": {
				"enabled": true,
				"secret_name": "mks-audit-logs"
			},
			"oidc": {
				"enabled": true,
				"provider_name": "keycloak",
				"client_id": "kubernetes",
				"groups_claim": "groups",
				"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
				"issuer_url": "https://example.com/",
				"username_claim": "email"
			}
        }
    }
}
`

// testCreateClusterOpts represents options for the Create request.
var testCreateClusterOpts = &cluster.CreateOpts{
	Name:        "test-cluster-0",
	KubeVersion: "1.15.7",
	Region:      "ru-1",
	Nodegroups: []*nodegroup.CreateOpts{
		{
			Count:            1,
			CPUs:             1,
			RAMMB:            2048,
			VolumeGB:         10,
			VolumeType:       "fast.ru-1b",
			KeypairName:      "ssh-key",
			AvailabilityZone: "ru-1b",
			Labels: map[string]string{
				"test-label-key": "test-label-value",
			},
			Taints: []nodegroup.Taint{
				{
					Key:    "test-key-0",
					Value:  "test-value-0",
					Effect: nodegroup.NoScheduleEffect,
				},
			},
		},
	},
	KubernetesOptions: &cluster.KubernetesOptions{
		EnablePodSecurityPolicy: false,
		FeatureGates:            []string{"CSIMigrationOpenStack"},
		AdmissionControllers:    []string{"LimitRanger"},
		AuditLogs: cluster.AuditLogs{
			Enabled:    true,
			SecretName: "mks-audit-logs",
		},
		OIDC: cluster.OIDC{
			Enabled:       true,
			ProviderName:  "keycloak",
			IssuerURL:     "https://example.com/",
			ClientID:      "kubernetes",
			UsernameClaim: "email",
			GroupsClaim:   "groups",
			CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
		},
	},
}

// testCreateClusterResponseRaw represents a raw response from the Create request.
const testCreateClusterResponseRaw = `
{
    "cluster": {
        "additional_software": null,
        "created_at": "2020-02-13T09:18:32.05753Z",
        "enable_autorepair": true,
        "enable_patch_version_auto_upgrade": true,
        "id": "1f4dfa6a-0468-45e0-be13-74c6481820f5",
        "kube_api_ip": "",
        "kube_version": "1.15.7",
        "maintenance_last_start": "2020-02-13T09:18:32.05753Z",
        "maintenance_window_end": "03:00:00",
        "maintenance_window_start": "01:00:00",
        "name": "test-cluster-0",
        "network_id": "",
        "pki_tree_updated_at": null,
        "project_id": "69744a03bebe4fd0a77e5a4c882e3059",
        "region": "ru-1",
        "status": "PENDING_CREATE",
        "subnet_id": "",
        "updated_at": null,
        "zonal": false,
        "private_kube_api": false,
        "kubernetes_options": {
            "enable_pod_security_policy": false,
            "feature_gates": [
                "CSIMigrationOpenStack"
      	    ],
            "admission_controllers": [
                "LimitRanger"
            ],
			"audit_logs": {
				"enabled": true,
				"secret_name": "mks-audit-logs"
			},
			"oidc": {
				"enabled": true,
				"provider_name": "keycloak",
				"client_id": "kubernetes",
				"groups_claim": "groups",
				"ca_certs": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
				"issuer_url": "https://example.com/",
				"username_claim": "email"
			}
        }
    }
}
`

// expectedCreateClusterResponse represents an unmarshalled testCreateClusterResponseRaw.
var expectedCreateClusterResponse = &cluster.View{
	ID:                            "1f4dfa6a-0468-45e0-be13-74c6481820f5",
	CreatedAt:                     &clusterResponseTimestamp,
	UpdatedAt:                     nil,
	Name:                          "test-cluster-0",
	Status:                        "PENDING_CREATE",
	ProjectID:                     "69744a03bebe4fd0a77e5a4c882e3059",
	NetworkID:                     "",
	SubnetID:                      "",
	KubeAPIIP:                     "",
	KubeVersion:                   "1.15.7",
	Region:                        "ru-1",
	AdditionalSoftware:            nil,
	PKITreeUpdatedAt:              nil,
	MaintenanceWindowStart:        "01:00:00",
	MaintenanceWindowEnd:          "03:00:00",
	MaintenanceLastStart:          &clusterResponseTimestamp,
	EnableAutorepair:              true,
	EnablePatchVersionAutoUpgrade: true,
	Zonal:                         false,
	KubernetesOptions: &cluster.KubernetesOptions{
		EnablePodSecurityPolicy: false,
		FeatureGates:            []string{"CSIMigrationOpenStack"},
		AdmissionControllers:    []string{"LimitRanger"},
		AuditLogs: cluster.AuditLogs{
			Enabled:    true,
			SecretName: "mks-audit-logs",
		},
		OIDC: cluster.OIDC{
			Enabled:       true,
			ProviderName:  "keycloak",
			IssuerURL:     "https://example.com/",
			ClientID:      "kubernetes",
			UsernameClaim: "email",
			GroupsClaim:   "groups",
			CACerts:       "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
		},
	},
	PrivateKubeAPI: false,
}

// testCreateClusterEnableBoolsOptsRaw represents marshalled options for the Create request
// with enabled booleans opts.
const testCreateClusterEnableBoolsOptsRaw = `
{
    "cluster": {
        "name": "test-cluster-0",
        "kube_version": "1.15.7",
        "region": "ru-1",
        "enable_autorepair": true,
        "enable_patch_version_auto_upgrade": true,
        "nodegroups": [
            {
                "count": 1,
                "cpus": 1,
                "ram_mb": 2048,
                "volume_gb": 10,
                "volume_type": "fast.ru-1b",
                "keypair_name": "ssh-key",
                "availability_zone": "ru-1b",
                "labels": {
                  "test-label-key": "test-label-value"
                },
                "taints": []
            }
        ]
    }
}
`

// testCreateClusterEnableBoolsOpts represents options for the Create request with enabled booleans opts.
var testCreateClusterEnableBoolsOpts = &cluster.CreateOpts{
	Name:                          "test-cluster-0",
	KubeVersion:                   "1.15.7",
	Region:                        "ru-1",
	EnableAutorepair:              testutils.BoolToPtr(true),
	EnablePatchVersionAutoUpgrade: testutils.BoolToPtr(true),
	Nodegroups: []*nodegroup.CreateOpts{
		{
			Count:            1,
			CPUs:             1,
			RAMMB:            2048,
			VolumeGB:         10,
			VolumeType:       "fast.ru-1b",
			KeypairName:      "ssh-key",
			AvailabilityZone: "ru-1b",
			Labels: map[string]string{
				"test-label-key": "test-label-value",
			},
			Taints: []nodegroup.Taint{},
		},
	},
}

// testCreateClusterDisableBoolsOptsRaw represents marshalled options for the Create request
// with disabled booleans opts.
const testCreateClusterDisableBoolsOptsRaw = `
{
    "cluster": {
        "name": "test-cluster-0",
        "kube_version": "1.15.7",
        "region": "ru-1",
        "enable_autorepair": false,
        "enable_patch_version_auto_upgrade": false,
        "nodegroups": [
            {
                "count": 1,
                "cpus": 1,
                "ram_mb": 2048,
                "volume_gb": 10,
                "volume_type": "fast.ru-1b",
                "keypair_name": "ssh-key",
                "availability_zone": "ru-1b",
                "labels": {
                  "test-label-key": "test-label-value"
                },
                "taints": []
            }
        ],
        "zonal": false,
        "private_kube_api": false
    }
}
`

// testCreateClusterDisableBoolsOpts represents options for the Create request with disabled booleans opts.
var testCreateClusterDisableBoolsOpts = &cluster.CreateOpts{
	Name:                          "test-cluster-0",
	KubeVersion:                   "1.15.7",
	Region:                        "ru-1",
	EnableAutorepair:              testutils.BoolToPtr(false),
	EnablePatchVersionAutoUpgrade: testutils.BoolToPtr(false),
	Zonal:                         testutils.BoolToPtr(false),
	Nodegroups: []*nodegroup.CreateOpts{
		{
			Count:            1,
			CPUs:             1,
			RAMMB:            2048,
			VolumeGB:         10,
			VolumeType:       "fast.ru-1b",
			KeypairName:      "ssh-key",
			AvailabilityZone: "ru-1b",
			Labels: map[string]string{
				"test-label-key": "test-label-value",
			},
			Taints: []nodegroup.Taint{},
		},
	},
	PrivateKubeAPI: testutils.BoolToPtr(false),
}

// testCreateClusterDisableBoolsResponseRaw represents a raw response from the Create request
// with disabled booleans opts.
const testCreateClusterDisableBoolsResponseRaw = `
{
    "cluster": {
        "additional_software": null,
        "created_at": "2020-02-13T09:18:32.05753Z",
        "enable_autorepair": false,
        "enable_patch_version_auto_upgrade": false,
        "id": "1f4dfa6a-0468-45e0-be13-74c6481820f5",
        "kube_api_ip": "",
        "kube_version": "1.15.7",
        "maintenance_last_start": "2020-02-13T09:18:32.05753Z",
        "maintenance_window_end": "03:00:00",
        "maintenance_window_start": "01:00:00",
        "name": "test-cluster-0",
        "network_id": "",
        "pki_tree_updated_at": null,
        "project_id": "69744a03bebe4fd0a77e5a4c882e3059",
        "region": "ru-1",
        "status": "PENDING_CREATE",
        "subnet_id": "",
        "updated_at": null,
        "zonal": false,
        "private_kube_api": false
    }
}
`

// expectedCreateClusterDisableBoolsResponse represents an unmarshalled testCreateClusterResponseRaw
// with disabled booleans opts.
var expectedCreateClusterDisableBoolsResponse = &cluster.View{
	ID:                            "1f4dfa6a-0468-45e0-be13-74c6481820f5",
	CreatedAt:                     &clusterResponseTimestamp,
	UpdatedAt:                     nil,
	Name:                          "test-cluster-0",
	Status:                        "PENDING_CREATE",
	ProjectID:                     "69744a03bebe4fd0a77e5a4c882e3059",
	NetworkID:                     "",
	SubnetID:                      "",
	KubeAPIIP:                     "",
	KubeVersion:                   "1.15.7",
	Region:                        "ru-1",
	AdditionalSoftware:            nil,
	PKITreeUpdatedAt:              nil,
	MaintenanceWindowStart:        "01:00:00",
	MaintenanceWindowEnd:          "03:00:00",
	MaintenanceLastStart:          &clusterResponseTimestamp,
	EnableAutorepair:              false,
	EnablePatchVersionAutoUpgrade: false,
	Zonal:                         false,
	PrivateKubeAPI:                false,
}

// testCreateZonalClusterOptsRaw represents marshalled options for the Create request
// with zonal attribute set to true.
const testCreateZonalClusterOptsRaw = `
{
    "cluster": {
        "name": "test-zonal-cluster-0",
        "kube_version": "1.16.9",
        "region": "ru-3",
        "nodegroups": [
            {
                "count": 1,
                "cpus": 1,
                "ram_mb": 2048,
                "volume_gb": 10,
                "volume_type": "fast.ru-3a",
                "keypair_name": "ssh-key",
                "availability_zone": "ru-3a",
                "labels": {
                  "test-label-key": "test-label-value"
                },
                "taints": []
            }
        ],
        "zonal": true
    }
}
`

// testCreateZonalClusterOpts represents options for the Create request with zonal attribute set to true.
var testCreateZonalClusterOpts = &cluster.CreateOpts{
	Name:        "test-zonal-cluster-0",
	KubeVersion: "1.16.9",
	Region:      "ru-3",
	Nodegroups: []*nodegroup.CreateOpts{
		{
			Count:            1,
			CPUs:             1,
			RAMMB:            2048,
			VolumeGB:         10,
			VolumeType:       "fast.ru-3a",
			KeypairName:      "ssh-key",
			AvailabilityZone: "ru-3a",
			Labels: map[string]string{
				"test-label-key": "test-label-value",
			},
			Taints: []nodegroup.Taint{},
		},
	},
	Zonal: testutils.BoolToPtr(true),
}

// testCreateZonalClusterResponseRaw represents a raw response from the Create zonal cluster request.
const testCreateZonalClusterResponseRaw = `
{
    "cluster": {
        "additional_software": null,
        "created_at": "2020-02-13T09:18:32.05753Z",
        "enable_autorepair": true,
        "enable_patch_version_auto_upgrade": false,
        "id": "effe751d-501a-4b06-8e23-3f686dbfccf6",
        "kube_api_ip": "",
        "kube_version": "1.16.9",
        "maintenance_last_start": "2020-02-13T09:18:32.05753Z",
        "maintenance_window_end": "03:00:00",
        "maintenance_window_start": "01:00:00",
        "name": "test-zonal-cluster-0",
        "network_id": "",
        "pki_tree_updated_at": null,
        "project_id": "69744a03bebe4fd0a77e5a4c882e3059",
        "region": "ru-3",
        "status": "PENDING_CREATE",
        "subnet_id": "",
        "updated_at": null,
        "zonal": true,
        "private_kube_api": false
    }
}
`

// expectedCreateZonalClusterResponse represents an unmarshalled testCreateZonalClusterResponseRaw.
var expectedCreateZonalClusterResponse = &cluster.View{
	ID:                            "effe751d-501a-4b06-8e23-3f686dbfccf6",
	CreatedAt:                     &clusterResponseTimestamp,
	UpdatedAt:                     nil,
	Name:                          "test-zonal-cluster-0",
	Status:                        "PENDING_CREATE",
	ProjectID:                     "69744a03bebe4fd0a77e5a4c882e3059",
	NetworkID:                     "",
	SubnetID:                      "",
	KubeAPIIP:                     "",
	KubeVersion:                   "1.16.9",
	Region:                        "ru-3",
	AdditionalSoftware:            nil,
	PKITreeUpdatedAt:              nil,
	MaintenanceWindowStart:        "01:00:00",
	MaintenanceWindowEnd:          "03:00:00",
	MaintenanceLastStart:          &clusterResponseTimestamp,
	EnableAutorepair:              true,
	EnablePatchVersionAutoUpgrade: false,
	Zonal:                         true,
}

// testCreatePrivateKubeAPIClusterOptsRaw represents marshalled options for the Create request
// with private kube API attribute set to true.
const testCreatePrivateKubeAPIClusterOptsRaw = `
{
    "cluster": {
        "name": "test-private-cluster-0",
        "kube_version": "1.16.9",
        "region": "ru-3",
        "enable_autorepair": true,
        "enable_patch_version_auto_upgrade": false,
        "nodegroups": [
            {
                "count": 1,
                "cpus": 1,
                "ram_mb": 2048,
                "volume_gb": 10,
                "volume_type": "fast.ru-3a",
                "keypair_name": "ssh-key",
                "availability_zone": "ru-3a",
                "labels": {
                  "test-label-key": "test-label-value"
                },
                "taints": []
            }
        ],
        "zonal": false,
        "private_kube_api": true
    }
}
`

// testCreatePrivateKubeAPIClusterOpts represents options for the Create request with private kube API attribute set to true.
var testCreatePrivateKubeAPIClusterOpts = &cluster.CreateOpts{
	Name:                          "test-private-cluster-0",
	KubeVersion:                   "1.16.9",
	Region:                        "ru-3",
	EnableAutorepair:              testutils.BoolToPtr(true),
	EnablePatchVersionAutoUpgrade: testutils.BoolToPtr(false),
	Nodegroups: []*nodegroup.CreateOpts{
		{
			Count:            1,
			CPUs:             1,
			RAMMB:            2048,
			VolumeGB:         10,
			VolumeType:       "fast.ru-3a",
			KeypairName:      "ssh-key",
			AvailabilityZone: "ru-3a",
			Labels: map[string]string{
				"test-label-key": "test-label-value",
			},
			Taints: []nodegroup.Taint{},
		},
	},
	PrivateKubeAPI: testutils.BoolToPtr(true),
	Zonal:          testutils.BoolToPtr(false),
}

// testCreatePrivateKubeAPIClusterResponseRaw represents a raw response from the Create cluster with private kube API request.
const testCreatePrivateKubeAPIClusterResponseRaw = `
{
    "cluster": {
        "additional_software": null,
        "created_at": "2020-02-13T09:18:32.05753Z",
        "enable_autorepair": true,
        "enable_patch_version_auto_upgrade": false,
        "id": "effe751d-501a-4b06-8e23-3f686dbfccf6",
        "kube_api_ip": "",
        "kube_version": "1.16.9",
        "maintenance_last_start": "2020-02-13T09:18:32.05753Z",
        "maintenance_window_end": "03:00:00",
        "maintenance_window_start": "01:00:00",
        "name": "test-private-cluster-0",
        "network_id": "",
        "pki_tree_updated_at": null,
        "project_id": "69744a03bebe4fd0a77e5a4c882e3059",
        "region": "ru-3",
        "status": "PENDING_CREATE",
        "subnet_id": "",
        "updated_at": null,
        "zonal": false,
        "private_kube_api": true
    }
}
`

// expectedCreatePrivateKubeAPIClusterResponse represents an unmarshalled testCreatePrivateKubeAPIClusterResponseRaw.
var expectedCreatePrivateKubeAPIClusterResponse = &cluster.View{
	ID:                            "effe751d-501a-4b06-8e23-3f686dbfccf6",
	CreatedAt:                     &clusterResponseTimestamp,
	UpdatedAt:                     nil,
	Name:                          "test-private-cluster-0",
	Status:                        "PENDING_CREATE",
	ProjectID:                     "69744a03bebe4fd0a77e5a4c882e3059",
	NetworkID:                     "",
	SubnetID:                      "",
	KubeAPIIP:                     "",
	KubeVersion:                   "1.16.9",
	Region:                        "ru-3",
	AdditionalSoftware:            nil,
	PKITreeUpdatedAt:              nil,
	MaintenanceWindowStart:        "01:00:00",
	MaintenanceWindowEnd:          "03:00:00",
	MaintenanceLastStart:          &clusterResponseTimestamp,
	EnableAutorepair:              true,
	EnablePatchVersionAutoUpgrade: false,
	Zonal:                         false,
	PrivateKubeAPI:                true,
}

// testManyClustersInvalidResponseRaw represents a raw invalid response with several clusters.
const testManyClustersInvalidResponseRaw = `
{
    "clusters": [
        {
            "id": "dbe7559b-55d8-4f65-9230-6a22b985ff73",
        }
    ]
}
`

// testSingleClusterInvalidResponseRaw represents a raw invalid response with a single cluster.
const testSingleClusterInvalidResponseRaw = `
{
    "cluster": {
        "id": "dbe7559b-55d8-4f65-9230-6a22b985ff74",
    }
}
`

// testErrGenericResponseRaw represents a raw response with an error in the generic format.
const testErrGenericResponseRaw = `{"error":{"message":"bad gateway"}}`

// testGetKubeconfig represents a raw response from the GetKubeconfig request.
const testGetKubeconfig = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=
    server: https://203.0.113.101:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: admin
  name: admin@kubernetes
current-context: admin@kubernetes
kind: Config
preferences: {}
users:
- name: admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=
    client-key-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=
`

// testGetKubeconfigInvalidServer represents a raw response from the GetKubeconfig request with invalid server field.
const testGetKubeconfigInvalidServer = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=
    server:
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: admin
  name: admin@kubernetes
current-context: admin@kubernetes
kind: Config
preferences: {}
users:
- name: admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=
    client-key-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=
`

// testGetKubeconfigEmptyServer represents a raw response from the GetKubeconfig request without server field.
const testGetKubeconfigEmptyServer = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: admin
  name: admin@kubernetes
current-context: admin@kubernetes
kind: Config
preferences: {}
users:
- name: admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=
    client-key-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=
`

var testGetParsedKubeconfig = cluster.KubeconfigFields{
	ClusterCA:  "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
	Server:     "https://203.0.113.101:6443",
	ClientCert: "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
	ClientKey:  "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=",
	KubeconfigRaw: `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=
    server: https://203.0.113.101:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: admin
  name: admin@kubernetes
current-context: admin@kubernetes
kind: Config
preferences: {}
users:
- name: admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=
    client-key-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tS0tLQo=
`,
}

// testUpdateClusterOptsRaw represents marshalled options variant A for the Update request.
const testUpdateClusterOptsRaw = `
{
    "cluster": {
        "maintenance_window_start": "07:00:00",
        "enable_autorepair": true,
        "enable_patch_version_auto_upgrade": false
    }
}
`

// testUpdateClusterOpts represents options for the Update request.
var testUpdateClusterOpts = &cluster.UpdateOpts{
	MaintenanceWindowStart:        "07:00:00",
	EnableAutorepair:              testutils.BoolToPtr(true),
	EnablePatchVersionAutoUpgrade: testutils.BoolToPtr(false),
}

// testUpdateClusterWithEnabledPSPOptsRaw represents marshalled options variant A for the Update request.
const testUpdateClusterWithEnabledPSPOptsRaw = `
{
    "cluster": {
        "maintenance_window_start": "07:00:00",
        "enable_autorepair": true,
        "enable_patch_version_auto_upgrade": false,
        "kubernetes_options": {
            "enable_pod_security_policy": true,
            "feature_gates": [
                "CSIMigrationOpenStack"
            ],
            "admission_controllers": [
                "LimitRanger"
            ],
			"audit_logs": {
				"enabled": false,
				"secret_name": ""
			},
			"oidc": {
				"enabled": false,
				"provider_name": "",
				"client_id": "",
				"groups_claim": "",
				"issuer_url": "",
				"username_claim": ""
			}
        }
    }
}
`

// testUpdateClusterWithEnabledPSPOpts represents options for the Update request.
var testUpdateClusterWithEnabledPSPOpts = &cluster.UpdateOpts{
	MaintenanceWindowStart:        "07:00:00",
	EnableAutorepair:              testutils.BoolToPtr(true),
	EnablePatchVersionAutoUpgrade: testutils.BoolToPtr(false),
	KubernetesOptions: &cluster.KubernetesOptions{
		EnablePodSecurityPolicy: true,
		FeatureGates:            []string{"CSIMigrationOpenStack"},
		AdmissionControllers:    []string{"LimitRanger"},
	},
}

// testUpdateClusterWithEnabledPSPResponseRaw represents a raw response from the Create cluster request.
const testUpdateClusterWithEnabledPSPResponseRaw = `
{
    "cluster": {
        "additional_software": null,
        "created_at": "2020-02-13T09:18:32.05753Z",
        "enable_autorepair": true,
        "enable_patch_version_auto_upgrade": false,
        "id": "effe751d-501a-4b06-8e23-3f686dbfccf6",
        "kube_api_ip": "",
        "kube_version": "1.16.9",
        "maintenance_last_start": "2020-02-13T09:18:32.05753Z",
        "maintenance_window_end": "03:00:00",
        "maintenance_window_start": "01:00:00",
        "name": "test-zonal-cluster-0",
        "network_id": "",
        "pki_tree_updated_at": null,
        "project_id": "69744a03bebe4fd0a77e5a4c882e3059",
        "region": "ru-3",
        "status": "PENDING_UPGRADE_MASTERS_CONFIGURATION",
        "subnet_id": "",
        "updated_at": null,
        "zonal": false,
        "kubernetes_options": {
            "enable_pod_security_policy": true,
            "feature_gates": [
                "CSIMigrationOpenStack"
            ],
            "admission_controllers": [
                "LimitRanger"
            ] 
        }
    }
}
`

// expectedUpdateClusterWithEnabledPSPResponse represents an unmarshalled testUpdateClusterWithEnabledPSPResponseRaw.
var expectedUpdateClusterWithEnabledPSPResponse = &cluster.View{
	ID:                            "effe751d-501a-4b06-8e23-3f686dbfccf6",
	CreatedAt:                     &clusterResponseTimestamp,
	UpdatedAt:                     nil,
	Name:                          "test-zonal-cluster-0",
	Status:                        cluster.StatusPendingUpgradeMastersConfiguration,
	ProjectID:                     "69744a03bebe4fd0a77e5a4c882e3059",
	NetworkID:                     "",
	SubnetID:                      "",
	KubeAPIIP:                     "",
	KubeVersion:                   "1.16.9",
	Region:                        "ru-3",
	AdditionalSoftware:            nil,
	PKITreeUpdatedAt:              nil,
	MaintenanceWindowStart:        "01:00:00",
	MaintenanceWindowEnd:          "03:00:00",
	MaintenanceLastStart:          &clusterResponseTimestamp,
	EnableAutorepair:              true,
	EnablePatchVersionAutoUpgrade: false,
	Zonal:                         false,
	KubernetesOptions: &cluster.KubernetesOptions{
		EnablePodSecurityPolicy: true,
		FeatureGates:            []string{"CSIMigrationOpenStack"},
		AdmissionControllers:    []string{"LimitRanger"},
	},
}
