package customresources

import (
	"crypto/md5"
	"encoding/base64"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	azureblob "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		assert.Len(t, m, 1)
		assert.Equal(t, "v1", *m["k1"])
	})
}

func TestAzblobToPulumiProperties(t *testing.T) {
	rg := "myrg"
	storageAccount := "mystorageaccount"
	container := "mycontainer"
	name := "myblob"
	blobUrl := "https://mystorageaccount.blob.core.windows.net/mycontainer/myblob"

	setup := func() azureblob.GetPropertiesResponse {
		md5Sum := md5.Sum([]byte("content"))
		return azureblob.GetPropertiesResponse{
			AccessTier:  ref(string(azureblob.AccessTierHot)),
			BlobType:    ref(azureblob.BlobTypeBlockBlob),
			ContentMD5:  md5Sum[:],
			ContentType: ref("text/plain"),
			Metadata: map[string]*string{
				"k1": ref("v1"),
			},
		}
	}

	t.Run("all properties", func(t *testing.T) {
		azblob := setup()
		pulumiBlob := azblobToPulumiProperties(name, rg, storageAccount, container, blobUrl, azblob)

		assert.Contains(t, pulumiBlob, accountName)
		assert.Equal(t, storageAccount, pulumiBlob[accountName])

		assert.Contains(t, pulumiBlob, containerName)
		assert.Equal(t, container, pulumiBlob[containerName])

		assert.Contains(t, pulumiBlob, blobName)
		assert.Equal(t, name, pulumiBlob[blobName])

		assert.Contains(t, pulumiBlob, accessTier)
		assert.Equal(t, string(azureblob.AccessTierHot), pulumiBlob[accessTier])

		assert.Contains(t, pulumiBlob, typeProp)
		assert.Equal(t, "Block", pulumiBlob[typeProp])

		assert.Contains(t, pulumiBlob, contentMd5)
		assert.Equal(t, base64.StdEncoding.EncodeToString(azblob.ContentMD5), pulumiBlob[contentMd5])

		assert.Contains(t, pulumiBlob, contentType)
		assert.Equal(t, "text/plain", pulumiBlob[contentType])

		require.Contains(t, pulumiBlob, metadata)
		m := pulumiBlob[metadata].(map[string]*string)
		assert.Equal(t, ref("v1"), m["k1"])

		assert.Contains(t, pulumiBlob, url)
		assert.Equal(t, blobUrl, pulumiBlob[url])
	})

	t.Run("no access tier", func(t *testing.T) {
		azblob := setup()

		azblob.AccessTier = nil
		pulumiBlob := azblobToPulumiProperties(name, rg, storageAccount, container, blobUrl, azblob)
		assert.NotContains(t, pulumiBlob, accessTier)

		azblob.AccessTier = ref("")
		pulumiBlob = azblobToPulumiProperties(name, rg, storageAccount, container, blobUrl, azblob)
		assert.NotContains(t, pulumiBlob, accessTier)

	})
}

func ref[T any](t T) *T { return &t }
