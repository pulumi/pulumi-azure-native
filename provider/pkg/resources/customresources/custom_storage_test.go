// Copyright 2016-2023, Pulumi Corporation.

package customresources

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
