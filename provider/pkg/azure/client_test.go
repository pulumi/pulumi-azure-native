package azure

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleError(t *testing.T) {
	resp := http.Response{
		Status: "404 Not Found",
		Body:   io.NopCloser(strings.NewReader("body")),
	}

	t.Run("normal error", func(t *testing.T) {
		origErr := errors.New("original error")
		err := handleResponseError(origErr, &resp)
		assert.Equal(t, origErr, err)
	})

	t.Run("non-JSON response", func(t *testing.T) {
		origErr := errors.New("autorest/azure: error response cannot be parsed")
		err := handleResponseError(origErr, &resp)
		require.NotNil(t, err)
		assert.Contains(t, err.Error(), "unexpected response from Azure: 'body'")
		assert.Contains(t, err.Error(), "404")
	})
}
