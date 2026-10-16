package cluster

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/selectel/mks-go/v2/internal/common"
	mks "github.com/selectel/mks-go/v2/pkg"
	"github.com/selectel/mks-go/v2/pkg/mksclient"
	mksmock "github.com/selectel/mks-go/v2/pkg/mksclient/mocks"
)

const (
	testClusterName = "test-cluster"
	testKubeVersion = "1.28.0"
	testKubeAPIIP   = "10.0.0.1"
)

func TestGet(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.GetClusterV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.GetClusterV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.ClusterResp{
					Cluster: &mksclient.ClusterDetailed{
						Id:                            clusterID,
						Name:                          testClusterName,
						KubeVersion:                   testKubeVersion,
						Basic:                         false,
						EnableAutorepair:              true,
						EnablePatchVersionAutoUpgrade: true,
						KubeApiIp:                     testKubeAPIIP,
						CreatedAt:                     time.Now(),
						CniType:                       mksclient.ClusterDetailedCniType("cilium"),
						NetworkType:                   mksclient.ClusterDetailedNetworkType("default"),
						Status:                        mksclient.ClusterDetailedStatus("active"),
					},
				},
			},
		},
		{
			name: common.NameInternalError,
			clientResponse: &mksclient.GetClusterV2Response{
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
			name: common.NameNotFound,
			clientResponse: &mksclient.GetClusterV2Response{
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
			name: common.NameUnknownStatus,
			clientResponse: &mksclient.GetClusterV2Response{
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
			mksClient.EXPECT().GetClusterV2WithResponse(mock.Anything, clusterID).Return(test.clientResponse, test.clientError)

			cluster, err := Get(context.Background(), &mks.ServiceClient{MKSClient: mksClient}, clusterID)

			if test.errExpected != nil {
				assert.Nil(t, cluster)
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

			assert.Equal(t, test.clientResponse.JSON200.Cluster, cluster)
		})
	}
}

func TestCreate(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.CreateClusterV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.CreateClusterV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusCreated,
					Status:     http.StatusText(http.StatusCreated),
				},
				JSON201: &mksclient.ClusterResp{
					Cluster: &mksclient.ClusterDetailed{
						Id:                            clusterID,
						Name:                          testClusterName,
						KubeVersion:                   testKubeVersion,
						Basic:                         false,
						EnableAutorepair:              true,
						EnablePatchVersionAutoUpgrade: true,
						KubeApiIp:                     testKubeAPIIP,
						CreatedAt:                     time.Now(),
						CniType:                       mksclient.ClusterDetailedCniType("cilium"),
						NetworkType:                   mksclient.ClusterDetailedNetworkType("default"),
						Status:                        mksclient.ClusterDetailedStatus("active"),
					},
				},
			},
		},
		{
			name: common.NameInternalError,
			clientResponse: &mksclient.CreateClusterV2Response{
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
			clientResponse: &mksclient.CreateClusterV2Response{
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
			name: common.NameAccessDenied,
			clientResponse: &mksclient.CreateClusterV2Response{
				Body: []byte(`{"error": {"message": "` + common.MsgAccessDenied + `"}}`),
				HTTPResponse: &http.Response{
					StatusCode: http.StatusForbidden,
					Status:     http.StatusText(http.StatusForbidden),
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusForbidden,
				Message:    http.StatusText(http.StatusForbidden) + ": " + common.MsgAccessDenied,
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
			mksClient.EXPECT().CreateClusterV2WithResponse(mock.Anything, mock.Anything).Return(test.clientResponse, test.clientError)

			cluster, err := Create(context.Background(), &mks.ServiceClient{MKSClient: mksClient}, nil)

			if test.errExpected != nil {
				assert.Nil(t, cluster)
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
			assert.Equal(t, test.clientResponse.JSON201.Cluster, cluster)
		})
	}
}

func TestUpdate(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.UpdateClusterV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.UpdateClusterV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.ClusterResp{
					Cluster: &mksclient.ClusterDetailed{
						Id:                            clusterID,
						Name:                          testClusterName,
						KubeVersion:                   testKubeVersion,
						Basic:                         false,
						EnableAutorepair:              true,
						EnablePatchVersionAutoUpgrade: true,
						KubeApiIp:                     testKubeAPIIP,
						CreatedAt:                     time.Now(),
						CniType:                       mksclient.ClusterDetailedCniType("cilium"),
						NetworkType:                   mksclient.ClusterDetailedNetworkType("default"),
						Status:                        mksclient.ClusterDetailedStatus("active"),
					},
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.UpdateClusterV2Response{
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
			clientResponse: &mksclient.UpdateClusterV2Response{
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
			clientResponse: &mksclient.UpdateClusterV2Response{
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
			mksClient.EXPECT().UpdateClusterV2WithResponse(mock.Anything, clusterID, mock.Anything).Return(test.clientResponse, test.clientError)

			cluster, err := Update(context.Background(), &mks.ServiceClient{MKSClient: mksClient}, clusterID, nil)

			if test.errExpected != nil {
				assert.Nil(t, cluster)
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
			assert.Equal(t, test.clientResponse.JSON200.Cluster, cluster)
		})
	}
}

func TestDelete(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.DeleteClusterV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.DeleteClusterV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNoContent,
					Status:     http.StatusText(http.StatusNoContent),
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.DeleteClusterV2Response{
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
			clientResponse: &mksclient.DeleteClusterV2Response{
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
			clientResponse: &mksclient.DeleteClusterV2Response{
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
			mksClient.EXPECT().DeleteClusterV2WithResponse(mock.Anything, clusterID).Return(test.clientResponse, test.clientError)

			err := Delete(context.Background(), &mks.ServiceClient{MKSClient: mksClient}, clusterID)

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

func TestGetKubeconfig(t *testing.T) {
	const clusterID = "test-cluster-id"

	testKubeconfig := []byte("kubeconfig")

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.GetClusterKubeconfigV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.GetClusterKubeconfigV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				Body: testKubeconfig,
			},
		},
		{
			name: common.NameAccessDenied,
			clientResponse: &mksclient.GetClusterKubeconfigV2Response{
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
			clientResponse: &mksclient.GetClusterKubeconfigV2Response{
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
			name: common.NameBadRequest,
			clientResponse: &mksclient.GetClusterKubeconfigV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusBadRequest,
					Status:     http.StatusText(http.StatusBadRequest),
				},
				Body: []byte(`{"error": {"message": "` + common.MsgBadRequest + `"}}`),
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusBadRequest,
				Message:    http.StatusText(http.StatusBadRequest) + ": " + common.MsgBadRequest,
			},
		},
		{
			name: common.NameConflict,
			clientResponse: &mksclient.GetClusterKubeconfigV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusConflict,
					Status:     http.StatusText(http.StatusConflict),
				},
				Body: []byte(`{"error": {"message": "` + common.MsgConflict + `"}}`),
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusConflict,
				Message:    http.StatusText(http.StatusConflict) + ": " + common.MsgConflict,
			},
		},
		{
			name: common.NameInternalError,
			clientResponse: &mksclient.GetClusterKubeconfigV2Response{
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
			clientResponse: &mksclient.GetClusterKubeconfigV2Response{
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
			mksClient.EXPECT().GetClusterKubeconfigV2WithResponse(mock.Anything, clusterID).Return(test.clientResponse, test.clientError)

			kubeconfig, err := GetKubeconfig(context.Background(), &mks.ServiceClient{MKSClient: mksClient}, clusterID)

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
			assert.Equal(t, test.clientResponse.Body, kubeconfig)
		})
	}
}

func TestRotateCerts(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.RotateClusterCertsV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.RotateClusterCertsV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNoContent,
					Status:     http.StatusText(http.StatusNoContent),
				},
			},
		},
		{
			name: common.NameAccessDenied,
			clientResponse: &mksclient.RotateClusterCertsV2Response{
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
			clientResponse: &mksclient.RotateClusterCertsV2Response{
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
			clientResponse: &mksclient.RotateClusterCertsV2Response{
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
			clientResponse: &mksclient.RotateClusterCertsV2Response{
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
			mksClient.EXPECT().RotateClusterCertsV2WithResponse(mock.Anything, clusterID).Return(test.clientResponse, test.clientError)

			err := RotateCerts(context.Background(), &mks.ServiceClient{MKSClient: mksClient}, clusterID)

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

func TestUpgradePatchVersion(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.UpgradePatchVersionV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.UpgradePatchVersionV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.ClusterResp{
					Cluster: &mksclient.ClusterDetailed{
						Id:                            clusterID,
						Name:                          testClusterName,
						KubeVersion:                   "1.28.1",
						Basic:                         false,
						EnableAutorepair:              true,
						EnablePatchVersionAutoUpgrade: true,
						KubeApiIp:                     testKubeAPIIP,
						CreatedAt:                     time.Now(),
						CniType:                       mksclient.ClusterDetailedCniType("cilium"),
						NetworkType:                   mksclient.ClusterDetailedNetworkType("default"),
						Status:                        mksclient.ClusterDetailedStatus("active"),
					},
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.UpgradePatchVersionV2Response{
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
			clientResponse: &mksclient.UpgradePatchVersionV2Response{
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
			clientResponse: &mksclient.UpgradePatchVersionV2Response{
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
			mksClient.EXPECT().UpgradePatchVersionV2WithResponse(mock.Anything, clusterID).Return(test.clientResponse, test.clientError)

			cluster, err := UpgradePatchVersion(context.Background(), &mks.ServiceClient{MKSClient: mksClient}, clusterID)

			if test.errExpected != nil {
				assert.Nil(t, cluster)
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
			assert.Equal(t, test.clientResponse.JSON200.Cluster, cluster)
		})
	}
}

func TestUpgradeMinorVersion(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.UpgradeMinorVersionV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.UpgradeMinorVersionV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.ClusterResp{
					Cluster: &mksclient.ClusterDetailed{
						Id:                            clusterID,
						Name:                          testClusterName,
						KubeVersion:                   "1.29.0",
						Basic:                         false,
						EnableAutorepair:              true,
						EnablePatchVersionAutoUpgrade: true,
						KubeApiIp:                     testKubeAPIIP,
						CreatedAt:                     time.Now(),
						CniType:                       mksclient.ClusterDetailedCniType("cilium"),
						NetworkType:                   mksclient.ClusterDetailedNetworkType("default"),
						Status:                        mksclient.ClusterDetailedStatus("active"),
					},
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.UpgradeMinorVersionV2Response{
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
			clientResponse: &mksclient.UpgradeMinorVersionV2Response{
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
			clientResponse: &mksclient.UpgradeMinorVersionV2Response{
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
			mksClient.EXPECT().UpgradeMinorVersionV2WithResponse(mock.Anything, clusterID).Return(test.clientResponse, test.clientError)

			cluster, err := UpgradeMinorVersion(context.Background(), &mks.ServiceClient{MKSClient: mksClient}, clusterID)

			if test.errExpected != nil {
				assert.Nil(t, cluster)
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
			assert.Equal(t, test.clientResponse.JSON200.Cluster, cluster)
		})
	}
}

func TestPatch(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")
	enablePatchVersionAutoUpgrade := false
	opts := &mksclient.ClusterUpdateStruct{
		EnablePatchVersionAutoUpgrade: &enablePatchVersionAutoUpgrade,
		KubernetesOptions: &mksclient.KubernetesOptions{
			FeatureGates: []string{"TTLAfterFinished"},
			AuditLogs:    mksclient.AuditLogs{Enabled: true},
		},
	}

	tests := []struct {
		name           string
		clientResponse *mksclient.PatchClusterV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.PatchClusterV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.ClusterResp{
					Cluster: &mksclient.ClusterDetailed{
						Id:                            clusterID,
						Name:                          testClusterName,
						KubeVersion:                   testKubeVersion,
						Basic:                         false,
						EnableAutorepair:              true,
						EnablePatchVersionAutoUpgrade: true,
						KubeApiIp:                     testKubeAPIIP,
						CreatedAt:                     time.Now(),
						CniType:                       mksclient.ClusterDetailedCniType("cilium"),
						NetworkType:                   mksclient.ClusterDetailedNetworkType("default"),
						Status:                        mksclient.ClusterDetailedStatus("active"),
					},
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.PatchClusterV2Response{
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
			name: common.NameBadRequest,
			clientResponse: &mksclient.PatchClusterV2Response{
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
			name: common.NameConflict,
			clientResponse: &mksclient.PatchClusterV2Response{
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
			clientResponse: &mksclient.PatchClusterV2Response{
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
			clientResponse: &mksclient.PatchClusterV2Response{
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
			mksClient.EXPECT().
				PatchClusterV2WithResponse(mock.Anything, clusterID, mksclient.PatchClusterV2JSONRequestBody{Cluster: opts}).
				Return(test.clientResponse, test.clientError)

			cluster, err := Patch(context.Background(), &mks.ServiceClient{MKSClient: mksClient}, clusterID, opts)

			if test.errExpected != nil {
				assert.Nil(t, cluster)
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
			assert.Equal(t, test.clientResponse.JSON200.Cluster, cluster)
		})
	}
}
