package registries

import (
	"context"
	"errors"
	"net/http"
	"testing"

	v2 "github.com/selectel/mks-go/pkg/v2"
	"github.com/selectel/mks-go/pkg/v2/internal/common"
	"github.com/selectel/mks-go/pkg/v2/mksclient"
	mksmock "github.com/selectel/mks-go/pkg/v2/mksclient/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const testRegistryID = "test-registry-id"

func TestGet(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.GetRegistriesV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.GetRegistriesV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.RegistriesIntegration{
					Registries: []mksclient.RegistryWithMeta{
						{
							Id:   testRegistryID,
							Name: "test-registry",
						},
					},
					Status: mksclient.RegistriesIntegrationStatusACTIVE,
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.GetRegistriesV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     http.StatusText(http.StatusNotFound),
				},
				JSON404: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgRegistriesNotFound,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusNotFound,
				Message:    common.MsgRegistriesNotFound,
			},
		},
		{
			name: common.NameBadRequest,
			clientResponse: &mksclient.GetRegistriesV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusBadRequest,
					Status:     http.StatusText(http.StatusBadRequest),
				},
				JSON400: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgBadRequest,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusBadRequest,
				Message:    common.MsgBadRequest,
			},
		},
		{
			name: common.NameInternalError,
			clientResponse: &mksclient.GetRegistriesV2Response{
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
			clientResponse: &mksclient.GetRegistriesV2Response{
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
			mksClient.EXPECT().GetRegistriesV2WithResponse(mock.Anything, clusterID).Return(test.clientResponse, test.clientError)

			registries, err := Get(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID)

			if test.errExpected != nil {
				assert.Nil(t, registries)
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

			assert.Equal(t, test.clientResponse.JSON200.Registries, registries.Registries)
			assert.Equal(t, test.clientResponse.JSON200.Status, registries.Status)
		})
	}
}

func TestCreate(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.CreateRegistriesV2Response
		requestBody    mksclient.RegistriesIntergrationCreateBody
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			requestBody: mksclient.RegistriesIntergrationCreateBody{
				Registries: []mksclient.RegistriesIntergrationCreateStruct{{Id: testRegistryID}},
			},
			clientResponse: &mksclient.CreateRegistriesV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusCreated,
					Status:     http.StatusText(http.StatusCreated),
				},
				JSON201: &mksclient.RegistriesIntegration{
					Registries: []mksclient.RegistryWithMeta{
						{
							Id:   testRegistryID,
							Name: "test-registry",
						},
					},
					Status: mksclient.RegistriesIntegrationStatusACTIVE,
				},
			},
		},
		{
			name: common.NameNotFound,
			requestBody: mksclient.RegistriesIntergrationCreateBody{
				Registries: []mksclient.RegistriesIntergrationCreateStruct{{Id: testRegistryID}},
			},
			clientResponse: &mksclient.CreateRegistriesV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     http.StatusText(http.StatusNotFound),
				},
				JSON404: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
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
			name: common.NameBadRequest,
			requestBody: mksclient.RegistriesIntergrationCreateBody{
				Registries: []mksclient.RegistriesIntergrationCreateStruct{{Id: testRegistryID}},
			},
			clientResponse: &mksclient.CreateRegistriesV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusBadRequest,
					Status:     http.StatusText(http.StatusBadRequest),
				},
				JSON400: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgBadRequest,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusBadRequest,
				Message:    common.MsgBadRequest,
			},
		},
		{
			name: common.NameInternalError,
			requestBody: mksclient.RegistriesIntergrationCreateBody{
				Registries: []mksclient.RegistriesIntergrationCreateStruct{{Id: testRegistryID}},
			},
			clientResponse: &mksclient.CreateRegistriesV2Response{
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
			requestBody: mksclient.RegistriesIntergrationCreateBody{
				Registries: []mksclient.RegistriesIntergrationCreateStruct{{Id: testRegistryID}},
			},
			clientResponse: &mksclient.CreateRegistriesV2Response{
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
			requestBody: mksclient.RegistriesIntergrationCreateBody{
				Registries: []mksclient.RegistriesIntergrationCreateStruct{{Id: testRegistryID}},
			},
			clientError: httpError,
			errExpected: httpError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mksClient := mksmock.NewMockClientWithResponsesInterface(t)
			mksClient.EXPECT().CreateRegistriesV2WithResponse(
				mock.Anything, clusterID, test.requestBody,
			).Return(test.clientResponse, test.clientError)

			registries, err := Create(
				context.Background(),
				&v2.ServiceClient{MKSClient: mksClient},
				clusterID,
				test.requestBody.Registries,
			)

			if test.errExpected != nil {
				assert.Nil(t, registries)
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

			assert.Equal(t, test.clientResponse.JSON201.Registries, registries.Registries)
			assert.Equal(t, test.clientResponse.JSON201.Status, registries.Status)
		})
	}
}

