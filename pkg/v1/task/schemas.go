package task

import "time"

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
	Status string `json:"status"`

	// Task represents task's type.
	Type string `json:"type"`
}
