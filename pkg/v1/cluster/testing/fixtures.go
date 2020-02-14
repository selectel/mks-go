package testing

import (
	"time"

	"github.com/selectel/mks-go/pkg/v1/cluster"
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

// testListClustersInvalidResponseRaw represents a raw invalid response with several clusters.
const testListClustersInvalidResponseRaw = `
{
    "clusters": [
        {
            "id": "dbe7559b-55d8-4f65-9230-6a22b985ff73",
        }
    ]
}
`

// testGetClusterInvalidResponseRaw represents a raw invalid response with a single cluster.
const testGetClusterInvalidResponseRaw = `
{
    "cluster": {
        "id": "dbe7559b-55d8-4f65-9230-6a22b985ff74",
    }
}
`
