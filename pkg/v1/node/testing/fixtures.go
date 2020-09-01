package testing

import (
	"time"

	"github.com/selectel/mks-go/pkg/v1/node"
)

// testGetNodeResponseRaw represents a raw response from the Get request.
const testGetNodeResponseRaw = `
{
    "node": {
        "created_at": "2020-02-25T08:49:21.780542Z",
        "hostname": "test-cluster-node-gap1g",
        "id": "203d0f8c-547d-48a7-98ed-3075254b8d4a",
        "ip": "198.51.100.11",
        "nodegroup_id": "f174b65d-442a-4423-aaf7-5654789b8a9d",
        "os_server_id": "dc56abe9-d0d4-4099-9b5f-e5cabfccf276",
        "updated_at": "2020-02-25T08:49:21.780542Z"
    }
}
`

var nodeResponseTimestamp, _ = time.Parse(time.RFC3339, "2020-02-25T08:49:21.780542Z")

// expectedGetNodeResponse represents an unmarshalled testGetNodeResponseRaw.
var expectedGetNodeResponse = &node.View{
	ID:          "203d0f8c-547d-48a7-98ed-3075254b8d4a",
	CreatedAt:   &nodeResponseTimestamp,
	UpdatedAt:   &nodeResponseTimestamp,
	Hostname:    "test-cluster-node-gap1g",
	IP:          "198.51.100.11",
	NodegroupID: "f174b65d-442a-4423-aaf7-5654789b8a9d",
	OSServerID:  "dc56abe9-d0d4-4099-9b5f-e5cabfccf276",
}

// testSingleNodeInvalidResponseRaw represents a raw invalid response with a single node.
const testSingleNodeInvalidResponseRaw = `
{
    "node": {
        "id": "203d0f8c-547d-48a7-98ed-3075254b8d4a",
    }
}
`

// testErrGenericResponseRaw represents a raw response with an error in the generic format.
const testErrGenericResponseRaw = `{"error":{"message":"bad gateway"}}`
