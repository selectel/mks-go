package cluster

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	v1 "github.com/selectel/mks-go/pkg/v1"
)

// Get returns a single cluster by its id.
func Get(ctx context.Context, client *v1.ServiceClient, clusterID string) (*View, *v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster, clusterID}, "/")
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

// Update requests an update of an existing cluster.
func Update(ctx context.Context, client *v1.ServiceClient, clusterID string, opts *UpdateOpts) (*View, *v1.ResponseResult, error) {
	updateClusterOpts := struct {
		Cluster *UpdateOpts `json:"cluster"`
	}{
		Cluster: opts,
	}
	requestBody, err := json.Marshal(updateClusterOpts)
	if err != nil {
		return nil, nil, err
	}

	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster, clusterID}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodPut, url, bytes.NewReader(requestBody))
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
func Delete(ctx context.Context, client *v1.ServiceClient, clusterID string) (*v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster, clusterID}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		err = responseResult.Err
	}

	return responseResult, err
}

// GetKubeconfig returns a kubeconfig by cluster id.
func GetKubeconfig(ctx context.Context, client *v1.ServiceClient, clusterID string) ([]byte, *v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster, clusterID, v1.ResourceURLKubeconfig}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract kubeconfig from the response body.
	kubeconfig, err := responseResult.ExtractRaw()
	if err != nil {
		return nil, responseResult, err
	}

	return kubeconfig, responseResult, nil
}

func getFieldFromKubeconfig(kubeconfig []byte, fieldName string) (string, error) {
	r := regexp.MustCompile(fmt.Sprintf("%s.*", fieldName))
	if s := r.FindString(string(kubeconfig)); len(s) != 0 {
		if subS := strings.Split(s, " "); len(subS) > 1 {
			return subS[1], nil
		}
		return "", fmt.Errorf("invalid %s field in the kubeconfig", fieldName)
	}
	return "", fmt.Errorf("unable to find %s field in kubeconfig", fieldName)
}

// GetParsedKubeconfig is a small helper function to get map of values from kubeconfig that can be useful for tf provider for example.
func GetParsedKubeconfig(ctx context.Context, client *v1.ServiceClient, clusterID string) (map[string]string, *v1.ResponseResult, error) {
	kubeconfig, responceResult, err := GetKubeconfig(ctx, client, clusterID)
	if err != nil {
		return nil, nil, err
	}
	if responceResult.Err != nil {
		return nil, responceResult, responceResult.Err
	}

	fieldsToParseAs := map[string]string{
		"certificate-authority-data": "cluster_ca",
		"server":                     "server",
		"client-certificate-data":    "client_cert",
		"client-key-data":            "client_key",
	}
	parsedKubeconfig := make(map[string]string)

	for rawName, parsedName := range fieldsToParseAs {
		if s, err := getFieldFromKubeconfig(kubeconfig, rawName); err == nil {
			parsedKubeconfig[parsedName] = s
		} else {
			return nil, responceResult, err
		}
	}

	parsedKubeconfig["raw_config"] = string(kubeconfig)

	return parsedKubeconfig, responceResult, nil
}

// RotateCerts requests a rotation of cluster certificates by cluster id.
func RotateCerts(ctx context.Context, client *v1.ServiceClient, clusterID string) (*v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster, clusterID, v1.ResourceURLRotateCerts}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		err = responseResult.Err
	}

	return responseResult, err
}

// UpgradePatchVersion requests a Kubernetes patch version upgrade by cluster id.
func UpgradePatchVersion(ctx context.Context, client *v1.ServiceClient, clusterID string) (*View, *v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster, clusterID, v1.ResourceURLUpgradePatchVersion}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodPost, url, nil)
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

// UpgradeMinorVersion requests a Kubernetes minor version upgrade by cluster id.
func UpgradeMinorVersion(ctx context.Context, client *v1.ServiceClient, clusterID string) (*View, *v1.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLCluster, clusterID, v1.ResourceURLUpgradeMinorVersion}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodPost, url, nil)
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
