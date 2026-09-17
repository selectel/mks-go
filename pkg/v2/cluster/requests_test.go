package cluster

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	v2 "github.com/selectel/mks-go/pkg/v2"
	"github.com/selectel/mks-go/pkg/v2/internal/common"
	"github.com/selectel/mks-go/pkg/v2/mksclient"
	mksmock "github.com/selectel/mks-go/pkg/v2/mksclient/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	testClusterName        = "test-cluster"
	testKubeVersion        = "1.28.0"
	testKubeAPIIP          = "10.0.0.1"
	testAdditionalSoftware = "enabled"
	testSoftwareKey        = "nginx-ingress"
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
						AdditionalSoftware:            map[string]any{testSoftwareKey: testAdditionalSoftware},
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

			cluster, err := Get(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID)

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
						AdditionalSoftware:            map[string]any{testSoftwareKey: testAdditionalSoftware},
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
			name:        common.NameHTTPError,
			clientError: httpError,
			errExpected: httpError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mksClient := mksmock.NewMockClientWithResponsesInterface(t)
			mksClient.EXPECT().CreateClusterV2WithResponse(mock.Anything, mock.Anything).Return(test.clientResponse, test.clientError)

			cluster, err := Create(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, nil)

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
						AdditionalSoftware:            map[string]any{testSoftwareKey: testAdditionalSoftware},
					},
				},
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

			cluster, err := Update(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID, nil)

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

			err := Delete(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID)

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

			kubeconfig, err := GetKubeconfig(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID)

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

			err := RotateCerts(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID)

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
						AdditionalSoftware:            map[string]any{testSoftwareKey: testAdditionalSoftware},
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

			cluster, err := UpgradePatchVersion(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID)

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
						AdditionalSoftware:            map[string]any{testSoftwareKey: testAdditionalSoftware},
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

			cluster, err := UpgradeMinorVersion(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID)

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
