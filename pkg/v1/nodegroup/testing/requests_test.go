package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/mks-go/pkg/testutils"
	v1 "github.com/selectel/mks-go/pkg/v1"
	"github.com/selectel/mks-go/pkg/v1/nodegroup"
)

func TestGetNodegroup(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/79265515-3700-49fa-af0e-7f547bce788a/nodegroups/a376745a-fbcb-413d-b418-169d059d79ce",
		RawResponse: testGetNodegroupResponseRaw,
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
	clusterID := "79265515-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "a376745a-fbcb-413d-b418-169d059d79ce"

	actual, httpResponse, err := nodegroup.Get(ctx, testClient, clusterID, nodegroupID)
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
	if !reflect.DeepEqual(expectedGetNodegroupResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetNodegroupResponse, actual)
	}
}

func TestGetNodegroupHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/79265515-3700-49fa-af0e-7f547bce777a/nodegroups/a376745a-fbcb-413d-b418-172d059d79ce",
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
	clusterID := "79265515-3700-49fa-af0e-7f547bce777a"
	nodegroupID := "a376745a-fbcb-413d-b418-172d059d79ce"

	actual, httpResponse, err := nodegroup.Get(ctx, testClient, clusterID, nodegroupID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no nodegroup from the Get method")
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

func TestGetNodegroupTimeoutError(t *testing.T) {
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
	clusterID := "79265515-3700-49fa-af0e-7f547bce888a"
	nodegroupID := "a481745a-fbcb-413d-b418-172d059d79ce"

	actual, httpResponse, err := nodegroup.Get(ctx, testClient, clusterID, nodegroupID)

	if actual != nil {
		t.Fatal("expected no nodegroup from the Get method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetNodegroupUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/15965515-3700-49fa-af0e-7f547bce788a/nodegroups/a371145a-fbcb-413d-b418-169d059d79ce",
		RawResponse: testSingleNodegroupInvalidResponseRaw,
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
	clusterID := "15965515-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "a371145a-fbcb-413d-b418-169d059d79ce"

	actual, httpResponse, err := nodegroup.Get(ctx, testClient, clusterID, nodegroupID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no nodegroup from the Get method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestListNodegroups(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/79265515-3700-49fa-af0e-7f547bce788a/nodegroups",
		RawResponse: testListNodegroupsResponseRaw,
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
	clusterID := "79265515-3700-49fa-af0e-7f547bce788a"

	actual, httpResponse, err := nodegroup.List(ctx, testClient, clusterID)
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
	if !reflect.DeepEqual(expectedListNodegroupsResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedListNodegroupsResponse, actual)
	}
}

func TestListNodegroupsHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/89265515-3700-49fa-af0e-7f547bce788a/nodegroups",
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
	clusterID := "89265515-3700-49fa-af0e-7f547bce788a"

	actual, httpResponse, err := nodegroup.List(ctx, testClient, clusterID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no nodegroup from the List method")
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

func TestListNodegroupsTimeoutError(t *testing.T) {
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
	clusterID := "89265515-3700-49fa-af0e-7f547bce788b"

	actual, httpResponse, err := nodegroup.List(ctx, testClient, clusterID)

	if actual != nil {
		t.Fatal("expected no nodegroup from the List method")
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
		URL:         "/v1/clusters/d1265515-3700-49fa-af0e-7f547bce788a/nodegroups",
		RawResponse: testManyNodegroupsInvalidResponseRaw,
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
	clusterID := "d1265515-3700-49fa-af0e-7f547bce788a"

	actual, httpResponse, err := nodegroup.List(ctx, testClient, clusterID)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no nodegroup from the List method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestCreateNodegroup(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:        testEnv.Mux,
		URL:        "/v1/clusters/d1465515-3700-49fa-af0e-7f547bce788a/nodegroups",
		RawRequest: testCreateNodegroupOptsRaw,
		Method:     http.MethodPost,
		Status:     http.StatusNoContent,
		CallFlag:   &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "d1465515-3700-49fa-af0e-7f547bce788a"

	httpResponse, err := nodegroup.Create(ctx, testClient, clusterID, testCreateNodegroupOpts)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Create method")
	}
	if httpResponse.StatusCode != http.StatusNoContent {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusNoContent, httpResponse.StatusCode)
	}
}

func TestCreateNodegroupHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/d1565515-3700-49fa-af0e-7f547bce788a/nodegroups",
		RawResponse: testErrGenericResponseRaw,
		RawRequest:  testCreateNodegroupOptsRaw,
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
	clusterID := "d1565515-3700-49fa-af0e-7f547bce788a"

	httpResponse, err := nodegroup.Create(ctx, testClient, clusterID, testCreateNodegroupOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
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

func TestCreateNodegroupTimeoutError(t *testing.T) {
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
	clusterID := "d1c65515-3700-49fa-af0e-7f547bce788a"

	httpResponse, err := nodegroup.Create(ctx, testClient, clusterID, testCreateNodegroupOpts)

	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestDeleteNodegroup(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/v1/clusters/c1d7559c-55d8-4f65-9230-6a22b985ff93/nodegroups/b376745a-fbcb-413d-b418-180d059d79cd",
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
	clusterID := "c1d7559c-55d8-4f65-9230-6a22b985ff93"
	nodegroupID := "b376745a-fbcb-413d-b418-180d059d79cd"

	httpResponse, err := nodegroup.Delete(ctx, testClient, clusterID, nodegroupID)
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

func TestDeleteNodegroupHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/c1d7559c-55d8-7f65-9230-6a22b985ff93/nodegroups/b376745a-0bcb-413d-b418-180d059d79cd",
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
	clusterID := "c1d7559c-55d8-7f65-9230-6a22b985ff93"
	nodegroupID := "b376745a-0bcb-413d-b418-180d059d79cd"

	httpResponse, err := nodegroup.Delete(ctx, testClient, clusterID, nodegroupID)

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

func TestDeleteNodegroupTimeoutError(t *testing.T) {
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
	clusterID := "c1e7559c-55d8-7f65-9230-6a22b985ff93"
	nodegroupID := "b376845a-0bcb-413d-b418-180d059d79cd"

	httpResponse, err := nodegroup.Delete(ctx, testClient, clusterID, nodegroupID)

	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Delete method")
	}
	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
}

func TestResizeNodegroup(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:        testEnv.Mux,
		URL:        "/v1/clusters/e172551c-3700-49fa-af0e-7f547bce788a/nodegroups/c476d45a-0bcc-e13d-b418-180d059d79cd/resize",
		RawRequest: testResizeNodegroupOptsRaw,
		Method:     http.MethodPost,
		Status:     http.StatusNoContent,
		CallFlag:   &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "e172551c-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "c476d45a-0bcc-e13d-b418-180d059d79cd"

	httpResponse, err := nodegroup.Resize(ctx, testClient, clusterID, nodegroupID, testResizeNodegroupOpts)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Resize method")
	}
	if httpResponse.StatusCode != http.StatusNoContent {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusNoContent, httpResponse.StatusCode)
	}
}

func TestResizeNodegroupHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/e17d551c-3700-49fa-af0e-7f547bce788a/nodegroups/c476b45a-0bcc-e13d-b418-180d059d79cd/resize",
		RawResponse: testErrGenericResponseRaw,
		RawRequest:  testResizeNodegroupOptsRaw,
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
	clusterID := "e17d551c-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "c476b45a-0bcc-e13d-b418-180d059d79cd"

	httpResponse, err := nodegroup.Resize(ctx, testClient, clusterID, nodegroupID, testResizeNodegroupOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Resize method")
	}
	if err == nil {
		t.Fatal("expected error from the Resize method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestResizeNodegroupTimeoutError(t *testing.T) {
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
	clusterID := "e17d651c-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "c476b75a-0bcc-e13d-b418-180d059d79cd"

	httpResponse, err := nodegroup.Resize(ctx, testClient, clusterID, nodegroupID, testResizeNodegroupOpts)

	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Resize method")
	}
	if err == nil {
		t.Fatal("expected error from the Resize method")
	}
}

func TestUpdateNodegroup(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:        testEnv.Mux,
		URL:        "/v1/clusters/e172551c-3700-49fa-af0e-7f547bce788a/nodegroups/c476d45a-0bcc-e13d-b418-180d059d79cd",
		RawRequest: testUpdateNodegroupOptsRaw,
		Method:     http.MethodPut,
		Status:     http.StatusNoContent,
		CallFlag:   &endpointCalled,
	})

	ctx := context.Background()
	testClient := &v1.ServiceClient{
		HTTPClient: &http.Client{},
		TokenID:    testutils.TokenID,
		Endpoint:   testEnv.Server.URL + "/v1",
		UserAgent:  testutils.UserAgent,
	}
	clusterID := "e172551c-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "c476d45a-0bcc-e13d-b418-180d059d79cd"

	httpResponse, err := nodegroup.Update(ctx, testClient, clusterID, nodegroupID, testUpdateNodegroupOpts)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Update method")
	}
	if httpResponse.StatusCode != http.StatusNoContent {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusNoContent, httpResponse.StatusCode)
	}
}

func TestUpdateNodegroupHTTPError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/clusters/e17d551c-3700-49fa-af0e-7f547bce788a/nodegroups/c476b45a-0bcc-e13d-b418-180d059d79cd",
		RawResponse: testErrGenericResponseRaw,
		RawRequest:  testUpdateNodegroupOptsRaw,
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
	clusterID := "e17d551c-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "c476b45a-0bcc-e13d-b418-180d059d79cd"

	httpResponse, err := nodegroup.Update(ctx, testClient, clusterID, nodegroupID, testUpdateNodegroupOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
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

func TestUpdateNodegroupTimeoutError(t *testing.T) {
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
	clusterID := "e17d651c-3700-49fa-af0e-7f547bce788a"
	nodegroupID := "c476b75a-0bcc-e13d-b418-180d059d79cd"

	httpResponse, err := nodegroup.Update(ctx, testClient, clusterID, nodegroupID, testUpdateNodegroupOpts)

	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
}
