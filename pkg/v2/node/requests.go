package node

import (
	"context"
	"net/http"

	v2 "github.com/selectel/mks-go/pkg/v2"
	"github.com/selectel/mks-go/pkg/v2/mksclient"
)

// Get returns a node of a cluster nodegroup by its id.
func Get(ctx context.Context, client *v2.ServiceClient, clusterID, nodegroupID, nodeID string) (*mksclient.Node, error) {
	responseResult, err := client.MKSClient.GetNodeV2WithResponse(ctx, clusterID, nodegroupID, nodeID)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		return &responseResult.JSON200.Node, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON404, responseResult.JSON500,
	)
}

// Reinstall requests to make reinstall of a single node of a cluster nodegroup by its id.
func Reinstall(ctx context.Context, client *v2.ServiceClient, clusterID, nodegroupID, nodeID string) error {
	responseResult, err := client.MKSClient.ReinstallNodeV2WithResponse(ctx, clusterID, nodegroupID, nodeID)
	if err != nil {
		return err
	}

	if responseResult.StatusCode() == http.StatusNoContent {
		return nil
	}

	return mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404,
		responseResult.JSON409, responseResult.JSON500,
	)
}

// Delete deletes a node of a cluster nodegroup by its id.
func Delete(ctx context.Context, client *v2.ServiceClient, clusterID, nodegroupID, nodeID string) error {
	responseResult, err := client.MKSClient.DeleteNodeV2WithResponse(ctx, clusterID, nodegroupID, nodeID)
	if err != nil {
		return err
	}

	if responseResult.StatusCode() == http.StatusNoContent {
		return nil
	}

	return mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404,
		responseResult.JSON409, responseResult.JSON500,
	)
}
