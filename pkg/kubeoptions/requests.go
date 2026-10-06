package kubeoptions

import (
	"context"

	mks "github.com/selectel/mks-go/v2/pkg"
	"github.com/selectel/mks-go/v2/pkg/mksclient"
)

// ListFeatureGates gets a list of available feature gates by Kubernetes versions.
func ListFeatureGates(ctx context.Context, client *mks.ServiceClient) ([]mksclient.AvailableFeatureGates, error) {
	responseResult, err := client.MKSClient.ListFeatureGatesV2WithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		if responseResult.JSON200.FeatureGates == nil {
			return []mksclient.AvailableFeatureGates{}, nil
		}

		return *responseResult.JSON200.FeatureGates, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON500,
	)
}

// ListAdmissionControllers gets a list of available admission controllers by Kubernetes versions.
func ListAdmissionControllers(ctx context.Context, client *mks.ServiceClient) ([]mksclient.AvailableAdmissionControllers, error) {
	responseResult, err := client.MKSClient.ListAdmissionControllersV2WithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		if responseResult.JSON200.AdmissionControllers == nil {
			return []mksclient.AvailableAdmissionControllers{}, nil
		}

		return *responseResult.JSON200.AdmissionControllers, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON500,
	)
}
