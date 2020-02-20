package testing

import (
	"time"

	"github.com/selectel/mks-go/pkg/v1/task"
)

// testGetTaskResponseRaw represents a raw response from the Get request.
const testGetTaskResponseRaw = `
{
    "task": {
        "cluster_id": "d2e16a48-a9c5-4449-8b71-71f21fc872db",
        "id": "2f6fb93c-cf0d-4289-a78c-34393ac75f92",
        "started_at": "2020-02-19T11:43:02.868387Z",
        "status": "DONE",
        "type": "CREATE_CLUSTER",
        "updated_at": "2020-02-19T11:43:02.868387Z"
    }
}
`

var taskResponseTimestamp, _ = time.Parse(time.RFC3339, "2020-02-19T11:43:02.868387Z")

// expectedGetTaskResponse represents an unmarshalled testGetTaskResponseRaw.
var expectedGetTaskResponse = &task.View{
	ID:        "2f6fb93c-cf0d-4289-a78c-34393ac75f92",
	StartedAt: &taskResponseTimestamp,
	UpdatedAt: &taskResponseTimestamp,
	ClusterID: "d2e16a48-a9c5-4449-8b71-71f21fc872db",
	Status:    task.StatusDone,
	Type:      task.TypeCreateCluster,
}

// testGetTaskUnknownStatusAndTypeResponseRaw represents a raw response from the Get request
// with unknown status and type.
const testGetTaskUnknownStatusAndTypeResponseRaw = `
{
    "task": {
        "cluster_id": "d2e16a48-a9c5-4449-8b71-71f21fc872dc",
        "id": "2f6fb93c-cf0d-4289-a78c-34393ac75f92",
        "started_at": "2020-02-19T11:43:02.868387Z",
        "status": "FAKE_STATUS",
        "type": "FAKE_TYPE",
        "updated_at": "2020-02-19T11:43:02.868387Z"
    }
}
`

// expectedGetTaskUnknownStatusAndTypeResponse represents an unmarshalled
// testGetTaskUnknownStatusAndTypeResponseRaw.
var expectedGetTaskUnknownStatusAndTypeResponse = &task.View{
	ID:        "2f6fb93c-cf0d-4289-a78c-34393ac75f92",
	StartedAt: &taskResponseTimestamp,
	UpdatedAt: &taskResponseTimestamp,
	ClusterID: "d2e16a48-a9c5-4449-8b71-71f21fc872dc",
	Status:    task.StatusUnknown,
	Type:      task.TypeUnknown,
}

// testListTasksResponseRaw represents a raw response from the List method.
const testListTasksResponseRaw = `
{
    "tasks": [
        {
            "cluster_id": "d2e16a48-a9c5-4449-9b71-81f21fc872db",
            "id": "2f8fb93c-cf9e-4289-a78c-34393ac75f92",
            "started_at": "2020-02-19T11:43:02.868387Z",
            "status": "DONE",
            "type": "CREATE_CLUSTER",
            "updated_at": "2020-02-19T11:43:02.868387Z"
        },
        {
            "cluster_id": "d2e16a48-a9c5-4449-9b71-81f21fc872db",
            "id": "3f8fb93c-cf9e-4289-a78c-34393ac75f92",
            "started_at": "2020-02-19T11:43:02.868387Z",
            "status": "DONE",
            "type": "ROTATE_CERTS",
            "updated_at": "2020-02-19T11:43:02.868387Z"
        },
        {
            "cluster_id": "d2e16a48-a9c5-4449-9b71-81f21fc872db",
            "id": "4f8fb93c-cf9e-4289-a78c-34393ac75f92",
            "started_at": "2020-02-19T11:43:02.868387Z",
            "status": "DONE",
            "type": "NODE_GROUP_RESIZE",
            "updated_at": "2020-02-19T11:43:02.868387Z"
        },
        {
            "cluster_id": "d2e16a48-a9c5-4449-9b71-81f21fc872db",
            "id": "5f8fb93c-cf9e-4289-a78c-34393ac75f92",
            "started_at": "2020-02-19T11:43:02.868387Z",
            "status": "DONE",
            "type": "NODE_REINSTALL",
            "updated_at": "2020-02-19T11:43:02.868387Z"
        },
        {
            "cluster_id": "d2e16a48-a9c5-4449-9b71-81f21fc872db",
            "id": "6f8fb93c-cf9e-4289-a78c-34393ac75f92",
            "started_at": "2020-02-19T11:43:02.868387Z",
            "status": "DONE",
            "type": "CLUSTER_RESIZE",
            "updated_at": "2020-02-19T11:43:02.868387Z"
        },
        {
            "cluster_id": "d2e16a48-a9c5-4449-9b71-81f21fc872db",
            "id": "7f8fb93c-cf9e-4289-a78c-34393ac75f92",
            "started_at": "2020-02-19T11:43:02.868387Z",
            "status": "ERROR",
            "type": "UPGRADE_PATCH_VERSION",
            "updated_at": "2020-02-19T11:43:02.868387Z"
        },
        {
            "cluster_id": "d2e16a48-a9c5-4449-9b71-81f21fc872db",
            "id": "8f8fb93c-cf9e-4289-a78c-34393ac75f92",
            "started_at": "2020-02-19T11:43:02.868387Z",
            "status": "IN_PROGRESS",
            "type": "DELETE_CLUSTER",
            "updated_at": "2020-02-19T11:43:02.868387Z"
        }
    ]
}
`

