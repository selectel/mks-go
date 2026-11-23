package cplogs

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/selectel/mks-go/v2/internal/common"
	mks "github.com/selectel/mks-go/v2/pkg"
	"github.com/selectel/mks-go/v2/pkg/mksclient"
	mksmock "github.com/selectel/mks-go/v2/pkg/mksclient/mocks"
)

func TestGet(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.GetControlPlaneLoggingV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.GetControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.ControlPlaneLogging{
					Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
					Group:      common.Ptr("test-group"),
				},
			},
		},
		{
			name: common.NameAccessDenied,
			clientResponse: &mksclient.GetControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusForbidden,
					Status:     http.StatusText(http.StatusForbidden),
				},
				JSON403: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgAccessDenied,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusForbidden,
				Message:    common.MsgAccessDenied,
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.GetControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     http.StatusText(http.StatusNotFound),
				},
				JSON404: &mksclient.GenericNotFoundError{
					Error: struct {
						Id      string `json:"id"` //nolint:revive // it's generated struct
						Message string `json:"message"`
					}{
						Id:      clusterID,
						Message: common.MsgClusterNotFound,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusNotFound,
				Message:    common.MsgClusterNotFound,
			},
		},
		{
			name: common.NameInternalError,
			clientResponse: &mksclient.GetControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusInternalServerError,
					Status:     http.StatusText(http.StatusInternalServerError),
				},
				JSON500: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgInternalError,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusInternalServerError,
				Message:    common.MsgInternalError,
			},
		},
		{
			name: common.NameUnknownStatus,
			clientResponse: &mksclient.GetControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusServiceUnavailable,
					Status:     http.StatusText(http.StatusServiceUnavailable),
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusServiceUnavailable,
				Message:    http.StatusText(http.StatusServiceUnavailable),
			},
		},
		{
			name: common.NameUnknownStatus + " with body",
			clientResponse: &mksclient.GetControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusTeapot,
					Status:     http.StatusText(http.StatusTeapot),
				},
				Body: []byte(`{"error":{"message":"teapot"}}`),
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusTeapot,
				Message:    http.StatusText(http.StatusTeapot) + ": teapot",
			},
		},
		{
			name:        common.NameHTTPError,
			clientError: httpError,
			errExpected: httpError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mksClient := mksmock.NewMockClientWithResponsesInterface(t)
			mksClient.EXPECT().GetControlPlaneLoggingV2WithResponse(mock.Anything, clusterID).Return(test.clientResponse, test.clientError)

			logging, err := Get(context.Background(), &mks.ServiceClient{MKSClient: mksClient}, clusterID)

			if test.errExpected != nil {
				assert.Nil(t, logging)
				require.Error(t, err)

				var mksErrExp *mksclient.MKSError
				if !errors.As(test.errExpected, &mksErrExp) {
					assert.ErrorIs(t, err, test.errExpected)

					return
				}

				var mksErr *mksclient.MKSError
				require.ErrorAs(t, err, &mksErr)
				assert.Equal(t, mksErrExp, mksErr)

				return
			}

			require.NoError(t, err)

			assert.Equal(t, test.clientResponse.JSON200.Components, logging.Components)
			assert.Equal(t, test.clientResponse.JSON200.Group, logging.Group)
		})
	}
}

