package cplogs

import (
	"context"
	"net/http"

	mks "github.com/selectel/mks-go/v2/pkg"
	"github.com/selectel/mks-go/v2/pkg/mksclient"
)

// Get returns control plane logs collection settings for provided cluster.
func Get(ctx context.Context, client *mks.ServiceClient, clusterID string) (*mksclient.ControlPlaneLogging, error) {
	responseResult, err := client.MKSClient.GetControlPlaneLoggingV2WithResponse(ctx, clusterID)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		return responseResult.JSON200, nil
	}

	return nil, mksclient.HandleAPIErrorsWithBody(
		responseResult.StatusCode(), responseResult.Status(), responseResult.Body,
		responseResult.JSON403, responseResult.JSON404, responseResult.JSON500,
	)
}

// Init initializes control plane logs collection for provided cluster.
func Init(
	ctx context.Context, client *mks.ServiceClient,
	clusterID string, opts *mksclient.ControlPlaneLogging,
) (*mksclient.ControlPlaneLogging, error) {
	responseResult, err := client.MKSClient.InitControlPlaneLoggingV2WithResponse(ctx, clusterID, *opts)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON201 != nil {
		return responseResult.JSON201, nil
	}

	return nil, mksclient.HandleAPIErrorsWithBody(
		responseResult.StatusCode(), responseResult.Status(), responseResult.Body,
		responseResult.JSON403, responseResult.JSON404, responseResult.JSON409, responseResult.JSON500,
	)
}

// Update updates control plane logs collection settings for provided cluster.
func Update(
	ctx context.Context, client *mks.ServiceClient,
	clusterID string, opts *mksclient.ControlPlaneLoggingUpdateBody,
) error {
	responseResult, err := client.MKSClient.UpdateControlPlaneLoggingV2WithResponse(ctx, clusterID, *opts)
	if err != nil {
		return err
	}

	if responseResult.StatusCode() == http.StatusNoContent {
		return nil
	}

	return mksclient.HandleAPIErrorsWithBody(
		responseResult.StatusCode(), responseResult.Status(), responseResult.Body,
		responseResult.JSON403, responseResult.JSON404, responseResult.JSON409, responseResult.JSON500,
	)
}

// ListComponents returns all available components for control plane logs collection.
func ListComponents(ctx context.Context, client *mks.ServiceClient) ([]mksclient.LogComponent, error) {
	responseResult, err := client.MKSClient.ListLogComponentsV2WithResponse(ctx, &mksclient.ListLogComponentsV2Params{})
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		if responseResult.JSON200.Components == nil {
			return []mksclient.LogComponent{}, nil
		}

		return *responseResult.JSON200.Components, nil
	}

	return nil, mksclient.HandleAPIErrorsWithBody(
		responseResult.StatusCode(), responseResult.Status(), responseResult.Body,
		responseResult.JSON403, responseResult.JSON500,
	)
}
