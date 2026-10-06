package mksclient

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMKSError_Status(t *testing.T) {
	const code = http.StatusInternalServerError

	err := MKSError{StatusCode: code}

	assert.Equal(t, code, err.StatusCode)
}

func TestMKSError_Error(t *testing.T) {
	const message = "message"

	err := MKSError{Message: message}

	assert.Equal(t, message, err.Error())
}

func TestGenericError_ToMKSError(t *testing.T) {
	const (
		code    = http.StatusInternalServerError
		message = "message"
	)

	var genericErr GenericError
	genericErr.Error.Message = message

	mksErr := genericErr.ToMKSError(code)

	assert.Equal(t, message, mksErr.Error())
	assert.Equal(t, code, mksErr.Status())
}

func TestGenericErrorNotFound_ToMKSError(t *testing.T) {
	const (
		code    = http.StatusNotFound
		message = "message"
	)

	var genericErr GenericNotFoundError
	genericErr.Error.Message = message

	mksErr := genericErr.ToMKSError(code)

	assert.Equal(t, message, mksErr.Error())
	assert.Equal(t, code, mksErr.Status())
}

func TestHandleAPIErrors(t *testing.T) {
	const (
		code                = http.StatusInternalServerError
		messageHTTP         = "message http"
		messageGenericError = "message generic"
	)

	t.Run("no generic errors", func(t *testing.T) {
		err := HandleAPIErrors(code, messageHTTP)
		require.Error(t, err)

		var mksErr *MKSError
		require.ErrorAs(t, err, &mksErr)

		assert.Equal(t, messageHTTP, mksErr.Error())
		assert.Equal(t, code, mksErr.Status())
	})

	t.Run("unknown server error", func(t *testing.T) {
		var (
			genericErr         *GenericError
			genericErrNotFound *GenericNotFoundError
		)

		err := HandleAPIErrors(code, messageHTTP, genericErr, genericErrNotFound)
		require.Error(t, err)

		var mksErr *MKSError
		require.ErrorAs(t, err, &mksErr)

		assert.Equal(t, messageHTTP, mksErr.Error())
		assert.Equal(t, code, mksErr.Status())
	})
	t.Run("get generic error", func(t *testing.T) {
		var (
			genericErr         *GenericError
			genericErrNotFound *GenericNotFoundError
		)

		genericErr = &GenericError{}
		genericErr.Error.Message = messageGenericError

		err := HandleAPIErrors(code, messageHTTP, genericErrNotFound, genericErr)
		require.Error(t, err)

		var mksErr *MKSError
		require.ErrorAs(t, err, &mksErr)

		assert.Equal(t, messageGenericError, mksErr.Error())
		assert.Equal(t, code, mksErr.Status())
	})
}