func TestInit(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		requestBody    *mksclient.ControlPlaneLogging
		clientResponse *mksclient.InitControlPlaneLoggingV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			requestBody: &mksclient.ControlPlaneLogging{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
				Group:      common.Ptr("test-group"),
			},
			clientResponse: &mksclient.InitControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusCreated,
					Status:     http.StatusText(http.StatusCreated),
				},
				JSON201: &mksclient.ControlPlaneLogging{
					Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
					Group:      common.Ptr("test-group"),
				},
			},
		},
		{
			name: common.NameAccessDenied,
			requestBody: &mksclient.ControlPlaneLogging{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
				Group:      common.Ptr("test-group"),
			},
			clientResponse: &mksclient.InitControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusForbidden,
					Status:     http.StatusText(http.StatusForbidden),
				},
				JSON403: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgAccessDenied,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusForbidden,
				Message:    common.MsgAccessDenied,
			},
		},
		{
			name: common.NameNotFound,
			requestBody: &mksclient.ControlPlaneLogging{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
				Group:      common.Ptr("test-group"),
			},
			clientResponse: &mksclient.InitControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     http.StatusText(http.StatusNotFound),
				},
				JSON404: &mksclient.GenericNotFoundError{
					Error: struct {
						Id      string `json:"id"` //nolint:revive // it's generated struct
						Message string `json:"message"`
					}{
						Id:      clusterID,
						Message: common.MsgClusterNotFound,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusNotFound,
				Message:    common.MsgClusterNotFound,
			},
		},
		{
			name: common.NameConflict,
			requestBody: &mksclient.ControlPlaneLogging{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
				Group:      common.Ptr("test-group"),
			},
			clientResponse: &mksclient.InitControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusConflict,
					Status:     http.StatusText(http.StatusConflict),
				},
				JSON409: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgConflict,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusConflict,
				Message:    common.MsgConflict,
			},
		},
		{
			name: common.NameInternalError,
			requestBody: &mksclient.ControlPlaneLogging{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
				Group:      common.Ptr("test-group"),
			},
			clientResponse: &mksclient.InitControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusInternalServerError,
					Status:     http.StatusText(http.StatusInternalServerError),
				},
				JSON500: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgInternalError,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusInternalServerError,
				Message:    common.MsgInternalError,
			},
		},
		{
			name: common.NameUnknownStatus,
			requestBody: &mksclient.ControlPlaneLogging{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
				Group:      common.Ptr("test-group"),
			},
			clientResponse: &mksclient.InitControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusServiceUnavailable,
					Status:     http.StatusText(http.StatusServiceUnavailable),
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusServiceUnavailable,
				Message:    http.StatusText(http.StatusServiceUnavailable),
			},
		},
		{
			name: common.NameHTTPError,
			requestBody: &mksclient.ControlPlaneLogging{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
				Group:      common.Ptr("test-group"),
			},
			clientError: httpError,
			errExpected: httpError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mksClient := mksmock.NewMockClientWithResponsesInterface(t)
			mksClient.EXPECT().InitControlPlaneLoggingV2WithResponse(
				mock.Anything, clusterID, *test.requestBody,
			).Return(test.clientResponse, test.clientError)

			logging, err := Init(
				context.Background(),
				&mks.ServiceClient{MKSClient: mksClient},
				clusterID,
				test.requestBody,
			)

			if test.errExpected != nil {
				assert.Nil(t, logging)
				require.Error(t, err)

				var mksErrExp *mksclient.MKSError
				if !errors.As(test.errExpected, &mksErrExp) {
					assert.ErrorIs(t, err, test.errExpected)

					return
				}

				var mksErr *mksclient.MKSError
				require.ErrorAs(t, err, &mksErr)
				assert.Equal(t, mksErrExp, mksErr)

				return
			}

			require.NoError(t, err)

			assert.Equal(t, test.clientResponse.JSON201.Components, logging.Components)
			assert.Equal(t, test.clientResponse.JSON201.Group, logging.Group)
		})
	}
}

func TestUpdate(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		requestBody    *mksclient.ControlPlaneLoggingUpdateBody
		clientResponse *mksclient.UpdateControlPlaneLoggingV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			requestBody: &mksclient.ControlPlaneLoggingUpdateBody{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
			},
			clientResponse: &mksclient.UpdateControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNoContent,
					Status:     http.StatusText(http.StatusNoContent),
				},
			},
		},
		{
			name: common.NameAccessDenied,
			requestBody: &mksclient.ControlPlaneLoggingUpdateBody{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
			},
			clientResponse: &mksclient.UpdateControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusForbidden,
					Status:     http.StatusText(http.StatusForbidden),
				},
				JSON403: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgAccessDenied,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusForbidden,
				Message:    common.MsgAccessDenied,
			},
		},
		{
			name: common.NameNotFound,
			requestBody: &mksclient.ControlPlaneLoggingUpdateBody{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
			},
			clientResponse: &mksclient.UpdateControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     http.StatusText(http.StatusNotFound),
				},
				JSON404: &mksclient.GenericNotFoundError{
					Error: struct {
						Id      string `json:"id"` //nolint:revive // it's generated struct
						Message string `json:"message"`
					}{
						Id:      clusterID,
						Message: common.MsgClusterNotFound,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusNotFound,
				Message:    common.MsgClusterNotFound,
			},
		},
		{
			name: common.NameConflict,
			requestBody: &mksclient.ControlPlaneLoggingUpdateBody{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
			},
			clientResponse: &mksclient.UpdateControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusConflict,
					Status:     http.StatusText(http.StatusConflict),
				},
				JSON409: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgConflict,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusConflict,
				Message:    common.MsgConflict,
			},
		},
		{
			name: common.NameInternalError,
			requestBody: &mksclient.ControlPlaneLoggingUpdateBody{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
			},
			clientResponse: &mksclient.UpdateControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusInternalServerError,
					Status:     http.StatusText(http.StatusInternalServerError),
				},
				JSON500: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgInternalError,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusInternalServerError,
				Message:    common.MsgInternalError,
			},
		},
		{
			name: common.NameUnknownStatus,
			requestBody: &mksclient.ControlPlaneLoggingUpdateBody{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
			},
			clientResponse: &mksclient.UpdateControlPlaneLoggingV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusServiceUnavailable,
					Status:     http.StatusText(http.StatusServiceUnavailable),
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusServiceUnavailable,
				Message:    http.StatusText(http.StatusServiceUnavailable),
			},
		},
		{
			name: common.NameHTTPError,
			requestBody: &mksclient.ControlPlaneLoggingUpdateBody{
				Components: common.Ptr([]string{"APISERVER", "AUTOSCALER"}),
			},
			clientError: httpError,
			errExpected: httpError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mksClient := mksmock.NewMockClientWithResponsesInterface(t)
			mksClient.EXPECT().UpdateControlPlaneLoggingV2WithResponse(
				mock.Anything, clusterID, *test.requestBody,
			).Return(test.clientResponse, test.clientError)

			err := Update(
				context.Background(),
				&mks.ServiceClient{MKSClient: mksClient},
				clusterID,
				test.requestBody,
			)

			if test.errExpected != nil {
				require.Error(t, err)

				var mksErrExp *mksclient.MKSError
				if !errors.As(test.errExpected, &mksErrExp) {
					assert.ErrorIs(t, err, test.errExpected)

					return
				}

				var mksErr *mksclient.MKSError
				require.ErrorAs(t, err, &mksErr)
				assert.Equal(t, mksErrExp, mksErr)

				return
			}

			require.NoError(t, err)
		})
	}
}

