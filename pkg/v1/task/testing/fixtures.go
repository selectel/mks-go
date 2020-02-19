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
	Status:    "DONE",
	Type:      "CREATE_CLUSTER",
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
		Status:    "DONE",
		Type:      "CREATE_CLUSTER",
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
