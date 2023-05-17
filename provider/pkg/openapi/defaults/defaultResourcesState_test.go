package defaults

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GetDefaultResourceState returns expected values for a given resource type.
func TestNotFound(t *testing.T) {
	actual := GetDefaultResourceState("notapath")
	assert.Nil(t, actual)
}

func TestSkipDelete(t *testing.T) {
	actual := GetDefaultResourceState("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{tdeName}")
	expected := DefaultResourceState{
		SkipDelete: true,
	}
	assert.Equal(t, &expected, actual)
}

func TestPathNormalisation(t *testing.T) {
	// Same as TestSkipDelete, but with a different casing and parameter naming.
	actual := GetDefaultResourceState("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{transparentDataEncryptionName}")
	expected := DefaultResourceState{
		SkipDelete: true,
	}
	assert.Equal(t, &expected, actual)
}

func TestEmptyBody(t *testing.T) {
	actual := GetDefaultResourceState("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}")
	expected := DefaultResourceState{
		State: map[string]interface{}{},
	}
	assert.Equal(t, &expected, actual)
}

func TestNonEmptyCollection(t *testing.T) {
	actual := GetDefaultResourceState("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/currentbillingfeatures")
	expected := DefaultResourceState{
		State: map[string]interface{}{
			"currentBillingFeatures": []string{"Basic"},
			"dataVolumeCap":          map[string]string{},
		},
		HasNonEmptyCollections: true,
	}
	assert.Equal(t, &expected, actual)
}
