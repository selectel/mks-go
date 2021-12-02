package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/mks-go/pkg/testutils"
	v1 "github.com/selectel/mks-go/pkg/v1"
	"github.com/selectel/mks-go/pkg/v1/kubeoptions"
)

func TestListFeatureGates(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/feature-gates",
		RawResponse: testListFeatureGatesResponseRaw,
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

	actual, httpResponse, err := kubeoptions.ListFeatureGates(ctx, testClient)
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
	if !reflect.DeepEqual(expectedFeatureGates, actual) {
		t.Fatalf("expected %#v, \nbut got %#v", expectedFeatureGates, actual)
	}
}

func TestListFeatureGatesTimeoutError(t *testing.T) {
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

	actual, httpResponse, err := kubeoptions.ListFeatureGates(ctx, testClient)

	if actual != nil {
		t.Fatal("expected no Kubernetes versions from the List method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListFeatureGatesUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/feature-gates",
		RawResponse: testFeatureGatesAsRawList,
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

	actual, httpResponse, err := kubeoptions.ListFeatureGates(ctx, testClient)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no Kubernetes versions from the List method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListAdmissionControllers(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/admission-controllers",
		RawResponse: testListAdmissionControllersResponseRaw,
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

	actual, httpResponse, err := kubeoptions.ListAdmissionControllers(ctx, testClient)
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
	if !reflect.DeepEqual(expectedAdmissionControllers, actual) {
		t.Fatalf("expected %#v, \nbut got %#v", expectedAdmissionControllers, actual)
	}
}

func TestListAdmissionControllersTimeoutError(t *testing.T) {
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

	actual, httpResponse, err := kubeoptions.ListAdmissionControllers(ctx, testClient)

	if actual != nil {
		t.Fatal("expected no Kubernetes versions from the List method")
	}
	if httpResponse != nil {
		t.Fatal("expected no HTTP response from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListAdmissionControllersUnmarshallError(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/v1/admission-controllers",
		RawResponse: testAdmissionControllersAsRawList,
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

	actual, httpResponse, err := kubeoptions.ListAdmissionControllers(ctx, testClient)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
		t.Fatal("expected no Kubernetes versions from the List method")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}
