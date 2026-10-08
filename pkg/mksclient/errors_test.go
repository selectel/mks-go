package mksclient

import (
	"net/http"
	"strconv"
	"strings"
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

func TestHandleAPIErrorsWithBody(t *testing.T) {
	const (
		messageGenericError = "message generic"
		messageBody         = "reason from body"
	)

	genericBody := []byte(`{"error": {"message": "` + messageBody + `"}}`)

	typedErr := &GenericError{}
	typedErr.Error.Message = messageGenericError

	type testCase struct {
		name        string
		code        int
		body        []byte
		errors      []APIError
		msgExpected string
	}

	tests := make([]testCase, 0, 18)
	tests = append(tests, []testCase{
		{
			name:        "typed error wins over body",
			code:        http.StatusInternalServerError,
			body:        genericBody,
			errors:      []APIError{(*GenericNotFoundError)(nil), typedErr},
			msgExpected: messageGenericError,
		},
		{
			name:        "plain text body",
			code:        http.StatusForbidden,
			body:        []byte("  access\n\tdenied  "),
			msgExpected: http.StatusText(http.StatusForbidden) + ": access denied",
		},
		{
			name:        "html body",
			code:        http.StatusBadGateway,
			body:        []byte("<html>\n<body>Bad Gateway</body>\n</html>\n"),
			msgExpected: http.StatusText(http.StatusBadGateway) + ": <html> <body>Bad Gateway</body> </html>",
		},
		{
			name:        "json body without message",
			code:        http.StatusForbidden,
			body:        []byte(`{"error": {}}`),
			msgExpected: http.StatusText(http.StatusForbidden) + `: {"error": {}}`,
		},
		{
			name:        "empty body",
			code:        http.StatusServiceUnavailable,
			msgExpected: http.StatusText(http.StatusServiceUnavailable),
		},
		{
			name:        "whitespace body",
			code:        http.StatusServiceUnavailable,
			body:        []byte(" \n\t "),
			msgExpected: http.StatusText(http.StatusServiceUnavailable),
		},
		{
			name:        "long body is cut",
			code:        http.StatusInternalServerError,
			body:        []byte(strings.Repeat("a", maxErrorBodyLen+100)),
			msgExpected: http.StatusText(http.StatusInternalServerError) + ": " + strings.Repeat("a", maxErrorBodyLen) + "...",
		},
		{
			name:        "long body is cut on a rune boundary",
			code:        http.StatusInternalServerError,
			body:        []byte("a" + strings.Repeat("я", maxErrorBodyLen)),
			msgExpected: http.StatusText(http.StatusInternalServerError) + ": a" + strings.Repeat("я", (maxErrorBodyLen-1)/2) + "...",
		},
	}...)

	for _, code := range []int{
		http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound,
		http.StatusConflict, http.StatusUnprocessableEntity, http.StatusTooManyRequests,
		http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable,
	} {
		tests = append(tests, testCase{
			name:        "generic json body " + strconv.Itoa(code),
			code:        code,
			body:        genericBody,
			msgExpected: http.StatusText(code) + ": " + messageBody,
		})
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := HandleAPIErrorsWithBody(test.code, http.StatusText(test.code), test.body, test.errors...)
			require.Error(t, err)

			var mksErr *MKSError
			require.ErrorAs(t, err, &mksErr)

			assert.Equal(t, test.msgExpected, mksErr.Error())
			assert.Equal(t, test.code, mksErr.Status())
		})
	}
}
