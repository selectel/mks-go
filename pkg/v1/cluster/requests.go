package cluster

import (
	"context"
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
