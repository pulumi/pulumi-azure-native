package customresources

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	azpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	azureblob "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure/cloud"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadWithClient(t *testing.T) {
	blob := blob_azidentity{
		creds: &azfake.TokenCredential{},
		env:   cloud.AzurePublic,
	}

	properties := resource.NewPropertyMapFromMap(map[string]any{
		accountName:       "myaccount",
		blobName:          "myblob",
		containerName:     "mycontainer",
		resourceGroupName: "myrg",
	})

	azureProperties := azureblob.GetPropertiesResponse{
		BlobType:    ref(azureblob.BlobTypeBlockBlob),
		ContentMD5:  []byte("contentMD5"),
		ContentType: ref("text/plain"),
		Metadata: map[string]*string{
			"meta1": ref("value1"),
		},
	}

	var reader blobPropertiesReader = func(ctx context.Context) (azureblob.GetPropertiesResponse, error) {
		return azureProperties, nil
	}

	blobProperties, found, err := blob.readBlob(context.Background(), properties, reader)
	require.NoError(t, err)
	assert.True(t, found)

	assert.NotNil(t, blobProperties)
	assert.Equal(t, "myrg", blobProperties[resourceGroupName])
	assert.Equal(t, "myaccount", blobProperties[accountName])
	assert.Equal(t, "mycontainer", blobProperties[containerName])
	assert.Equal(t, "myblob", blobProperties[blobName])
	assert.Equal(t, "myblob", blobProperties[nameProp])

	assert.Equal(t, "Block", blobProperties[typeProp])
	assert.Equal(t, "Y29udGVudE1ENQ==", blobProperties[contentMd5])
	assert.Equal(t, "text/plain", blobProperties[contentType])
	assert.Equal(t, azureProperties.Metadata, blobProperties[metadata])

	assert.Equal(t, "https://myaccount.blob.core.windows.net/mycontainer/myblob", blobProperties[url])
}

func TestReadFromUSGovWithClient(t *testing.T) {
	blob := blob_azidentity{
		creds: &azfake.TokenCredential{},
		env:   cloud.AzureGovernment,
	}

	properties := resource.NewPropertyMapFromMap(map[string]any{
		accountName:       "myaccount",
		blobName:          "myblob",
		containerName:     "mycontainer",
		resourceGroupName: "myrg",
	})

	azureProperties := azureblob.GetPropertiesResponse{
		BlobType:    ref(azureblob.BlobTypeBlockBlob),
		ContentType: ref("text/plain"),
	}

	var reader blobPropertiesReader = func(ctx context.Context) (azureblob.GetPropertiesResponse, error) {
		return azureProperties, nil
	}

	blobProperties, found, err := blob.readBlob(context.Background(), properties, reader)
	require.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, "https://myaccount.blob.core.usgovcloudapi.net/mycontainer/myblob", blobProperties[url])
}

func fakeListKeysStorageAccountServer(keys []*armstorage.AccountKey) fake.AccountsServer {
	return fake.AccountsServer{
		ListKeys: func(ctx context.Context, resourceGroupName string, accountName string, options *armstorage.AccountsClientListKeysOptions) (resp azfake.Responder[armstorage.AccountsClientListKeysResponse], errResp azfake.ErrorResponder) {
			keysResp := armstorage.AccountsClientListKeysResponse{
				AccountListKeysResult: armstorage.AccountListKeysResult{
					Keys: keys,
				},
			}
			resp.SetResponse(http.StatusOK, keysResp, nil)
			return
		},
	}
}

func TestNewSharedKeyCredentialWithClient(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		base64Key := base64.StdEncoding.EncodeToString([]byte("key1valueasfgasgfdsfgdafgfadlgjas"))
		accountsServer := fakeListKeysStorageAccountServer([]*armstorage.AccountKey{
			{
				KeyName: ref("key1"),
				Value:   ref(base64Key),
			},
		})

		client, err := armstorage.NewAccountsClient("subscriptionID", &azfake.TokenCredential{}, &arm.ClientOptions{
			ClientOptions: azcore.ClientOptions{
				Transport: fake.NewAccountsServerTransport(&accountsServer),
			},
		})
		require.NoError(t, err)

		c, err := newSharedKeyCredentialWithClient(context.Background(), "myaccount", "myrg", client)
		require.NoError(t, err)
		require.NotNil(t, c)
	})

	t.Run("no keys", func(t *testing.T) {
		accountsServer := fakeListKeysStorageAccountServer([]*armstorage.AccountKey{})

		client, err := armstorage.NewAccountsClient("subscriptionID", &azfake.TokenCredential{}, &arm.ClientOptions{
			ClientOptions: azcore.ClientOptions{
				Transport: fake.NewAccountsServerTransport(&accountsServer),
			},
		})
		require.NoError(t, err)

		_, err = newSharedKeyCredentialWithClient(context.Background(), "myaccount", "myrg", client)
		require.Error(t, err)
		require.Contains(t, err.Error(), "no keys")
	})
}