// expectedListTasksResponse represents an unmarshalled testListTasksResponseRaw.
var expectedListTasksResponse = []*task.View{
	{
		ID:        "2f8fb93c-cf9e-4289-a78c-34393ac75f92",
		StartedAt: &taskResponseTimestamp,
		UpdatedAt: &taskResponseTimestamp,
		ClusterID: "d2e16a48-a9c5-4449-9b71-81f21fc872db",
		Status:    task.StatusDone,
		Type:      task.TypeCreateCluster,
	},
	{
		ID:        "3f8fb93c-cf9e-4289-a78c-34393ac75f92",
		StartedAt: &taskResponseTimestamp,
		UpdatedAt: &taskResponseTimestamp,
		ClusterID: "d2e16a48-a9c5-4449-9b71-81f21fc872db",
		Status:    task.StatusDone,
		Type:      task.TypeRotateCerts,
	},
	{
		ID:        "4f8fb93c-cf9e-4289-a78c-34393ac75f92",
		StartedAt: &taskResponseTimestamp,
		UpdatedAt: &taskResponseTimestamp,
		ClusterID: "d2e16a48-a9c5-4449-9b71-81f21fc872db",
		Status:    task.StatusDone,
		Type:      task.TypeNodeGroupResize,
	},
	{
		ID:        "5f8fb93c-cf9e-4289-a78c-34393ac75f92",
		StartedAt: &taskResponseTimestamp,
		UpdatedAt: &taskResponseTimestamp,
		ClusterID: "d2e16a48-a9c5-4449-9b71-81f21fc872db",
		Status:    task.StatusDone,
		Type:      task.TypeNodeReinstall,
	},
	{
		ID:        "6f8fb93c-cf9e-4289-a78c-34393ac75f92",
		StartedAt: &taskResponseTimestamp,
		UpdatedAt: &taskResponseTimestamp,
		ClusterID: "d2e16a48-a9c5-4449-9b71-81f21fc872db",
		Status:    task.StatusDone,
		Type:      task.TypeClusterResize,
	},
	{
		ID:        "7f8fb93c-cf9e-4289-a78c-34393ac75f92",
		StartedAt: &taskResponseTimestamp,
		UpdatedAt: &taskResponseTimestamp,
		ClusterID: "d2e16a48-a9c5-4449-9b71-81f21fc872db",
		Status:    task.StatusError,
		Type:      task.TypeUpgradePatchVersion,
	},
	{
		ID:        "8f8fb93c-cf9e-4289-a78c-34393ac75f92",
		StartedAt: &taskResponseTimestamp,
		UpdatedAt: &taskResponseTimestamp,
		ClusterID: "d2e16a48-a9c5-4449-9b71-81f21fc872db",
		Status:    task.StatusInProgress,
		Type:      task.TypeDeleteCluster,
	},
}

// testListTasksUnknownStatusAndTypeResponseRaw represents a raw response from the List method
// with unknown status and type.
const testListTasksUnknownStatusAndTypeResponseRaw = `
{
    "tasks": [
        {
            "cluster_id": "d1e16a48-a9c5-4449-9b71-81f21fc872cb",
            "id": "2f9fb93c-cf9e-3289-a78c-34393ac75f92",
            "started_at": "2020-02-19T11:43:02.868387Z",
            "status": "FAKE_STATUS",
            "type": "FAKE_TYPE",
            "updated_at": "2020-02-19T11:43:02.868387Z"
        }
    ]
}
`

// expectedListTasksUnknownStatusAndTypeResponse represents an unmarshalled
// testListTasksUnknownStatusAndTypeResponseRaw.
var expectedListTasksUnknownStatusAndTypeResponse = []*task.View{
	{
		ID:        "2f9fb93c-cf9e-3289-a78c-34393ac75f92",
		StartedAt: &taskResponseTimestamp,
		UpdatedAt: &taskResponseTimestamp,
		ClusterID: "d1e16a48-a9c5-4449-9b71-81f21fc872cb",
		Status:    task.StatusUnknown,
		Type:      task.TypeUnknown,
	},
}

// testManyTasksInvalidResponse represents a raw invalid response with several tasks.
const testManyTasksInvalidResponse = `
{
    "tasks": [
        {
            "id": "2f8fb93c-cf9e-4289-a78c-34393ac75f92",
        }
    ]
}
`

// testSingleTaskInvalidResponse represents a raw invalid response with a single task.
const testSingleTaskInvalidResponseRaw = `
{
    "task": {
        "id": "2fbfb93c-cf0d-4289-a78c-34393ac75f92",
    }
}
`

// testErrGenericResponseRaw represents a raw response with an error in the generic format.
const testErrGenericResponseRaw = `{"error":{"message":"bad gateway"}}`
