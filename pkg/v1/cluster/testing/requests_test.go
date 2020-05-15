package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/mks-go/pkg/testutils"
	v1 "github.com/selectel/mks-go/pkg/v1"
	"github.com/selectel/mks-go/pkg/v1/cluster"
)

func TestGetCluster(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/dbe7559b-55d8-4f65-9230-6a22b985ff73",
		RawResponse: testGetClusterResponseRaw,
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
	id := "dbe7559b-55d8-4f65-9230-6a22b985ff73"

	actual, httpResponse, err := cluster.Get(ctx, testClient, id)

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
	if !reflect.DeepEqual(expectedGetClusterResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetClusterResponse, actual)
	}
}

func TestGetClusterHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/dbe7759b-55d8-4f65-9230-6a22b985ff73",
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
	id := "dbe7759b-55d8-4f65-9230-6a22b985ff73"

	actual, httpResponse, err := cluster.Get(ctx, testClient, id)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no cluster from the Get method")
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

func TestGetClusterTimeoutError(t *testing.T) {
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
	id := "dbe1159b-55d8-4f65-9230-6a22b985ff73"

	actual, httpResponse, err := cluster.Get(ctx, testClient, id)

	if actual != nil {
		t.Fatal("expected no cluster from the Get method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetClusterUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/dbe7559b-55d8-4f65-9230-6a22b985ff74",
		RawResponse: testSingleClusterInvalidResponseRaw,
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
	id := "dbe7559b-55d8-4f65-9230-6a22b985ff74"

	actual, httpResponse, err := cluster.Get(ctx, testClient, id)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no cluster from the Get method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestListClusters(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters",
		RawResponse: testListClustersResponseRaw,
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

	actual, httpResponse, err := cluster.List(ctx, testClient)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the List method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedListClustersResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedListClustersResponse, actual)
	}
}

func TestListClustersHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters",
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

	actual, httpResponse, err := cluster.List(ctx, testClient)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no cluster from the List method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListClustersTimeoutError(t *testing.T) {
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

	actual, httpResponse, err := cluster.List(ctx, testClient)

	if actual != nil {
		t.Fatal("expected no cluster from the List method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListClustersUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters",
		RawResponse: testManyClustersInvalidResponseRaw,
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

	actual, httpResponse, err := cluster.List(ctx, testClient)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no cluster from the List method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestCreateCluster(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters",
		RawResponse: testCreateClusterResponseRaw,
		RawRequest:  testCreateClusterOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := cluster.Create(ctx, testClient, testCreateClusterOpts)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Create method")
	}
	if httpResponse.StatusCode != http.StatusCreated {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusCreated, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedCreateClusterResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedCreateClusterResponse, actual)
	}
}

func TestCreateClusterEnableBools(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters",
		RawResponse: testCreateClusterResponseRaw,
		RawRequest:  testCreateClusterEnableBoolsOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := cluster.Create(ctx, testClient, testCreateClusterEnableBoolsOpts)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Create method")
	}
	if httpResponse.StatusCode != http.StatusCreated {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusCreated, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedCreateClusterResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedCreateClusterResponse, actual)
	}
}

func TestCreateClusterDisableBools(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters",
		RawResponse: testCreateClusterDisableBoolsResponseRaw,
		RawRequest:  testCreateClusterDisableBoolsOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := cluster.Create(ctx, testClient, testCreateClusterDisableBoolsOpts)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Create method")
	}
	if httpResponse.StatusCode != http.StatusCreated {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusCreated, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedCreateClusterDisableBoolsResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedCreateClusterDisableBoolsResponse, actual)
	}
}

func TestCreateClusterHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters",
		RawResponse: testErrGenericResponseRaw,
		RawRequest:  testCreateClusterOptsRaw,
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

	actual, httpResponse, err := cluster.Create(ctx, testClient, testCreateClusterOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no cluster from the Create method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestCreateClusterTimeoutError(t *testing.T) {
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

	actual, httpResponse, err := cluster.Create(ctx, testClient, testCreateClusterOpts)

	if actual != nil {
		t.Fatal("expected no cluster from the Create method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestCreateClusterUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters",
		RawResponse: testSingleClusterInvalidResponseRaw,
		RawRequest:  testCreateClusterOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}

	actual, httpResponse, err := cluster.Create(ctx, testClient, testCreateClusterOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no cluster from the Create method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestUpdateCluster(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/57b244b4-a49e-4067-a837-e0024a3a8aed",
		RawResponse: testGetClusterResponseRaw,
		RawRequest:  testUpdateClusterOptsRaw,
		Method:      http.MethodPut,
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
	id := "57b244b4-a49e-4067-a837-e0024a3a8aed"

	actual, httpResponse, err := cluster.Update(ctx, testClient, id, testUpdateClusterOpts)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Update method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedGetClusterResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetClusterResponse, actual)
	}
}

func TestUpdateClusterHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/9b1352c2-5c79-4ff0-abf8-3a6d57c35487",
		RawResponse: testErrGenericResponseRaw,
		RawRequest:  testUpdateClusterOptsRaw,
		Method:      http.MethodPut,
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
	id := "9b1352c2-5c79-4ff0-abf8-3a6d57c35487"

	actual, httpResponse, err := cluster.Update(ctx, testClient, id, testUpdateClusterOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no cluster from the Update method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestUpdateClusterTimeoutError(t *testing.T) {
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
	id := "b91ae78a-e72c-4d38-8d2c-594f9cf1f0e1"

	actual, httpResponse, err := cluster.Update(ctx, testClient, id, testUpdateClusterOpts)

	if actual != nil {
		t.Fatal("expected no cluster from the Update method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
}

func TestUpdateClusterUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/72139577-463c-4e7d-9048-d9496041016c",
		RawResponse: testSingleClusterInvalidResponseRaw,
		RawRequest:  testUpdateClusterOptsRaw,
		Method:      http.MethodPut,
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
	id := "72139577-463c-4e7d-9048-d9496041016c"

	actual, httpResponse, err := cluster.Update(ctx, testClient, id, testUpdateClusterOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no cluster from the Update method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestDeleteCluster(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/v1/clusters/dbe7559c-55d8-4f65-9230-6a22b985ff73",
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
	id := "dbe7559c-55d8-4f65-9230-6a22b985ff73"

	httpResponse, err := cluster.Delete(ctx, testClient, id)

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

func TestDeleteClusterHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/dbe7759d-55d8-4f65-9230-6a22b985ff73",
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
	id := "dbe7759d-55d8-4f65-9230-6a22b985ff73"

	httpResponse, err := cluster.Delete(ctx, testClient, id)

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

func TestDeleteClusterTimeoutError(t *testing.T) {
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
	id := "dbe1159e-55d8-4f65-9230-6a22b985ff73"

	httpResponse, err := cluster.Delete(ctx, testClient, id)

	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Delete method")
	}
	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
}

func TestGetKubeconfig(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/dcd7559a-55d8-4f65-9230-6a22b985ff73/kubeconfig",
		RawResponse: testGetKubeconfig,
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
	id := "dcd7559a-55d8-4f65-9230-6a22b985ff73"

	actual, httpResponse, err := cluster.GetKubeconfig(ctx, testClient, id)

	expected := []byte(testGetKubeconfig)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the GetKubeconfig method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetKubeconfigHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/dbe1150b-55d8-4f65-9230-6a22b985ff47/kubeconfig",
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
	id := "dbe1150b-55d8-4f65-9230-6a22b985ff47"

	actual, httpResponse, err := cluster.GetKubeconfig(ctx, testClient, id)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no kubeconfig from the GetKubeconfig method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the GetKubeconfig method")
	}
	if err == nil {
		t.Fatal("expected error from the GetKubeconfig method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetKubeconfigTimeoutError(t *testing.T) {
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
	id := "dbe11593b-55d8-4f65-9230-6a22b985ff47"

	actual, httpResponse, err := cluster.Get(ctx, testClient, id)

	if actual != nil {
		t.Fatal("expected no kubeconfig from the GetKubeconfig method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the GetKubeconfig method")
	}
	if err == nil {
		t.Fatal("expected error from the GetKubeconfig method")
	}
}

func TestRotateCerts(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/v1/clusters/dbe1259c-55d8-4f65-9230-6a22b985ff73/rotate-certs",
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
	id := "dbe1259c-55d8-4f65-9230-6a22b985ff73"

	httpResponse, err := cluster.RotateCerts(ctx, testClient, id)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the RotateCerts method")
	}
	if httpResponse.StatusCode != http.StatusNoContent {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusNoContent, httpResponse.StatusCode)
	}
}

func TestRotateCertsHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/dbe7761d-55d8-4f65-9230-6a22b985ff73/rotate-certs",
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
	id := "dbe7761d-55d8-4f65-9230-6a22b985ff73"

	httpResponse, err := cluster.RotateCerts(ctx, testClient, id)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the RotateCerts method")
	}
	if err == nil {
		t.Fatal("expected error from the RotateCerts method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestRotateCertsTimeoutError(t *testing.T) {
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
	id := "dbe1159e-55d8-4f65-9780-6a22b985ff73"

	httpResponse, err := cluster.RotateCerts(ctx, testClient, id)

	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the RotateCerts method")
	}
	if err == nil {
		t.Fatal("expected error from the RotateCerts method")
	}
}

func TestUpgradePatchVersion(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/fc1d8841-d8dc-4981-a97f-4cb251e3a8aa/upgrade-patch-version",
		RawResponse: testGetClusterResponseRaw,
		Method:      http.MethodPost,
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
	id := "fc1d8841-d8dc-4981-a97f-4cb251e3a8aa"

	actual, httpResponse, err := cluster.UpgradePatchVersion(ctx, testClient, id)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the UpgradePatchVersion method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusNoContent, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedGetClusterResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetClusterResponse, actual)
	}
}

func TestUpgradePatchVersionHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/6025be99-ee53-4a9f-8589-e43801b8f778/upgrade-patch-version",
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
	id := "6025be99-ee53-4a9f-8589-e43801b8f778"

	actual, httpResponse, err := cluster.UpgradePatchVersion(ctx, testClient, id)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no cluster from the UpgradePatchVersion method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the UpgradePatchVersion method")
	}
	if err == nil {
		t.Fatal("expected error from the UpgradePatchVersion method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestUpgradePatchVersionTimeoutError(t *testing.T) {
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
	id := "6af18629-f069-4bf1-888d-2476ab2bddff"

	actual, httpResponse, err := cluster.UpgradePatchVersion(ctx, testClient, id)

	if actual != nil {
		t.Fatal("expected no cluster from the UpgradePatchVersion method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the UpgradePatchVersion method")
	}
	if err == nil {
		t.Fatal("expected error from the UpgradePatchVersion method")
	}
}

func TestUpgradePatchVersionUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/7af18629-f069-4bf1-888d-2476ab2bddff/upgrade-patch-version",
		RawResponse: testSingleClusterInvalidResponseRaw,
		Method:      http.MethodPost,
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
	id := "7af18629-f069-4bf1-888d-2476ab2bddff"

	actual, httpResponse, err := cluster.UpgradePatchVersion(ctx, testClient, id)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no cluster from the UpgradePatchVersion method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the UpgradePatchVersion method")
	}
	if err == nil {
		t.Fatal("expected error from the UpgradePatchVersion method")
	}
}

