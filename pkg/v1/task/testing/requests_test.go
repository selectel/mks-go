package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/mks-go/pkg/testutils"
	v1 "github.com/selectel/mks-go/pkg/v1"
	"github.com/selectel/mks-go/pkg/v1/task"
)

func TestGetTask(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/d2e16a48-a9c5-4449-8b71-71f21fc872db/tasks/2f6fb93c-cf0d-4289-a78c-34393ac75f92",
		RawResponse: testGetTaskResponseRaw,
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
	clusterID := "d2e16a48-a9c5-4449-8b71-71f21fc872db"
	taskID := "2f6fb93c-cf0d-4289-a78c-34393ac75f92"

	actual, httpResponse, err := task.Get(ctx, testClient, clusterID, taskID)

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
	if !reflect.DeepEqual(expectedGetTaskResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetTaskResponse, actual)
	}
}

func TestGetTaskHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/d2d16a48-a9c5-4449-8b71-71f21fc872db/tasks/2f7fb93c-cf0d-4289-a78c-34393ac75f92",
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
	clusterID := "d2d16a48-a9c5-4449-8b71-71f21fc872db"
	taskID := "2f7fb93c-cf0d-4289-a78c-34393ac75f92"

	actual, httpResponse, err := task.Get(ctx, testClient, clusterID, taskID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no task from the Get method")
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
	clusterID := "d2e16a48-a9c5-4449-8b71-71f21fc872db"
	taskID := "2f8fb93c-cf0d-4289-a78c-34393ac75f92"

	actual, httpResponse, err := task.Get(ctx, testClient, clusterID, taskID)

	if actual != nil {
		t.Fatal("expected no task from the Get method")
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
		URL:         "/v1/clusters/d2916a40-a9c5-4449-8b71-71f21fc872db/tasks/2fbfb93c-cf0d-4289-a78c-34393ac75f92",
		RawResponse: testSingleTaskInvalidResponseRaw,
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
	clusterID := "d2916a40-a9c5-4449-8b71-71f21fc872db"
	taskID := "2fbfb93c-cf0d-4289-a78c-34393ac75f92"

	actual, httpResponse, err := task.Get(ctx, testClient, clusterID, taskID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no task from the Get method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}
