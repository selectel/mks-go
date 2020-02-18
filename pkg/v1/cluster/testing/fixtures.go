package testing

import (
	"time"

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
        "updated_at": "2020-02-13T09:18:32.05753Z"
    }
}
`

var clusterResponseTimestamp, _ = time.Parse(time.RFC3339, "2020-02-13T09:18:32.05753Z")

// expectedGetClusterResponse represents an unmarshalled testGetClusterResponseRaw.
var expectedGetClusterResponse = &cluster.View{
	ID:                     "dbe7559b-55d8-4f65-9230-6a22b985ff73",
	CreatedAt:              &clusterResponseTimestamp,
	UpdatedAt:              &clusterResponseTimestamp,
	Name:                   "test-cluster",
	Status:                 "ACTIVE",
	ProjectID:              "65044a03bede4fd0a77e5a4c882e3059",
	NetworkID:              "74a591be-6be7-4abc-d30f-1614c0f9721c",
	SubnetID:               "c872541d-2d83-419f-841d-8288201b8fb9",
	KubeAPIIP:              "203.0.113.101",
	KubeVersion:            "1.15.7",
	Region:                 "ru-1",
	AdditionalSoftware:     nil,
	PKITreeUpdatedAt:       &clusterResponseTimestamp,
	MaintenanceWindowStart: "01:00:00",
	MaintenanceWindowEnd:   "03:00:00",
	MaintenanceLastStart:   &clusterResponseTimestamp,
	EnableAutorepair:       true,
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
            "updated_at": "2020-02-13T09:18:32.05753Z"
        }
    ]
}
`

// expectedListClustersResponse represents an unmarshalled testListClustersResponseRaw.
var expectedListClustersResponse = []*cluster.View{
	{
		ID:                     "dbe7559b-55d8-4f65-9230-6a22b985ff73",
		CreatedAt:              &clusterResponseTimestamp,
		UpdatedAt:              &clusterResponseTimestamp,
		Name:                   "test-cluster",
		Status:                 "ACTIVE",
		ProjectID:              "65044a03bede4fd0a77e5a4c882e3059",
		NetworkID:              "74a591be-6be7-4abc-d30f-1614c0f9721c",
		SubnetID:               "c872541d-2d83-419f-841d-8288201b8fb9",
		KubeAPIIP:              "203.0.113.101",
		KubeVersion:            "1.15.7",
		Region:                 "ru-1",
		AdditionalSoftware:     nil,
		PKITreeUpdatedAt:       nil,
		MaintenanceWindowStart: "",
		MaintenanceWindowEnd:   "",
		MaintenanceLastStart:   nil,
		EnableAutorepair:       true,
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
                "availability_zone": "ru-1b"
            }
        ]
    }
}
`

// nolint
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
        "updated_at": null
    }
}
`

// expectedCreateClusterResponse represents an unmarshalled testCreateClusterResponseRaw.
var expectedCreateClusterResponse = &cluster.View{
	ID:                     "1f4dfa6a-0468-45e0-be13-74c6481820f5",
	CreatedAt:              &clusterResponseTimestamp,
	UpdatedAt:              nil,
	Name:                   "test-cluster-0",
	Status:                 "PENDING_CREATE",
	ProjectID:              "69744a03bebe4fd0a77e5a4c882e3059",
	NetworkID:              "",
	SubnetID:               "",
	KubeAPIIP:              "",
	KubeVersion:            "1.15.7",
	Region:                 "ru-1",
	AdditionalSoftware:     nil,
	PKITreeUpdatedAt:       nil,
	MaintenanceWindowStart: "01:00:00",
	MaintenanceWindowEnd:   "03:00:00",
	MaintenanceLastStart:   &clusterResponseTimestamp,
	EnableAutorepair:       true,
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
