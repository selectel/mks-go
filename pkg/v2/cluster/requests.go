package cluster

import (
	"context"
	"net/http"

	v2 "github.com/selectel/mks-go/pkg/v2"
	"github.com/selectel/mks-go/pkg/v2/mksclient"
)

// Get returns a single cluster by its id.
func Get(ctx context.Context, client *v2.ServiceClient, clusterID string) (*mksclient.ClusterDetailed, error) {
	responseResult, err := client.MKSClient.GetClusterV2WithResponse(ctx, clusterID)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		return responseResult.JSON200.Cluster, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON404, responseResult.JSON500,
	)
}

// Create requests a creation of a new cluster.
func Create(ctx context.Context, client *v2.ServiceClient, opts *mksclient.ClusterCreateStruct) (*mksclient.ClusterDetailed, error) {
	responseResult, err := client.MKSClient.CreateClusterV2WithResponse(ctx, mksclient.ClusterCreateBody{Cluster: opts})
	if err != nil {
		return nil, err
	}

	if responseResult.JSON201 != nil {
		return responseResult.JSON201.Cluster, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON500,
	)
}

// Update requests an update of an existing cluster.
func Update(ctx context.Context, client *v2.ServiceClient, clusterID string, opts *mksclient.ClusterUpdateStruct) (*mksclient.ClusterDetailed, error) {
	responseResult, err := client.MKSClient.UpdateClusterV2WithResponse(ctx, clusterID, mksclient.UpdateClusterV2JSONRequestBody{Cluster: opts})
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		return responseResult.JSON200.Cluster, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON500,
	)
}

// Delete deletes a single cluster by its id.
func Delete(ctx context.Context, client *v2.ServiceClient, clusterID string) error {
	responseResult, err := client.MKSClient.DeleteClusterV2WithResponse(ctx, clusterID)
	if err != nil {
		return err
	}

	if responseResult.StatusCode() == http.StatusNoContent {
		return nil
	}

	return mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON404, responseResult.JSON500,
	)
}

// GetKubeconfig returns a kubeconfig by cluster id.
func GetKubeconfig(ctx context.Context, client *v2.ServiceClient, clusterID string) ([]byte, error) {
	responseResult, err := client.MKSClient.GetClusterKubeconfigV2WithResponse(ctx, clusterID)
	if err != nil {
		return nil, err
	}

	if responseResult.StatusCode() == http.StatusOK {
		return responseResult.Body, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON404, responseResult.JSON500,
	)
}

// RotateCerts requests a rotation of cluster certificates by cluster id.
func RotateCerts(ctx context.Context, client *v2.ServiceClient, clusterID string) error {
	responseResult, err := client.MKSClient.RotateClusterCertsV2WithResponse(ctx, clusterID)
	if err != nil {
		return err
	}

	if responseResult.StatusCode() == http.StatusNoContent {
		return nil
	}

	return mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON404, responseResult.JSON500,
	)
}

// UpgradePatchVersion requests a Kubernetes patch version upgrade by cluster id.
func UpgradePatchVersion(ctx context.Context, client *v2.ServiceClient, clusterID string) (*mksclient.ClusterDetailed, error) {
	responseResult, err := client.MKSClient.UpgradePatchVersionV2WithResponse(ctx, clusterID)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		return responseResult.JSON200.Cluster, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON404, responseResult.JSON500,
	)
}

// UpgradeMinorVersion requests a Kubernetes minor version upgrade by cluster id.
func UpgradeMinorVersion(ctx context.Context, client *v2.ServiceClient, clusterID string) (*mksclient.ClusterDetailed, error) {
	responseResult, err := client.MKSClient.UpgradeMinorVersionV2WithResponse(ctx, clusterID)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		return responseResult.JSON200.Cluster, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON404, responseResult.JSON500,
	)
}
