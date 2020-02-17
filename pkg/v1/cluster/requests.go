package cluster

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	v1 "github.com/selectel/mks-go/pkg/v1"
)

// Get returns a single Cluster by its id.
func Get(ctx context.Context, client *v1.ServiceClient, id string) (*View, *v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster, id}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a cluster from the response body.
	var result struct {
		Cluster *View `json:"cluster"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Cluster, responseResult, nil
}

// List gets a list of all clusters.
func List(ctx context.Context, client *v1.ServiceClient) ([]*View, *v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract clusters from the response body.
	var result struct {
		Clusters []*View `json:"clusters"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Clusters, responseResult, nil
}

// Create requests a creation of a new cluster.
func Create(ctx context.Context, client *v1.ServiceClient, opts *CreateOpts) (*View, *v1.ResponseResult, error) {
	createClusterOpts := struct {
		Cluster *CreateOpts `json:"cluster"`
	}{
		Cluster: opts,
	}
	requestBody, err := json.Marshal(createClusterOpts)
	if err != nil {
		return nil, nil, err
	}

	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodPost, url, bytes.NewReader(requestBody))
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract cluster from the response body.
	var result struct {
		Cluster *View `json:"cluster"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Cluster, responseResult, nil
}

// Delete deletes a single cluster by its id.
func Delete(ctx context.Context, client *v1.ServiceClient, id string) (*v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster, id}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		err = responseResult.Err
	}

	return responseResult, err
}
