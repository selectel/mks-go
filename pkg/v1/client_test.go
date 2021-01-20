package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/selectel/mks-go/pkg/testutils"
)

func TestNewMKSClientV1(t *testing.T) {
	tokenID := "fakeID"
	endpoint := "http://example.org"
	expected := &ServiceClient{
		TokenID:   tokenID,
		Endpoint:  endpoint,
		UserAgent: userAgent,
	}

	actual := NewMKSClientV1(tokenID, endpoint)

	if expected.TokenID != actual.TokenID {
		t.Errorf("expected Endpoint %s, but got %s", expected.Endpoint, actual.Endpoint)
	}
	if expected.Endpoint != actual.Endpoint {
		t.Errorf("expected TokenID %s, but got %s", expected.TokenID, actual.TokenID)
	}
	if expected.UserAgent != actual.UserAgent {
		t.Errorf("expected UserAgent %s, but got %s", expected.UserAgent, actual.UserAgent)
	}
	if actual.HTTPClient == nil {
		t.Errorf("expected initialized HTTPClient but it's nil")
	}
}

func TestNewMKSClientV1WithCustomHTTP(t *testing.T) {
	tokenID := testutils.TokenID
	endpoint := "http://example.org"
	expected := &ServiceClient{
		TokenID:   tokenID,
		Endpoint:  endpoint,
		UserAgent: userAgent,
	}
	customHTTPClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	actual := NewMKSClientV1WithCustomHTTP(customHTTPClient, tokenID, endpoint)

	if expected.TokenID != actual.TokenID {
		t.Errorf("expected Endpoint %s, but got %s", expected.Endpoint, actual.Endpoint)
	}
	if expected.Endpoint != actual.Endpoint {
		t.Errorf("expected TokenID %s, but got %s", expected.TokenID, actual.TokenID)
	}
	if expected.UserAgent != actual.UserAgent {
		t.Errorf("expected UserAgent %s, but got %s", expected.UserAgent, actual.UserAgent)
	}
	if actual.HTTPClient == nil {
		t.Errorf("expected initialized HTTPClient but it's nil")
	}
}

func TestDoGetRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, "response")

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := &ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		TokenID:    "token",
		UserAgent:  "agent",
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("got %d response status, want 200", response.StatusCode)
	}
}

func TestDoPostRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, "response")

		if r.Method != http.MethodPost {
			t.Errorf("got %s method, want POST", r.Method)
		}

		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unable to read the request body: %v", err)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := &ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		TokenID:    "token",
		UserAgent:  "agent",
	}

	requestBody, err := json.Marshal(&struct {
		ID string `json:"id"`
	}{
		ID: "uuid",
	})
	if err != nil {
		t.Fatalf("can't marshal JSON: %v", err)
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodPost, endpoint, bytes.NewReader(requestBody))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("got %d response status, want 200", response.StatusCode)
	}
}

func TestDoErrNotFoundRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error":{"id":"9fb12d6e-0da2-4db1-a076-414059cfb448","message":"Cluster not found"}}`)

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := &ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		TokenID:    "token",
		UserAgent:  "agent",
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("got %d response status, want 404", response.StatusCode)
	}

	if response.ErrNotFound.Error.Message != "Cluster not found" {
		t.Fatalf("got %s error message, want 'Cluster not found'", response.ErrNotFound.Error.Message)
	}

	if response.ErrNotFound.Error.ID != "9fb12d6e-0da2-4db1-a076-414059cfb448" {
		t.Fatalf("got %s object id, want '9fb12d6e-0da2-4db1-a076-414059cfb448'", response.ErrNotFound.Error.ID)
	}
}

func TestDoErrGenericRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":{"message":"cluster_id value is invalid"}}`)

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := &ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		TokenID:    "token",
		UserAgent:  "agent",
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("got %d response status, want 400", response.StatusCode)
	}

	if response.ErrGeneric.Error.Message != "cluster_id value is invalid" {
		t.Fatalf("got %s error message, want 'cluster_id value is invalid'", response.ErrGeneric.Error.Message)
	}
}

func TestDoErrNoContentRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, "") // write no content in the response body.

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := &ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		TokenID:    "token",
		UserAgent:  "agent",
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusBadGateway {
		t.Fatalf("got %d response status, want 502", response.StatusCode)
	}

	if response.Err.Error() != "mks-go: got the 502 status code from the server" {
		t.Fatalf("got %s error message, want 'mks-go: got the 502 status code from the server'", response.Err.Error())
	}
}

func TestDoErrRequestUnmarshalError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{") // write invalid json in the response body.

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := &ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		TokenID:    "token",
		UserAgent:  "agent",
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("got %d response status, want 500", response.StatusCode)
	}

	if response.Err.Error() != "mks-go: got the 500 status code from the server" {
		t.Fatalf("got %s error message, want 'mks-go: got the 500 status code from the server'", response.Err.Error())
	}
}
