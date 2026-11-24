package mks

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAppVersion(t *testing.T) {
	assert.NotEmpty(t, getAppVersion())
}

func TestNewMKSClientV2(t *testing.T) {
	const (
		tokenID  = "token-id"
		endpoint = "https://ru-3.mks.selcloud.ru/v2"
	)

	client, err := NewMKSClientV2(tokenID, endpoint)

	require.NoError(t, err)

	assert.Equal(t, endpoint, client.Endpoint)
	assert.Equal(t, tokenID, client.TokenID)
	assert.Equal(t, userAgent, client.UserAgent)
}

func TestNewMKSClientV2WithCustomHTTP(t *testing.T) {
	const (
		tokenID  = "token-id"
		endpoint = "https://ru-3.mks.selcloud.ru/v2"
	)

	clientHTTP := &http.Client{
		Timeout: time.Hour,
	}

	client, err := NewMKSClientV2WithCustomHTTP(clientHTTP, tokenID, endpoint)

	require.NoError(t, err)

	assert.Equal(t, endpoint, client.Endpoint)
	assert.Equal(t, tokenID, client.TokenID)
	assert.Equal(t, userAgent, client.UserAgent)
}

func TestNewMKSClientV2RequestPath(t *testing.T) {
	for _, suffix := range []string{"", "/", "/v2", "/v2/"} {
		t.Run("suffix "+suffix, func(t *testing.T) {
			var requestPath string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requestPath = r.URL.Path
				w.WriteHeader(http.StatusNotFound)
			}))
			defer server.Close()

			client, err := NewMKSClientV2("token-id", server.URL+"/ru-3"+suffix)
			require.NoError(t, err)

			_, err = client.MKSClient.ListClustersV2WithResponse(context.Background())
			require.NoError(t, err)

			assert.Equal(t, "/ru-3/v2/clusters", requestPath)
		})
	}
}

func TestNewHTTPClient(t *testing.T) {
	clientHTTP := newHTTPClient()

	expectedTransport := newHTTPTransport()
	transport, ok := clientHTTP.Transport.(*http.Transport)
	require.True(t, ok)

	assert.Equal(t, defaultHTTPTimeout*time.Second, clientHTTP.Timeout)
	assert.NotNil(t, transport.Proxy)
	assert.NotNil(t, transport.DialContext)
	assert.Equal(t, expectedTransport.MaxIdleConns, transport.MaxIdleConns)
	assert.Equal(t, expectedTransport.IdleConnTimeout, transport.IdleConnTimeout)
	assert.Equal(t, expectedTransport.TLSHandshakeTimeout, transport.TLSHandshakeTimeout)
	assert.Equal(t, expectedTransport.ExpectContinueTimeout, transport.ExpectContinueTimeout)
}

func TestMKSClient(t *testing.T) {
	const (
		tokenID  = "token-id"
		endpoint = "https://ru-3.mks.selcloud.ru/v2"
	)

	var serviceClient ServiceClient

	clientHTTP := newHTTPClient()

	_, err := serviceClient.newMKSClient(clientHTTP, endpoint)

	require.NoError(t, err)
}

func TestWithUserAgent(t *testing.T) {
	const userAgent = "user agent"

	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://example.org", nil)

	require.NoError(t, err)

	serviceClient := ServiceClient{
		UserAgent: userAgent,
	}

	require.NoError(t, serviceClient.withUserAgent(context.Background(), request))

	require.Equal(t, userAgent, request.UserAgent())
}

func TestWithAuthToken(t *testing.T) {
	const token = "token-id"

	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://example.org", nil)

	require.NoError(t, err)

	serviceClient := ServiceClient{
		TokenID: token,
	}

	require.NoError(t, serviceClient.withAuthToken(context.Background(), request))

	require.Equal(t, token, request.Header.Get("X-Auth-Token"))
}