func TestUpgradeMinorVersion(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/dbe7559b-55d8-4f65-9230-6a22b985ff16/upgrade-minor-version",
		RawResponse: testGetClusterResponseRaw,
		Method:      http.MethodPost,
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
	id := "dbe7559b-55d8-4f65-9230-6a22b985ff16"

	actual, httpResponse, err := cluster.UpgradeMinorVersion(ctx, testClient, id)

	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the UpgradeMinorVersion method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedGetClusterResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetClusterResponse, actual)
	}
}

func TestUpgradeMinorVersionHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/dbe7759b-55d8-4f65-9230-6a22b985ff17/upgrade-minor-version",
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
	id := "dbe7759b-55d8-4f65-9230-6a22b985ff17"

	actual, httpResponse, err := cluster.UpgradeMinorVersion(ctx, testClient, id)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no cluster from the UpgradeMinorVersion method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the UpgradeMinorVersion method")
	}
	if err == nil {
		t.Fatal("expected error from the UpgradeMinorVersion method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestUpgradeMinorVersionTimeoutError(t *testing.T) {
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
	id := "dbe1159b-55d8-4f65-9230-6a22b985ff18"

	actual, httpResponse, err := cluster.UpgradeMinorVersion(ctx, testClient, id)

	if actual != nil {
		t.Fatal("expected no cluster from the UpgradeMinorVersion method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the UpgradeMinorVersion method")
	}
	if err == nil {
		t.Fatal("expected error from the UpgradeMinorVersion method")
	}
}

func TestUpgradeMinorVersionUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/dbe7559b-55d8-4f65-9230-6a22b985ff19/upgrade-minor-version",
		RawResponse: testSingleClusterInvalidResponseRaw,
		Method:      http.MethodPost,
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
	id := "dbe7559b-55d8-4f65-9230-6a22b985ff19"

	actual, httpResponse, err := cluster.UpgradeMinorVersion(ctx, testClient, id)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no cluster from the UpgradeMinorVersion method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the UpgradeMinorVersion method")
	}
	if err == nil {
		t.Fatal("expected error from the UpgradeMinorVersion method")
	}
}
