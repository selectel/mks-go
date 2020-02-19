package task

import (
	"context"
	"net/http"
	"strings"

	v1 "github.com/selectel/mks-go/pkg/v1"
)

// Get returns a cluster task by its id.
func Get(ctx context.Context, client *v1.ServiceClient, clusterID, taskID string) (*View, *v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster, clusterID, v1.ResourceURLTask, taskID}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a task from the response body.
	var result struct {
		Task *View `json:"task"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Task, responseResult, err
}

// List gets a list of all cluster tasks.
func List(ctx context.Context, client *v1.ServiceClient, clusterID string) ([]*View, *v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster, clusterID, v1.ResourceURLTask}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract tasks from the response body.
	var result struct {
		Tasks []*View `json:"tasks"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Tasks, responseResult, err
}
