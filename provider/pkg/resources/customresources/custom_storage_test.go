// Copyright 2016-2023, Pulumi Corporation.

package customresources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBlobIdProperties(t *testing.T) {
	id := "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Storage/storageAccounts/mysa/blobServices/default/containers/myc/blobs/folder%2Flog.txt"
	props, ok := parseBlobIdProperties(id)
	assert.True(t, ok)
	assert.Len(t, props, 5)
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", props[subscriptionId].StringValue())
	assert.Equal(t, "myrg", props[resourceGroupName].StringValue())
	assert.Equal(t, "mysa", props[accountName].StringValue())
	assert.Equal(t, "myc", props[containerName].StringValue())
	assert.Equal(t, "folder%2Flog.txt", props[blobName].StringValue())
}
