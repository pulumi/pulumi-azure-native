package openapi

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadSimpleDefaultVersionFromReadme(t *testing.T) {
	f, err := os.Open("testdata/quota-readme.md")
	require.NoError(t, err)

	version, err := readDefaultVersionFromReadme(f)
	require.NoError(t, err)
	require.Equal(t, "2025-03-01", version)
}
