package openapi

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadSimpleDefaultVersionFromReadme(t *testing.T) {
	f, err := os.Open("testdata/quota-readme.md")
	require.NoError(t, err)

	expectedVersion := ApiVersion("2025-03-01")
	versions, err := ReadDefaultVersionFromReadme(f)
	require.NoError(t, err)
	require.Len(t, versions, 1)
	require.Contains(t, versions, Service(""))
	require.Equal(t, expectedVersion, versions[""])
}