func TestUpdate(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.UpdateRegistriesV2Response
		requestBody    mksclient.RegistriesIntergrationUpdateBody
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			requestBody: mksclient.RegistriesIntergrationUpdateBody{
				Registries: []mksclient.RegistriesIntergrationUpdateStruct{{Id: testRegistryID}},
			},
			clientResponse: &mksclient.UpdateRegistriesV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.RegistriesIntegration{
					Registries: []mksclient.RegistryWithMeta{
						{
							Id:   testRegistryID,
							Name: "test-registry",
						},
					},
					Status: mksclient.RegistriesIntegrationStatusACTIVE,
				},
			},
		},
		{
			name: common.NameNotFound,
			requestBody: mksclient.RegistriesIntergrationUpdateBody{
				Registries: []mksclient.RegistriesIntergrationUpdateStruct{{Id: testRegistryID}},
			},
			clientResponse: &mksclient.UpdateRegistriesV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     http.StatusText(http.StatusNotFound),
				},
				JSON404: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
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
			name: common.NameBadRequest,
			requestBody: mksclient.RegistriesIntergrationUpdateBody{
				Registries: []mksclient.RegistriesIntergrationUpdateStruct{{Id: testRegistryID}},
			},
			clientResponse: &mksclient.UpdateRegistriesV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusBadRequest,
					Status:     http.StatusText(http.StatusBadRequest),
				},
				JSON400: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgBadRequest,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusBadRequest,
				Message:    common.MsgBadRequest,
			},
		},
		{
			name: common.NameInternalError,
			requestBody: mksclient.RegistriesIntergrationUpdateBody{
				Registries: []mksclient.RegistriesIntergrationUpdateStruct{{Id: testRegistryID}},
			},
			clientResponse: &mksclient.UpdateRegistriesV2Response{
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
			requestBody: mksclient.RegistriesIntergrationUpdateBody{
				Registries: []mksclient.RegistriesIntergrationUpdateStruct{{Id: testRegistryID}},
			},
			clientResponse: &mksclient.UpdateRegistriesV2Response{
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
			requestBody: mksclient.RegistriesIntergrationUpdateBody{
				Registries: []mksclient.RegistriesIntergrationUpdateStruct{{Id: testRegistryID}},
			},
			clientError: httpError,
			errExpected: httpError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mksClient := mksmock.NewMockClientWithResponsesInterface(t)
			mksClient.EXPECT().UpdateRegistriesV2WithResponse(
				mock.Anything, clusterID, test.requestBody,
			).Return(test.clientResponse, test.clientError)

			registries, err := Update(
				context.Background(),
				&v2.ServiceClient{MKSClient: mksClient},
				clusterID,
				test.requestBody.Registries,
			)

			if test.errExpected != nil {
				assert.Nil(t, registries)
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

			assert.Equal(t, test.clientResponse.JSON200.Registries, registries.Registries)
			assert.Equal(t, test.clientResponse.JSON200.Status, registries.Status)
		})
	}
}

func TestDelete(t *testing.T) {
	const (
		clusterID  = "test-cluster-id"
		registryID = testRegistryID
	)

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.DeleteRegistryV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.DeleteRegistryV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNoContent,
					Status:     http.StatusText(http.StatusNoContent),
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.DeleteRegistryV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     http.StatusText(http.StatusNotFound),
				},
				JSON404: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgRegistryNotFound,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusNotFound,
				Message:    common.MsgRegistryNotFound,
			},
		},
		{
			name: common.NameBadRequest,
			clientResponse: &mksclient.DeleteRegistryV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusBadRequest,
					Status:     http.StatusText(http.StatusBadRequest),
				},
				JSON400: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgBadRequest,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusBadRequest,
				Message:    common.MsgBadRequest,
			},
		},
		{
			name: common.NameInternalError,
			clientResponse: &mksclient.DeleteRegistryV2Response{
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
			clientResponse: &mksclient.DeleteRegistryV2Response{
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
			mksClient.EXPECT().DeleteRegistryV2WithResponse(mock.Anything, clusterID, registryID).Return(test.clientResponse, test.clientError)

			err := Delete(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID, registryID)

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

func TestDeleteAll(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.DeleteRegistriesV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.DeleteRegistriesV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNoContent,
					Status:     http.StatusText(http.StatusNoContent),
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.DeleteRegistriesV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     http.StatusText(http.StatusNotFound),
				},
				JSON404: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
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
			name: common.NameBadRequest,
			clientResponse: &mksclient.DeleteRegistriesV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusBadRequest,
					Status:     http.StatusText(http.StatusBadRequest),
				},
				JSON400: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgBadRequest,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusBadRequest,
				Message:    common.MsgBadRequest,
			},
		},
		{
			name: common.NameInternalError,
			clientResponse: &mksclient.DeleteRegistriesV2Response{
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
			clientResponse: &mksclient.DeleteRegistriesV2Response{
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
			mksClient.EXPECT().DeleteRegistriesV2WithResponse(mock.Anything, clusterID).Return(test.clientResponse, test.clientError)

			err := DeleteAll(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID)

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
