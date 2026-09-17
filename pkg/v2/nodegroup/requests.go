package nodegroup

import (
	"context"
	"net/http"

	v2 "github.com/selectel/mks-go/pkg/v2"
	"github.com/selectel/mks-go/pkg/v2/mksclient"
)

// Get returns a cluster nodegroup by its id.
func Get(ctx context.Context, client *v2.ServiceClient, clusterID, nodegroupID string) (*mksclient.NodegroupDetailed, error) {
	responseResult, err := client.MKSClient.GetNodegroupV2WithResponse(ctx, clusterID, nodegroupID)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		return &responseResult.JSON200.Nodegroup, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404, responseResult.JSON500,
	)
}

// List gets a list of all cluster nodegroups.
func List(ctx context.Context, client *v2.ServiceClient, clusterID string) ([]mksclient.NodegroupListItem, error) {
	responseResult, err := client.MKSClient.ListNodegroupsV2WithResponse(ctx, clusterID)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		if responseResult.JSON200.Nodegroups == nil {
			return []mksclient.NodegroupListItem{}, nil
		}

		return responseResult.JSON200.Nodegroups, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404, responseResult.JSON500,
	)
}

// Create requests a creation of a new cluster nodegroup.
func Create(ctx context.Context, client *v2.ServiceClient, clusterID string, nodegroups []mksclient.NodegroupCreateStruct) error {
	responseResult, err := client.MKSClient.CreateNodegroupsV2WithResponse(
		ctx, clusterID,
		mksclient.NodegroupsCreateBody{Nodegroups: nodegroups},
	)
	if err != nil {
		return err
	}

	if responseResult.StatusCode() == http.StatusNoContent {
		return nil
	}

	return mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404, responseResult.JSON409, responseResult.JSON500,
	)
}

// Delete deletes a cluster nodegroup by its id.
func Delete(ctx context.Context, client *v2.ServiceClient, clusterID, nodegroupID string) error {
	responseResult, err := client.MKSClient.DeleteNodegroupV2WithResponse(ctx, clusterID, nodegroupID)
	if err != nil {
		return err
	}

	if responseResult.StatusCode() == http.StatusNoContent {
		return nil
	}

	return mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404, responseResult.JSON409, responseResult.JSON500,
	)
}

// Resize requests a resize of a cluster nodegroup by its id.
func Resize(ctx context.Context, client *v2.ServiceClient, clusterID, nodegroupID string, desired int64) error {
	responseResult, err := client.MKSClient.ResizeNodegroupV2WithResponse(
		ctx, clusterID, nodegroupID,
		mksclient.NodegroupResizeBody{
			Nodegroup: mksclient.NodegroupResizeStruct{Desired: desired},
		},
	)
	if err != nil {
		return err
	}

	if responseResult.StatusCode() == http.StatusNoContent {
		return nil
	}

	return mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404, responseResult.JSON409, responseResult.JSON500,
	)
}

// Update requests an update of a cluster nodegroup by its id.
func Update(ctx context.Context, client *v2.ServiceClient, clusterID, nodegroupID string, nodegroup mksclient.NodegroupUpdateStruct) error {
	responseResult, err := client.MKSClient.UpdateNodegroupV2WithResponse(
		ctx, clusterID, nodegroupID,
		mksclient.NodegroupUpdateBody{Nodegroup: nodegroup},
	)
	if err != nil {
		return err
	}

	if responseResult.StatusCode() == http.StatusNoContent {
		return nil
	}

	return mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404, responseResult.JSON409, responseResult.JSON500,
	)
}
