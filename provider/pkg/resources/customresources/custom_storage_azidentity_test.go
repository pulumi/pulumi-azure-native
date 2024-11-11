package customresources

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestStorageAccountUrls(t *testing.T) {
	t.Run("empty account", func(t *testing.T) {
		_, err := getStorageAccountURL("", cloud.AzurePublic)
		assert.Error(t, err)
	})

	t.Run("invalid chars", func(t *testing.T) {
		_, err := getStorageAccountURL("not^^valid", cloud.AzurePublic)
		assert.Error(t, err)
	})

	t.Run("ok public", func(t *testing.T) {
		url, err := getStorageAccountURL("myaccount", cloud.AzurePublic)
		assert.NoError(t, err)
		assert.Equal(t, "https://myaccount.blob.core.windows.net", url)
	})

	t.Run("ok china", func(t *testing.T) {
		url, err := getStorageAccountURL("myaccount", cloud.AzureChina)
		assert.NoError(t, err)
		assert.Equal(t, "https://myaccount.blob.core.chinacloudapi.cn", url)
	})

	t.Run("ok usgov", func(t *testing.T) {
		url, err := getStorageAccountURL("myaccount", cloud.AzureGovernment)
		assert.NoError(t, err)
		assert.Equal(t, "https://myaccount.blob.core.usgovcloudapi.net", url)
	})
}

func TestBlobUrls(t *testing.T) {
	t.Run("ok public", func(t *testing.T) {
		properties := map[string]any{
			accountName:   "myaccount",
			containerName: "mycontainer",
			blobName:      "myblob",
		}
		url, err := getBlobURL(resource.NewPropertyMapFromMap(properties), cloud.AzurePublic)
		assert.NoError(t, err)
		assert.Equal(t, "https://myaccount.blob.core.windows.net/mycontainer/myblob", url)
	})

	t.Run("must have account", func(t *testing.T) {
		properties := map[string]any{
			containerName: "mycontainer",
			blobName:      "myblob",
		}
		_, err := getBlobURL(resource.NewPropertyMapFromMap(properties), cloud.AzurePublic)
		assert.Error(t, err)
	})

	t.Run("must have container", func(t *testing.T) {
		properties := map[string]any{
			accountName: "myaccount",
			blobName:    "myblob",
		}
		_, err := getBlobURL(resource.NewPropertyMapFromMap(properties), cloud.AzurePublic)
		assert.Error(t, err)
	})

	t.Run("must have blob", func(t *testing.T) {
		properties := map[string]any{
			accountName:   "myaccount",
			containerName: "mycontainer",
		}
		_, err := getBlobURL(resource.NewPropertyMapFromMap(properties), cloud.AzurePublic)
		assert.Error(t, err)
	})
}

func TestPopulateAzureBlobMetadata(t *testing.T) {
	t.Run("no metadata", func(t *testing.T) {
		properties := resource.PropertyMap{}
		m := populateAzureBlobMetadata(properties)
		assert.Empty(t, m)
	})

	t.Run("metadata is not an object", func(t *testing.T) {
		properties := resource.PropertyMap{
			metadata: resource.NewStringProperty("not an object"),
		}
		m := populateAzureBlobMetadata(properties)
		assert.Empty(t, m)
	})

	t.Run("metadata with mixed properties", func(t *testing.T) {
		properties := resource.PropertyMap{
			metadata: resource.NewObjectProperty(resource.PropertyMap{
				"k1": resource.NewStringProperty("v1"),
				"k2": resource.NewNumberProperty(42),
				"k3": resource.NewObjectProperty(resource.PropertyMap{}),
			}),
		}
		m := populateAzureBlobMetadata(properties)
		assert.Len(t, m, 2)
		assert.Equal(t, "v1", *m["k1"])
		assert.Equal(t, "42", *m["k2"])
	})
}
