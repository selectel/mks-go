package mks

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/selectel/mks-go/v2/pkg/mksclient"
)

const (
	// appName represents an application name.
	appName = "mks-go/v2"

	// defaultHTTPTimeout represents the default timeout (in seconds) for HTTP requests.
	defaultHTTPTimeout = 120

	// defaultDialTimeout represents the default timeout (in seconds) for HTTP connection establishments.
	defaultDialTimeout = 60

	// defaultKeepaliveTimeout represents the default keep-alive period for an active network connection.
	defaultKeepaliveTimeout = 60

	// defaultMaxIdleConns represents the maximum number of idle (keep-alive) connections.
	defaultMaxIdleConns = 100

	// defaultIdleConnTimeout represents the maximum amount of time an idle (keep-alive) connection will remain
	// idle before closing itself.
	defaultIdleConnTimeout = 100

	// defaultTLSHandshakeTimeout represents the default timeout (in seconds) for TLS handshake.
	defaultTLSHandshakeTimeout = 60

	// defaultExpectContinueTimeout represents the default amount of time to wait for a server's first
	// response headers.
	defaultExpectContinueTimeout = 1

	// selfPath represents the package self path.
	selfPath = "github.com/selectel/mks-go/v2"

	// defaultVersion represents the default version.
	defaultVersion = "0.0.0"
)

var (
	// appVersion is a version of the application.
	appVersion = getAppVersion()

	// userAgent contains a basic user agent that will be used in queries.
	userAgent = appName + "/" + appVersion
)

// appVersion tries to get the app version.
// If impossible to get build info, appVersion returns defaultVersion.
func getAppVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return defaultVersion
	}

	for _, dep := range info.Deps {
		if dep.Path == selfPath {
			version, _ := strings.CutPrefix(dep.Version, "v")

			return version
		}
	}

	return defaultVersion
}

// ServiceClient stores details that are needed to work with Selectel Managed Kubernetes Service API.
type ServiceClient struct {
	MKSClient mksclient.ClientWithResponsesInterface

	// TokenID is a client authentication token.
	TokenID string

	// Endpoint represents an endpoint that will be used in all requests.
	Endpoint string

	// UserAgent contains user agent that will be used in all requests.
	UserAgent string
}

// NewMKSClientV2 initializes a new MKS client for the V2 API.
func NewMKSClientV2(tokenID, endpoint string) (*ServiceClient, error) {
	return NewMKSClientV2WithCustomHTTP(nil, tokenID, endpoint)
}

// NewMKSClientV2WithCustomHTTP initializes a new MKS client for the V2 API using custom HTTP client.
// If custom HTTP client is nil - default HTTP client will be used.
func NewMKSClientV2WithCustomHTTP(customHTTPClient *http.Client, tokenID, endpoint string) (*ServiceClient, error) {
	if customHTTPClient == nil {
		customHTTPClient = newHTTPClient()
	}

	s := &ServiceClient{
		TokenID:   tokenID,
		Endpoint:  endpoint,
		UserAgent: userAgent,
	}

	mksClient, err := s.newMKSClient(customHTTPClient, endpoint)
	if err != nil {
		return nil, err
	}

	s.MKSClient = mksClient

	return s, nil
}

// newHTTPClient returns a reference to an initialized and configured HTTP client.
func newHTTPClient() *http.Client {
	return &http.Client{
		Timeout:   defaultHTTPTimeout * time.Second,
		Transport: newHTTPTransport(),
	}
}

// newMKSClient returns a reference to an initialized and configured MKS client.
func (s *ServiceClient) newMKSClient(httpClient *http.Client, endpoint string) (*mksclient.ClientWithResponses, error) {
	client, err := mksclient.NewClientWithResponses(
		endpoint,
		mksclient.WithHTTPClient(httpClient),
		mksclient.WithRequestEditorFn(s.withUserAgent),
		mksclient.WithRequestEditorFn(s.withAuthToken),
	)
	if err != nil {
		return nil, fmt.Errorf("create mks-client: %w", err)
	}

	return client, nil
}

// withUserAgent adds User-Agent header to request.
func (s *ServiceClient) withUserAgent(_ context.Context, req *http.Request) error {
	req.Header.Set("User-Agent", s.UserAgent)

	return nil
}

// withAuthToken adds X-Auth-Token header to request.
func (s *ServiceClient) withAuthToken(_ context.Context, req *http.Request) error {
	req.Header.Set("X-Auth-Token", s.TokenID)

	return nil
}

// newHTTPTransport returns a reference to an initialized and configured HTTP transport.
func newHTTPTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   defaultDialTimeout * time.Second,
			KeepAlive: defaultKeepaliveTimeout * time.Second,
		}).DialContext,
		MaxIdleConns:          defaultMaxIdleConns,
		IdleConnTimeout:       defaultIdleConnTimeout * time.Second,
		TLSHandshakeTimeout:   defaultTLSHandshakeTimeout * time.Second,
		ExpectContinueTimeout: defaultExpectContinueTimeout * time.Second,
	}
}
