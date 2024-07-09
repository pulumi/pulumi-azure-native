// Copyright 2016-2023, Pulumi Corporation.

package customresources

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tombuildsstuff/giovanni/storage/2018-11-09/blob/blobs"
)

func TestParseBlobIdProperties(t *testing.T) {
	id := "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Storage/storageAccounts/mysa/blobServices/default/containers/myc/blobs/folder%2Flog.txt"
	props, ok := parseBlobIdProperties(id)
	assert.Equal(t, ok, true)
	assert.Equal(t, len(props), 4)
	assert.Equal(t, props["resourceGroupName"].StringValue(), "myrg")
	assert.Equal(t, props["accountName"].StringValue(), "mysa")
	assert.Equal(t, props["containerName"].StringValue(), "myc")
	assert.Equal(t, props["blobName"].StringValue(), "folder%2Flog.txt")
}

func TestReadBlobProperties(t *testing.T) {
	t.Run("Access tier set", func(t *testing.T) {
		r := blobs.GetPropertiesResult{
			AccessTier: "Cool",
		}
		p := sdkBlobToPulumiProperties("myblob", "rg", "account", "container", "id", r)
		require.Contains(t, p, accessTier)
		assert.Equal(t, p[accessTier], "Cool")
	})

	t.Run("Access tier not set", func(t *testing.T) {
		r := blobs.GetPropertiesResult{
			AccessTier: "",
		}
		p := sdkBlobToPulumiProperties("myblob", "rg", "account", "container", "id", r)
		require.NotContains(t, p, accessTier)
	})

	t.Run("Access tier empty", func(t *testing.T) {
		r := blobs.GetPropertiesResult{
			AccessTier: "",
		}
		p := sdkBlobToPulumiProperties("myblob", "rg", "account", "container", "id", r)
		require.NotContains(t, p, accessTier)
	})
}
