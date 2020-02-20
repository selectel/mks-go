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

func TestGetTaskUnknownStatusAndType(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/d2e16a48-a9c5-4449-8b71-71f21fc872dc/tasks/2f6fb93c-cf0d-4289-a78c-34393ac75f92",
		RawResponse: testGetTaskUnknownStatusAndTypeResponseRaw,
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
	clusterID := "d2e16a48-a9c5-4449-8b71-71f21fc872dc"
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
	if !reflect.DeepEqual(expectedGetTaskUnknownStatusAndTypeResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetTaskUnknownStatusAndTypeResponse, actual)
	}
}

func TestListTasks(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/d2e16a48-a9c5-4449-9b71-81f21fc872db/tasks",
		RawResponse: testListTasksResponseRaw,
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
	clusterID := "d2e16a48-a9c5-4449-9b71-81f21fc872db"

	actual, httpResponse, err := task.List(ctx, testClient, clusterID)

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
	if !reflect.DeepEqual(expectedListTasksResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedListTasksResponse, actual)
	}
}

func TestListTasksHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/d4e16a48-a9c5-4449-9b71-81f21fc872db/tasks",
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
	clusterID := "d4e16a48-a9c5-4449-9b71-81f21fc872db"

	actual, httpResponse, err := task.List(ctx, testClient, clusterID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no tasks from the List method")
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

func TestListTasksTimeoutError(t *testing.T) {
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
	clusterID := "d4416a48-a9c5-4449-9b71-81f21fc872db"

	actual, httpResponse, err := task.List(ctx, testClient, clusterID)

	if actual != nil {
		t.Fatal("expected no tasks from the List method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListTasksUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/d2e16b48-a9c5-4440-8b71-71f21fc872db/tasks",
		RawResponse: testManyTasksInvalidResponse,
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
	clusterID := "d2e16b48-a9c5-4440-8b71-71f21fc872db"

	actual, httpResponse, err := task.List(ctx, testClient, clusterID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no tasks from the List method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListTasksUnknownStatusAndType(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/d1e16a48-a9c5-4449-9b71-81f21fc872cb/tasks",
		RawResponse: testListTasksUnknownStatusAndTypeResponseRaw,
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
	clusterID := "d1e16a48-a9c5-4449-9b71-81f21fc872cb"

	actual, httpResponse, err := task.List(ctx, testClient, clusterID)

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
	if !reflect.DeepEqual(expectedListTasksUnknownStatusAndTypeResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedListTasksUnknownStatusAndTypeResponse, actual)
	}
}
