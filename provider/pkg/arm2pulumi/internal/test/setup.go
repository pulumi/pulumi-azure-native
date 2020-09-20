package test

import (
	"encoding/json"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func LoadSchema(t *testing.T) *schema.PackageSpec {
	var pkgSpec schema.PackageSpec
	_, filename, _, _ := runtime.Caller(0)
	currDir := path.Dir(filename)
	f, err := os.Open(filepath.Join(currDir, "../testdata/schema-full.json"))
	require.NoError(t, err)
	require.NoError(t, json.NewDecoder(f).Decode(&pkgSpec))
	return &pkgSpec
}

func LoadMetadata(t *testing.T) *provider.AzureAPIMetadata {
	var metadata provider.AzureAPIMetadata
	_, filename, _, _ := runtime.Caller(0)
	currDir := path.Dir(filename)
	f, err := os.Open(filepath.Join(currDir, "../testdata/metadata.json"))
	require.NoError(t, err)
	require.NoError(t, json.NewDecoder(f).Decode(&metadata))
	return &metadata
}
