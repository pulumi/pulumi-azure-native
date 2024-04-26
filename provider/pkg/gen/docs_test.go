package gen

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanLookupExtraDocs(t *testing.T) {
	rootDir := t.TempDir()
	docsDir := filepath.Join(rootDir, "docs", "resources")
	require.NoError(t, os.MkdirAll(docsDir, 0755))

	for _, docs := range []struct {
		tok      string
		filename string
	}{
		{"azure-native:containerservice.ManagedCluster", "containerservice-ManagedCluster.md"},
		{"azure-native:sql.Server", "sql-Server.md"},
	} {
		require.NoError(t, os.WriteFile(filepath.Join(docsDir, docs.filename), []byte{'y'}, 0644))

		content := getAdditionalDocs(rootDir, docs.tok)
		require.NotNil(t, content, docs.tok)
		assert.Equal(t, "y", *content, docs.tok)
	}
}
