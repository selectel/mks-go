package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/mks-go/pkg/testutils"
	v1 "github.com/selectel/mks-go/pkg/v1"
	"github.com/selectel/mks-go/pkg/v1/node"
)

func TestGetNode(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/792de51c-3700-49fa-af0e-7f547bce788a/nodegroups/f174b65d-442a-4423-aaf7-5654789b8a9d/203d0f8c-547d-48a7-98ed-3075254b8d4a",
		RawResponse: testGetNodeResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "792de51c-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "f174b65d-442a-4423-aaf7-5654789b8a9d"
	nodeID := "203d0f8c-547d-48a7-98ed-3075254b8d4a"

	actual, httpResponse, err := node.Get(ctx, testClient, clusterID, nodegroupID, nodeID)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Get method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedGetNodeResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetNodeResponse, actual)
	}
}

func TestGetNodeHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/792de71c-3700-49fa-af0e-7f547bce788a/nodegroups/f174b66d-442a-4423-aaf7-5654789b8a9d/203c0f8c-547d-48a7-98ed-3075254b8d4a",
		RawResponse: testErrGenericResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "792de71c-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "f174b66d-442a-4423-aaf7-5654789b8a9d"
	nodeID := "203c0f8c-547d-48a7-98ed-3075254b8d4a"

	actual, httpResponse, err := node.Get(ctx, testClient, clusterID, nodegroupID, nodeID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no node from the Get method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetNodeTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "78ede71c-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "f174b66d-442a-41e3-aaf7-5b54789b8a9d"
	nodeID := "1cdc0f8c-547d-48a7-98ed-3075254b8d4a"

	actual, httpResponse, err := node.Get(ctx, testClient, clusterID, nodegroupID, nodeID)

	if actual != nil {
		t.Fatal("expected no node from the Get method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetNodeUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/1c2de51c-3700-49fa-af0e-7f547bce788a/nodegroups/e474b65d-442a-4423-aaf7-5654789b8a9d/3dbe0f8c-547d-48a7-98ed-3075254b8d4a",
		RawResponse: testSingleNodeInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "1c2de51c-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "e474b65d-442a-4423-aaf7-5654789b8a9d"
	nodeID := "3dbe0f8c-547d-48a7-98ed-3075254b8d4a"

	actual, httpResponse, err := node.Get(ctx, testClient, clusterID, nodegroupID, nodeID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no node from the Get method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestReinstallNode(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/v1/clusters/792de51c-54b9-49fa-af0e-7f547bce788a/nodegroups/e195b65d-442a-4423-aaf7-5654789b8a9d/c1800f8c-547d-48a7-98ed-3075254b8d4a/reinstall",
		Method:   http.MethodPost,
		Status:   http.StatusNoContent,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "792de51c-54b9-49fa-af0e-7f547bce788a"
	nodegroupID := "e195b65d-442a-4423-aaf7-5654789b8a9d"
	nodeID := "c1800f8c-547d-48a7-98ed-3075254b8d4a"

	httpResponse, err := node.Reinstall(ctx, testClient, clusterID, nodegroupID, nodeID)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Reinstall method")
	}
	if httpResponse.StatusCode != http.StatusNoContent {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusNoContent, httpResponse.StatusCode)
	}
}

func TestReinstallNodeHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/792de71c-3700-49fa-bc4e-7f547bce788a/nodegroups/f174b68d-442a-4423-aaf7-5654789b8a9d/203ce18c-547d-48a7-98ed-3075254b8d4a/reinstall",
		RawResponse: testErrGenericResponseRaw,
		Method:      http.MethodPost,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "792de71c-3700-49fa-bc4e-7f547bce788a"
	nodegroupID := "f174b68d-442a-4423-aaf7-5654789b8a9d"
	nodeID := "203ce18c-547d-48a7-98ed-3075254b8d4a"

	httpResponse, err := node.Reinstall(ctx, testClient, clusterID, nodegroupID, nodeID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Reinstall method")
	}
	if err == nil {
		t.Fatal("expected error from the Reinstall method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestReinstallNodeTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "78ed321c-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "f174b55d-442a-41e3-aaf7-5b54789b8a9d"
	nodeID := "1cdc0f8c-327d-48a7-98ed-3075254b8d4a"

	httpResponse, err := node.Reinstall(ctx, testClient, clusterID, nodegroupID, nodeID)

	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Reinstall method")
	}
	if err == nil {
		t.Fatal("expected error from the Reinstall method")
	}
}

func TestDeleteNode(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/v1/clusters/892de51c-54b9-49fa-af0e-7f547bce788a/nodegroups/c195b65d-442a-4423-aaf7-5654789b8a9d/d1800f8c-547d-48a7-98ed-3075254b8d4a",
		Method:   http.MethodDelete,
		Status:   http.StatusNoContent,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "892de51c-54b9-49fa-af0e-7f547bce788a"
	nodegroupID := "c195b65d-442a-4423-aaf7-5654789b8a9d"
	nodeID := "d1800f8c-547d-48a7-98ed-3075254b8d4a"

	httpResponse, err := node.Delete(ctx, testClient, clusterID, nodegroupID, nodeID)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Delete method")
	}
	if httpResponse.StatusCode != http.StatusNoContent {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusNoContent, httpResponse.StatusCode)
	}
}

func TestDeleteNodeHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/892de71c-3700-49fa-bc4e-7f547bce788a/nodegroups/d174b68d-442a-4423-aaf7-5654789b8a9d/c03ce18c-547d-48a7-98ed-3075254b8d4a",
		RawResponse: testErrGenericResponseRaw,
		Method:      http.MethodDelete,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "892de71c-3700-49fa-bc4e-7f547bce788a"
	nodegroupID := "d174b68d-442a-4423-aaf7-5654789b8a9d"
	nodeID := "c03ce18c-547d-48a7-98ed-3075254b8d4a"

	httpResponse, err := node.Delete(ctx, testClient, clusterID, nodegroupID, nodeID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Delete method")
	}
	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestDeleteNodeTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "18ed321c-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "2174b55d-442a-41e3-aaf7-5b54789b8a9d"
	nodeID := "3cdc0f8c-327d-48a7-98ed-3075254b8d4a"

	httpResponse, err := node.Delete(ctx, testClient, clusterID, nodegroupID, nodeID)

	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Delete method")
	}
	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
}
