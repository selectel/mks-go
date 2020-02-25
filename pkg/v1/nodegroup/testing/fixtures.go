package testing

import (
	"time"

	"github.com/selectel/mks-go/pkg/v1/node"
	"github.com/selectel/mks-go/pkg/v1/nodegroup"
)

// testGetNodegroupResponseRaw represents a raw response from the Get request.
const testGetNodegroupResponseRaw = `
{
    "nodegroup": {
        "availability_zone": "ru-1a",
        "cluster_id": "79265515-3700-49fa-af0e-7f547bce788a",
        "created_at": "2020-02-19T15:41:45.948646Z",
        "flavor_id": "99b62670-9d78-43fd-8f55-d184a4800f8d",
        "id": "a376745a-fbcb-413d-b418-169d059d79ce",
        "local_volume": false,
        "nodes": [
            {
                "created_at": "2020-02-19T15:41:45.948646Z",
                "hostname": "test-cluster-node-eegp9",
                "id": "39e5dd4d-5e23-4a00-8173-974bf844f21b",
                "ip": "198.51.100.11",
                "nodegroup_id": "a376745a-fbcb-413d-b418-169d059d79ce",
                "updated_at": "2020-02-19T15:41:45.948646Z"
            }
        ],
        "updated_at": "2020-02-19T15:41:45.948646Z",
        "volume_gb": 10,
        "volume_type": "basic.ru-1a"
    }
}
`

var nodegroupResponseTimestamp, _ = time.Parse(time.RFC3339, "2020-02-19T15:41:45.948646Z")

// nolint
// expectedGetNodegroupResponse represents an unmarshalled testGetNodegroupResponseRaw.
var expectedGetNodegroupResponse = &nodegroup.View{
	ID:               "a376745a-fbcb-413d-b418-169d059d79ce",
	CreatedAt:        &nodegroupResponseTimestamp,
	UpdatedAt:        &nodegroupResponseTimestamp,
	ClusterID:        "79265515-3700-49fa-af0e-7f547bce788a",
	FlavorID:         "99b62670-9d78-43fd-8f55-d184a4800f8d",
	VolumeGB:         10,
	VolumeType:       "basic.ru-1a",
	LocalVolume:      false,
	AvailabilityZone: "ru-1a",
	Nodes: []*node.View{
		{
			ID:          "39e5dd4d-5e23-4a00-8173-974bf844f21b",
			CreatedAt:   &nodegroupResponseTimestamp,
			UpdatedAt:   &nodegroupResponseTimestamp,
			Hostname:    "test-cluster-node-eegp9",
			IP:          "198.51.100.11",
			NodegroupID: "a376745a-fbcb-413d-b418-169d059d79ce",
		},
	},
}

// testListNodegroupsResponseRaw represents a raw response from the List method.
const testListNodegroupsResponseRaw = `
{
    "nodegroups": [
        {
            "availability_zone": "ru-1a",
            "cluster_id": "79265515-3700-49fa-af0e-7f547bce788a",
            "created_at": "2020-02-19T15:41:45.948646Z",
            "flavor_id": "99b62670-9d78-43fd-8f55-d184a4800f8d",
            "id": "a376745a-fbcb-413d-b418-169d059d79ce",
            "local_volume": false,
            "nodes": [
                {
                    "created_at": "2020-02-19T15:41:45.948646Z",
                    "hostname": "test-cluster-node-eegp9",
                    "id": "39e5dd4d-5e23-4a00-8173-974bf844f21b",
                    "ip": "198.51.100.11",
                    "nodegroup_id": "a376745a-fbcb-413d-b418-169d059d79ce",
                    "updated_at": "2020-02-19T15:41:45.948646Z"
                }
            ],
            "updated_at": "2020-02-19T15:41:45.948646Z",
            "volume_gb": 10,
            "volume_type": "basic.ru-1a"
        }
    ]
}
`

// nolint
// expectedListNodegroupsResponse represents an unmarshalled testListNodegroupsResponseRaw.
var expectedListNodegroupsResponse = []*nodegroup.View{
	{
		ID:               "a376745a-fbcb-413d-b418-169d059d79ce",
		CreatedAt:        &nodegroupResponseTimestamp,
		UpdatedAt:        &nodegroupResponseTimestamp,
		ClusterID:        "79265515-3700-49fa-af0e-7f547bce788a",
		FlavorID:         "99b62670-9d78-43fd-8f55-d184a4800f8d",
		VolumeGB:         10,
		VolumeType:       "basic.ru-1a",
		LocalVolume:      false,
		AvailabilityZone: "ru-1a",
		Nodes: []*node.View{
			{
				ID:          "39e5dd4d-5e23-4a00-8173-974bf844f21b",
				CreatedAt:   &nodegroupResponseTimestamp,
				UpdatedAt:   &nodegroupResponseTimestamp,
				Hostname:    "test-cluster-node-eegp9",
				IP:          "198.51.100.11",
				NodegroupID: "a376745a-fbcb-413d-b418-169d059d79ce",
			},
		},
	},
}

// testCreateNodegroupOptsRaw represents marshalled options for the Create request.
const testCreateNodegroupOptsRaw = `
{
    "nodegroup": {
        "count": 1,
        "cpus": 1,
        "ram_mb": 2048,
        "volume_gb": 10,
        "volume_type": "fast.ru-1b",
        "keypair_name": "ssh-key",
        "availability_zone": "ru-1b"
    }
}
`

// nolint
// testCreateNodegroupOpts represents options for the Create request.
var testCreateNodegroupOpts = &nodegroup.CreateOpts{
	Count:            1,
	CPUs:             1,
	RAMMB:            2048,
	VolumeGB:         10,
	VolumeType:       "fast.ru-1b",
	KeypairName:      "ssh-key",
	AvailabilityZone: "ru-1b",
}

// testResizeNodegroupOptsRaw represents marshalled options for the Resize request.
const testResizeNodegroupOptsRaw = `
{
    "nodegroup": {
        "desired": 1
    }
}
`

// nolint
// testResizeNodegroupOpts represents options for the Resize request.
var testResizeNodegroupOpts = &nodegroup.ResizeOpts{
	Desired: 1,
}

// testManyNodegroupsInvalidResponseRaw represents a raw invalid response with several nodegroups.
const testManyNodegroupsInvalidResponseRaw = `
{
    "nodegroups": [
        {
            "id": "a376745a-fbcb-413d-b418-169d059d79ce",
        }
    ]
}
`

// testSingleNodegroupInvalidResponseRaw represents a raw invalid response with a single nodegroup.
const testSingleNodegroupInvalidResponseRaw = `
{
    "nodegroup": {
        "id": "a376745a-fbcb-413d-b418-169d059d79ce",
    }
}
`

// testErrGenericResponseRaw represents a raw response with an error in the generic format.
const testErrGenericResponseRaw = `{"error":{"message":"bad gateway"}}`
