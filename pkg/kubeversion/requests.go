package kubeversion

import (
	"context"

	mks "github.com/selectel/mks-go/v2/pkg"
	"github.com/selectel/mks-go/v2/pkg/mksclient"
)

// List returns all supported Kubernetes versions.
func List(ctx context.Context, client *mks.ServiceClient) ([]mksclient.KubeVersionInfo, error) {
	responseResult, err := client.MKSClient.ListKubeVersionsV2WithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		if responseResult.JSON200.KubeVersions == nil {
			return []mksclient.KubeVersionInfo{}, nil
		}

		return *responseResult.JSON200.KubeVersions, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON500,
	)
}
