package registries

import (
	"context"
	"net/http"

	v2 "github.com/selectel/mks-go/pkg/v2"
	"github.com/selectel/mks-go/pkg/v2/mksclient"
)

// Get returns registries integrated with provided cluster.
func Get(ctx context.Context, client *v2.ServiceClient, clusterID string) (*mksclient.RegistriesIntegration, error) {
	responseResult, err := client.MKSClient.GetRegistriesV2WithResponse(ctx, clusterID)
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		return responseResult.JSON200, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404, responseResult.JSON500,
	)
}

// Create creates new registries integration in provided cluster.
func Create(
	ctx context.Context, client *v2.ServiceClient,
	clusterID string, registries []mksclient.RegistriesIntergrationCreateStruct,
) (*mksclient.RegistriesIntegration, error) {
	responseResult, err := client.MKSClient.CreateRegistriesV2WithResponse(
		ctx, clusterID, mksclient.RegistriesIntergrationCreateBody{Registries: registries})
	if err != nil {
		return nil, err
	}

	if responseResult.JSON201 != nil {
		return responseResult.JSON201, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404, responseResult.JSON500,
	)
}

// Update updates registries integration in provided cluster.
func Update(
	ctx context.Context, client *v2.ServiceClient,
	clusterID string, registries []mksclient.RegistriesIntergrationUpdateStruct,
) (*mksclient.RegistriesIntegration, error) {
	responseResult, err := client.MKSClient.UpdateRegistriesV2WithResponse(
		ctx, clusterID, mksclient.RegistriesIntergrationUpdateBody{Registries: registries})
	if err != nil {
		return nil, err
	}

	if responseResult.JSON200 != nil {
		return responseResult.JSON200, nil
	}

	return nil, mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404, responseResult.JSON500,
	)
}

// Delete deletes provided registry integration in provided cluster.
func Delete(
	ctx context.Context, client *v2.ServiceClient, clusterID, registryID string,
) error {
	responseResult, err := client.MKSClient.DeleteRegistryV2WithResponse(ctx, clusterID, registryID)
	if err != nil {
		return err
	}

	if responseResult.StatusCode() == http.StatusNoContent {
		return nil
	}

	return mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404, responseResult.JSON500,
	)
}

// DeleteAll deletes all registries integrations in provided cluster.
func DeleteAll(
	ctx context.Context, client *v2.ServiceClient, clusterID string,
) error {
	responseResult, err := client.MKSClient.DeleteRegistriesV2WithResponse(ctx, clusterID)
	if err != nil {
		return err
	}

	if responseResult.StatusCode() == http.StatusNoContent {
		return nil
	}

	return mksclient.HandleAPIErrors(
		responseResult.StatusCode(), responseResult.Status(),
		responseResult.JSON400, responseResult.JSON404, responseResult.JSON500,
	)
}