func TestListComponents(t *testing.T) {
	httpError := errors.New("error")

	components := []mksclient.LogComponent{
		{
			Name:        "apiserver",
			Description: "kube-apiserver logs",
			DisplayName: "kube-apiserver",
		},
		{
			Name:        "autoscaler",
			Description: "cluster-autoscaler logs",
			DisplayName: "cluster-autoscaler",
		},
	}

	tests := []struct {
		name           string
		clientResponse *mksclient.ListLogComponentsV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.ListLogComponentsV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.LogComponentsList{
					Components: &components,
				},
			},
		},
		{
			name: common.NameEmptyLogComponents,
			clientResponse: &mksclient.ListLogComponentsV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.LogComponentsList{},
			},
		},
		{
			name: common.NameAccessDenied,
			clientResponse: &mksclient.ListLogComponentsV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusForbidden,
					Status:     http.StatusText(http.StatusForbidden),
				},
				JSON403: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgAccessDenied,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusForbidden,
				Message:    common.MsgAccessDenied,
			},
		},
		{
			name: common.NameInternalError,
			clientResponse: &mksclient.ListLogComponentsV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusInternalServerError,
					Status:     http.StatusText(http.StatusInternalServerError),
				},
				JSON500: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgInternalError,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusInternalServerError,
				Message:    common.MsgInternalError,
			},
		},
		{
			name: common.NameUnknownStatus,
			clientResponse: &mksclient.ListLogComponentsV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusServiceUnavailable,
					Status:     http.StatusText(http.StatusServiceUnavailable),
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusServiceUnavailable,
				Message:    http.StatusText(http.StatusServiceUnavailable),
			},
		},
		{
			name:        common.NameHTTPError,
			clientError: httpError,
			errExpected: httpError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mksClient := mksmock.NewMockClientWithResponsesInterface(t)
			mksClient.EXPECT().ListLogComponentsV2WithResponse(
				mock.Anything, &mksclient.ListLogComponentsV2Params{},
			).Return(test.clientResponse, test.clientError)

			logComponents, err := ListComponents(context.Background(), &mks.ServiceClient{MKSClient: mksClient})

			if test.errExpected != nil {
				assert.Nil(t, logComponents)
				require.Error(t, err)

				var mksErrExp *mksclient.MKSError
				if !errors.As(test.errExpected, &mksErrExp) {
					assert.ErrorIs(t, err, test.errExpected)

					return
				}

				var mksErr *mksclient.MKSError
				require.ErrorAs(t, err, &mksErr)
				assert.Equal(t, mksErrExp, mksErr)

				return
			}

			require.NoError(t, err)

			if test.clientResponse.JSON200.Components != nil {
				assert.Equal(t, components, logComponents)
			} else {
				assert.Equal(t, []mksclient.LogComponent{}, logComponents)
			}
		})
	}
}
