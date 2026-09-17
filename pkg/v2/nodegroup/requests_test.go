package nodegroup

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

func TestGet(t *testing.T) {
	const (
		clusterID   = "test-cluster-id"
		nodegroupID = "test-nodegroup-id"
	)

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.GetNodegroupV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.GetNodegroupV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.NodegroupResp{
					Nodegroup: mksclient.NodegroupDetailed{
						AutoscaleMaxNodes: common.Ptr(int64(10)),
						AutoscaleMinNodes: common.Ptr(int64(1)),
						Cidr:              common.Ptr("192.168.1.0/24"),
						CloudNodegroupConfig: &mksclient.CloudNodegroupConfig{
							AffinityPolicy: "anti-affinity",
							Cpus:           4,
							FlavorId:       "test-flavor",
							KeypairName:    "test-key",
							LocalVolume:    false,
							RamMb:          8192,
							VolumeGb:       100,
							VolumeType:     "standard",
						},
						ClusterId:                 clusterID,
						CreatedAt:                 time.Now(),
						DedicatedNodegroupConfig:  nil,
						EnableAutoscale:           true,
						Id:                        nodegroupID,
						InstallNvidiaDevicePlugin: true,
						Labels:                    map[string]string{"key": "value"},
						NodegroupType:             mksclient.NodegroupDetailedNodegroupTypeSTANDARD,
						Nodes: []mksclient.Node{{
							Id:          "node-1",
							Hostname:    "node-1",
							Ip:          "10.0.0.1",
							NodegroupId: nodegroupID,
							CreatedAt:   time.Now(),
						}},
						Preemptible: false,
						Segment:     "default-segment",
						Status:      mksclient.NodegroupDetailedStatusACTIVE,
						Taints: []mksclient.NodegroupTaint{{
							Key:    "key",
							Value:  "value",
							Effect: mksclient.NoSchedule,
						}},
						UpdatedAt: common.Ptr(time.Now()),
						UserData:  "user data",
					},
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.GetNodegroupV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     http.StatusText(http.StatusNotFound),
				},
				JSON404: &mksclient.GenericNotFoundError{
					Error: struct {
						Id      string `json:"id"` //nolint:revive // it's generated struct
						Message string `json:"message"`
					}{
						Id:      nodegroupID,
						Message: common.MsgNodegroupNotFound,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusNotFound,
				Message:    common.MsgNodegroupNotFound,
			},
		},
		{
			name: common.NameInternalError,
			clientResponse: &mksclient.GetNodegroupV2Response{
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
			clientResponse: &mksclient.GetNodegroupV2Response{
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
			mksClient.EXPECT().GetNodegroupV2WithResponse(mock.Anything, clusterID, nodegroupID).Return(test.clientResponse, test.clientError)

			nodegroup, err := Get(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID, nodegroupID)

			if test.errExpected != nil {
				assert.Nil(t, nodegroup)
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

			assert.Equal(t, test.clientResponse.JSON200.Nodegroup.Id, nodegroup.Id)
			assert.Equal(t, test.clientResponse.JSON200.Nodegroup.ClusterId, nodegroup.ClusterId)
		})
	}
}

func TestList(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.ListNodegroupsV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.ListNodegroupsV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.NodegroupList{
					Nodegroups: []mksclient.NodegroupListItem{
						{
							AutoscaleMaxNodes:       common.Ptr(int64(10)),
							AutoscaleMinNodes:       common.Ptr(int64(1)),
							AvailableAdditionalInfo: mksclient.NodegorupAdditionalInfo{UserData: true},
							Cidr:                    common.Ptr("192.168.1.0/24"),
							CloudNodegroupConfig: &mksclient.CloudNodegroupConfig{
								AffinityPolicy: "anti-affinity",
								Cpus:           4,
								FlavorId:       "test-flavor",
								KeypairName:    "test-key",
								LocalVolume:    false,
								RamMb:          8192,
								VolumeGb:       100,
								VolumeType:     "standard",
							},
							ClusterId:                 clusterID,
							CreatedAt:                 time.Now(),
							DedicatedNodegroupConfig:  nil,
							EnableAutoscale:           true,
							Id:                        "test-nodegroup-id-1",
							InstallNvidiaDevicePlugin: true,
							Labels:                    map[string]string{"key": "value"},
							NodegroupType:             mksclient.STANDARD,
							Nodes: []mksclient.Node{{
								Id:          "node-1",
								Hostname:    "node-1",
								Ip:          "10.0.0.1",
								NodegroupId: "test-nodegroup-id-1",
								CreatedAt:   time.Now(),
							}},
							Preemptible: false,
							Segment:     "default-segment",
							Status:      mksclient.NodegroupListItemStatusACTIVE,
							Taints: []mksclient.NodegroupTaint{{
								Key:    "key",
								Value:  "value",
								Effect: mksclient.NoSchedule,
							}},
							UpdatedAt: common.Ptr(time.Now()),
						},
					},
				},
			},
		},
		{
			name: common.NameEmptyNodegroups,
			clientResponse: &mksclient.ListNodegroupsV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.NodegroupList{},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.ListNodegroupsV2Response{
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
			clientResponse: &mksclient.ListNodegroupsV2Response{
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
			clientResponse: &mksclient.ListNodegroupsV2Response{
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
			clientResponse: &mksclient.ListNodegroupsV2Response{
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
			mksClient.EXPECT().ListNodegroupsV2WithResponse(mock.Anything, clusterID).Return(test.clientResponse, test.clientError)

			nodegroups, err := List(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID)

			if test.errExpected != nil {
				assert.Nil(t, nodegroups)
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

			if test.clientResponse.JSON200 != nil && test.clientResponse.JSON200.Nodegroups != nil {
				assert.Equal(t, test.clientResponse.JSON200.Nodegroups, nodegroups)
			} else {
				assert.Empty(t, nodegroups)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	const clusterID = "test-cluster-id"

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.CreateNodegroupsV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.CreateNodegroupsV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNoContent,
					Status:     http.StatusText(http.StatusNoContent),
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.CreateNodegroupsV2Response{
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
			clientResponse: &mksclient.CreateNodegroupsV2Response{
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
			clientResponse: &mksclient.CreateNodegroupsV2Response{
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
			clientResponse: &mksclient.CreateNodegroupsV2Response{
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
			clientResponse: &mksclient.CreateNodegroupsV2Response{
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
			mksClient.EXPECT().CreateNodegroupsV2WithResponse(mock.Anything, clusterID, mock.Anything).Return(test.clientResponse, test.clientError)

			err := Create(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID, nil)

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

func TestDelete(t *testing.T) {
	const (
		clusterID   = "test-cluster-id"
		nodegroupID = "test-nodegroup-id"
	)

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.DeleteNodegroupV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.DeleteNodegroupV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNoContent,
					Status:     http.StatusText(http.StatusNoContent),
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.DeleteNodegroupV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     http.StatusText(http.StatusNotFound),
				},
				JSON404: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgNodegroupNotFound,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusNotFound,
				Message:    common.MsgNodegroupNotFound,
			},
		},
		{
			name: common.NameBadRequest,
			clientResponse: &mksclient.DeleteNodegroupV2Response{
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
			clientResponse: &mksclient.DeleteNodegroupV2Response{
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
			clientResponse: &mksclient.DeleteNodegroupV2Response{
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
			clientResponse: &mksclient.DeleteNodegroupV2Response{
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
			mksClient.EXPECT().DeleteNodegroupV2WithResponse(mock.Anything, clusterID, nodegroupID).Return(test.clientResponse, test.clientError)

			err := Delete(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID, nodegroupID)

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

func TestResize(t *testing.T) {
	const (
		clusterID   = "test-cluster-id"
		nodegroupID = "test-nodegroup-id"
	)

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.ResizeNodegroupV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.ResizeNodegroupV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNoContent,
					Status:     http.StatusText(http.StatusNoContent),
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.ResizeNodegroupV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     http.StatusText(http.StatusNotFound),
				},
				JSON404: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgNodegroupNotFound,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusNotFound,
				Message:    common.MsgNodegroupNotFound,
			},
		},
		{
			name: common.NameBadRequest,
			clientResponse: &mksclient.ResizeNodegroupV2Response{
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
			clientResponse: &mksclient.ResizeNodegroupV2Response{
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
			clientResponse: &mksclient.ResizeNodegroupV2Response{
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
			clientResponse: &mksclient.ResizeNodegroupV2Response{
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
			mksClient.EXPECT().ResizeNodegroupV2WithResponse(mock.Anything, clusterID, nodegroupID, mock.Anything).Return(test.clientResponse, test.clientError)

			err := Resize(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID, nodegroupID, 0)

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

func TestUpdate(t *testing.T) {
	const (
		clusterID   = "test-cluster-id"
		nodegroupID = "test-nodegroup-id"
	)

	httpError := errors.New("error")

	tests := []struct {
		name           string
		clientResponse *mksclient.UpdateNodegroupV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.UpdateNodegroupV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNoContent,
					Status:     http.StatusText(http.StatusNoContent),
				},
			},
		},
		{
			name: common.NameNotFound,
			clientResponse: &mksclient.UpdateNodegroupV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     http.StatusText(http.StatusNotFound),
				},
				JSON404: &mksclient.GenericError{
					Error: struct {
						Message string `json:"message"`
					}{
						Message: common.MsgNodegroupNotFound,
					},
				},
			},
			errExpected: &mksclient.MKSError{
				StatusCode: http.StatusNotFound,
				Message:    common.MsgNodegroupNotFound,
			},
		},
		{
			name: common.NameBadRequest,
			clientResponse: &mksclient.UpdateNodegroupV2Response{
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
			clientResponse: &mksclient.UpdateNodegroupV2Response{
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
			clientResponse: &mksclient.UpdateNodegroupV2Response{
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
			clientResponse: &mksclient.UpdateNodegroupV2Response{
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
			mksClient.EXPECT().UpdateNodegroupV2WithResponse(mock.Anything, clusterID, nodegroupID, mock.Anything).Return(test.clientResponse, test.clientError)

			err := Update(context.Background(), &v2.ServiceClient{MKSClient: mksClient}, clusterID, nodegroupID, mksclient.NodegroupUpdateStruct{})

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