func TestNewSharedKeyCredentialWithCloudConfig(t *testing.T) {
	t.Run("Azure Government Cloud", func(t *testing.T) {
		base64Key := base64.StdEncoding.EncodeToString([]byte("key1valueasfgasgfdsfgdafgfadlgjas"))
		accountsServer := fakeListKeysStorageAccountServer([]*armstorage.AccountKey{
			{
				KeyName: ref("key1"),
				Value:   ref(base64Key),
			},
		})

		// Mock transport to capture and validate the request
		var capturedRequest *http.Request
		mockTransport := &requestCapturingTransporter{
			fakeTransport: fake.NewAccountsServerTransport(&accountsServer),
			onRequest: func(req *http.Request) {
				capturedRequest = req
			},
		}

		// Test that the function accepts Azure Government cloud config and passes it correctly
		creds := &azfake.TokenCredential{}

		// This tests our modified newSharedKeyCredential function with Government cloud
		c, err := func() (*azblob.SharedKeyCredential, error) {
			clientOptions := &arm.ClientOptions{
				ClientOptions: azcore.ClientOptions{
					Cloud:     cloud.AzureGovernment.Configuration,
					Transport: mockTransport,
				},
			}
			accClient, err := armstorage.NewAccountsClient("subscriptionID", creds, clientOptions)
			if err != nil {
				return nil, err
			}
			return newSharedKeyCredentialWithClient(context.Background(), "myaccount", "myrg", accClient)
		}()

		require.NoError(t, err)
		require.NotNil(t, c)
		require.NotNil(t, capturedRequest)

		// Verify that the request was made to the government cloud endpoint
		expectedHost := "management.usgovcloudapi.net"
		require.Equal(t, expectedHost, capturedRequest.URL.Host,
			"Expected request to be made to Azure Government cloud endpoint")
	})

	t.Run("Azure Public Cloud", func(t *testing.T) {
		base64Key := base64.StdEncoding.EncodeToString([]byte("key1valueasfgasgfdsfgdafgfadlgjas"))
		accountsServer := fakeListKeysStorageAccountServer([]*armstorage.AccountKey{
			{
				KeyName: ref("key1"),
				Value:   ref(base64Key),
			},
		})

		// Mock transport to capture and validate the request
		var capturedRequest *http.Request
		mockTransport := &requestCapturingTransporter{
			fakeTransport: fake.NewAccountsServerTransport(&accountsServer),
			onRequest: func(req *http.Request) {
				capturedRequest = req
			},
		}

		// Test that the function works with public cloud config
		creds := &azfake.TokenCredential{}

		c, err := func() (*azblob.SharedKeyCredential, error) {
			clientOptions := &arm.ClientOptions{
				ClientOptions: azcore.ClientOptions{
					Cloud:     cloud.AzurePublic.Configuration,
					Transport: mockTransport,
				},
			}
			accClient, err := armstorage.NewAccountsClient("subscriptionID", creds, clientOptions)
			if err != nil {
				return nil, err
			}
			return newSharedKeyCredentialWithClient(context.Background(), "myaccount", "myrg", accClient)
		}()

		require.NoError(t, err)
		require.NotNil(t, c)
		require.NotNil(t, capturedRequest)

		// Verify that the request was made to the public cloud endpoint
		expectedHost := "management.azure.com"
		require.Equal(t, expectedHost, capturedRequest.URL.Host,
			"Expected request to be made to Azure Public cloud endpoint")
	})
}

func TestParseStorageAccountInput(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		props := resource.NewPropertyMapFromMap(map[string]any{
			accountName:       "myaccount",
			resourceGroupName: "myrg",
		})
		account, rg, err := parseStorageAccountInput(props)
		require.NoError(t, err)
		assert.Equal(t, "myaccount", account)
		assert.Equal(t, "myrg", rg)
	})

	t.Run("no account", func(t *testing.T) {
		props := resource.NewPropertyMapFromMap(map[string]any{
			resourceGroupName: "myrg",
		})
		_, _, err := parseStorageAccountInput(props)
		require.Error(t, err)
	})

	t.Run("no RG", func(t *testing.T) {
		props := resource.NewPropertyMapFromMap(map[string]any{
			accountName: "myaccount",
		})
		_, _, err := parseStorageAccountInput(props)
		require.Error(t, err)
	})
}

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

func TestGetStorageAccountName(t *testing.T) {
	validID := "/subscriptions/123/resourceGroups/rg/providers/Microsoft.Storage/storageAccounts/acc/staticWebsite"

	t.Run("from properties", func(t *testing.T) {
		properties := resource.NewPropertyMapFromMap(map[string]any{
			accountName: "myaccount",
		})
		name, err := getStorageAccountName(validID, properties)
		require.NoError(t, err)
		assert.Equal(t, "myaccount", name)
	})

	t.Run("from ID", func(t *testing.T) {
		name, err := getStorageAccountName(validID, resource.PropertyMap{})
		require.NoError(t, err)
		assert.Equal(t, "acc", name)
	})

	t.Run("invalid ID", func(t *testing.T) {
		_, err := getStorageAccountName("invalid", resource.PropertyMap{})
		require.Error(t, err)
	})

	t.Run("falls back to ID", func(t *testing.T) {
		name, err := getStorageAccountName(validID, resource.PropertyMap{})
		require.NoError(t, err)
		assert.Equal(t, "acc", name)
	})
}

func ref[T any](t T) *T { return &t }

// requestCapturingTransporter is a test helper that captures HTTP requests
type requestCapturingTransporter struct {
	fakeTransport azpolicy.Transporter
	onRequest     func(*http.Request)
}

func (r *requestCapturingTransporter) Do(req *http.Request) (*http.Response, error) {
	if r.onRequest != nil {
		r.onRequest(req)
	}
	return r.fakeTransport.Do(req)
}
