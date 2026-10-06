package kubeversion

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

func TestList(t *testing.T) {
	httpError := errors.New("error")

	versions := []mksclient.KubeVersionInfo{
		{
			Version:   common.Ptr("v1.27.0"),
			IsDefault: common.Ptr(false),
		},
		{
			Version:   common.Ptr("v1.28.0"),
			IsDefault: common.Ptr(true),
		},
	}

	tests := []struct {
		name           string
		clientResponse *mksclient.ListKubeVersionsV2Response
		clientError    error
		errExpected    error
	}{
		{
			name: common.NameSuccess,
			clientResponse: &mksclient.ListKubeVersionsV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.KubeVersionsList{
					KubeVersions: &versions,
				},
			},
		},
		{
			name: common.NameEmptyKubeVersions,
			clientResponse: &mksclient.ListKubeVersionsV2Response{
				HTTPResponse: &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				},
				JSON200: &mksclient.KubeVersionsList{},
			},
		},
		{
			name: common.NameInternalError,
			clientResponse: &mksclient.ListKubeVersionsV2Response{
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
			clientResponse: &mksclient.ListKubeVersionsV2Response{
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
			mksClient.EXPECT().ListKubeVersionsV2WithResponse(mock.Anything).Return(test.clientResponse, test.clientError)

			kubeVersions, err := List(context.Background(), &mks.ServiceClient{MKSClient: mksClient})

			if test.errExpected != nil {
				assert.Nil(t, kubeVersions)
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

			if test.clientResponse.JSON200 != nil && test.clientResponse.JSON200.KubeVersions != nil {
				assert.Equal(t, *test.clientResponse.JSON200.KubeVersions, kubeVersions)
			} else {
				assert.Empty(t, kubeVersions)
			}
		})
	}
}
