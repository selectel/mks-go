package task

import (
	"context"

	"github.com/selectel/mks-go/v2/internal/common"
	mks "github.com/selectel/mks-go/v2/pkg"
	"github.com/selectel/mks-go/v2/pkg/mksclient"
)

// Get returns a cluster task by its id.
func Get(ctx context.Context, client *mks.ServiceClient, clusterID, taskID string) (*mksclient.Task, error) {
	responseResult, err := client.MKSClient.GetTaskV2WithResponse(
		ctx, clusterID,
		taskID, &mksclient.GetTaskV2Params{WithErrorDetails: common.Ptr(true)},
	)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		return &responseResult.JSON200.Task, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON404, responseResult.JSON500,
	)
}

// List gets a list of all cluster tasks.
func List(ctx context.Context, client *mks.ServiceClient, clusterID string, limit, offset uint64) ([]mksclient.Task, error) {
	responseResult, err := client.MKSClient.ListTasksV2WithResponse(
		ctx, clusterID,
		&mksclient.ListTasksV2Params{
			Limit: &limit, Offset: &offset, WithErrorDetails: common.Ptr(true),
		},
	)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		if responseResult.JSON200.Tasks == nil {
			return []mksclient.Task{}, nil
		}

		return responseResult.JSON200.Tasks, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON404, responseResult.JSON500,
	)
}
