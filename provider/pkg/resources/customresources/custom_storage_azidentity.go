// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"
	"net/http"
	neturl "net/url"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/appendblob"
	azureblob "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// newStorageAccountStaticWebsite creates a custom resource to mark Azure Storage Account as a static website.
func storageAccountStaticWebsite_azidentity(env cloud.Configuration, creds azcore.TokenCredential) *CustomResource {
	r := staticWebsite_azidentity{
		env:   env,
		creds: creds,
	}
	return &CustomResource{
		path:   staticWebsitePath,
		tok:    "azure-native:storage:StorageAccountStaticWebsite",
		Create: r.createOrUpdate,
		Update: r.update,
		Read:   r.read,
		Delete: r.delete,
		LegacySchema: &schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "Enables the static website feature of a storage account.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					containerName: {
						Description: "The name of the container to upload blobs to.",
						TypeSpec:    schema.TypeSpec{Type: "string"},
					},
					indexDocument: {
						Description: "The webpage that Azure Storage serves for requests to the root of a website or any sub-folder. For example, 'index.html'. The value is case-sensitive.",
						TypeSpec:    schema.TypeSpec{Type: "string"},
					},
					error404Document: {
						Description: "The absolute path to a custom webpage that should be used when a request is made which does not correspond to an existing file.",
						TypeSpec:    schema.TypeSpec{Type: "string"},
					},
				},
				Required: []string{containerName},
			},
			InputProperties: map[string]schema.PropertySpec{
				resourceGroupName: {
					Description: "The name of the resource group within the user's subscription. The name is case insensitive.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
				accountName: {
					Description: "The name of the storage account within the specified resource group.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
				indexDocument: {
					Description: "The webpage that Azure Storage serves for requests to the root of a website or any sub-folder. For example, 'index.html'. The value is case-sensitive.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
				error404Document: {
					Description: "The absolute path to a custom webpage that should be used when a request is made which does not correspond to an existing file.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
			},
			RequiredInputs: []string{resourceGroupName, accountName},
		},
		Meta: &resources.AzureAPIResource{
			Path: staticWebsitePath,
			PutParameters: []resources.AzureAPIParameter{
				{Name: subscriptionId, Location: "path", IsRequired: true, Value: &resources.AzureAPIProperty{Type: "string"}},
				{Name: resourceGroupName, Location: "path", IsRequired: true, Value: &resources.AzureAPIProperty{Type: "string"}},
				{Name: accountName, Location: "path", IsRequired: true, Value: &resources.AzureAPIProperty{Type: "string"}},
				{
					Name:     "properties",
					Location: "body",
					Body: &resources.AzureAPIType{
						Properties: map[string]resources.AzureAPIProperty{
							indexDocument:    {Type: "string"},
							error404Document: {Type: "string"},
						},
					},
				},
			},
		},
	}
}

func getStorageAccountURL(saName string, env cloud.Configuration) (string, error) {
	baseUrl := "blob.core.windows.net"
	if env.ActiveDirectoryAuthorityHost == cloud.AzureChina.ActiveDirectoryAuthorityHost {
		baseUrl = "blob.core.chinacloudapi.cn"
	} else if env.ActiveDirectoryAuthorityHost == cloud.AzureGovernment.ActiveDirectoryAuthorityHost {
		baseUrl = "blob.core.usgovcloudapi.net"
	}

	urlStr := fmt.Sprintf("https://%s.%s", saName, baseUrl)
	_, err := neturl.Parse(urlStr)
	if err != nil || saName == "" {
		return "", errors.Errorf("invalid storage account URL: %v", err)
	}
	return urlStr, nil
}

func (r *staticWebsite_azidentity) newStorageAccountClient(properties resource.PropertyMap) (*azblob.Client, error) {
	return newStorageAccountClient(properties, r.env, r.creds)
}

func newStorageAccountClient(properties resource.PropertyMap, env cloud.Configuration, creds azcore.TokenCredential) (*azblob.Client, error) {
	acc := properties[accountName]
	if !acc.HasValue() || !acc.IsString() {
		return nil, errors.Errorf("%q not found in resource state", accountName)
	}
	accName := acc.StringValue()

	saUrl, err := getStorageAccountURL(accName, env)
	if err != nil {
		return nil, err
	}

	return azblob.NewClient(saUrl, creds, nil)
}

type staticWebsite_azidentity struct {
	env   cloud.Configuration
	creds azcore.TokenCredential
}

func (r *staticWebsite_azidentity) update(ctx context.Context, id string, properties, oldState resource.PropertyMap) (map[string]interface{}, error) {
	return r.createOrUpdate(ctx, id, properties)
}

func (r *staticWebsite_azidentity) createOrUpdate(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, error) {
	accountClient, err := r.newStorageAccountClient(properties)
	if err != nil {
		return nil, err
	}

	siteProps := &service.SetPropertiesOptions{
		StaticWebsite: &service.StaticWebsite{
			Enabled: pulumi.BoolRef(true),
		},
	}
	if p := properties[indexDocument]; p.HasValue() && p.IsString() {
		siteProps.StaticWebsite.IndexDocument = pulumi.StringRef(p.StringValue())
	}
	if p := properties[error404Document]; p.HasValue() && p.IsString() {
		siteProps.StaticWebsite.ErrorDocument404Path = pulumi.StringRef(p.StringValue())
	}

	if _, err := accountClient.ServiceClient().SetProperties(ctx, siteProps); err != nil {
		acc := properties[accountName].StringValue()
		return nil, errors.Wrapf(err, "error updating storage account %q service properties", acc)
	}

	outputs := properties.Mappable()
	outputs[containerName] = "$web"
	return outputs, nil
}

func is404StorageError(err error) bool {
	var respErr *azcore.ResponseError
	if errors.As(err, &respErr) {
		return respErr.StatusCode == http.StatusNotFound
	}
	return false
}

func (r *staticWebsite_azidentity) read(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, bool, error) {
	dataClient, err := r.newStorageAccountClient(properties)
	if err != nil {
		return nil, false, err
	}

	// Storage Account does not exist => return empty state.
	_, err = dataClient.ServiceClient().GetAccountInfo(ctx, nil)
	if err != nil {
		if is404StorageError(err) {
			return nil, false, nil
		}
		return nil, false, err
	}

	accountProps, err := dataClient.ServiceClient().GetProperties(ctx, nil)
	if err != nil {
		return nil, false, errors.Wrapf(err, "error reading storage account %q service properties", properties[accountName].StringValue())
	}

	if accountProps.StaticWebsite == nil || (accountProps.StaticWebsite.Enabled == nil || !*accountProps.StaticWebsite.Enabled) {
		// Static Website not enabled, return empty state.
		return nil, false, nil
	}

	outputs := properties.Mappable()
	outputs[containerName] = "$web"
	outputs[indexDocument] = accountProps.StaticWebsite.IndexDocument
	outputs[error404Document] = accountProps.StaticWebsite.ErrorDocument404Path
	return outputs, true, nil
}

func (r *staticWebsite_azidentity) delete(ctx context.Context, id string, properties resource.PropertyMap) error {
	accountClient, err := r.newStorageAccountClient(properties)
	if err != nil {
		return err
	}

	// Storage Account does not exist => delete is a no-op.
	_, err = accountClient.ServiceClient().GetAccountInfo(ctx, nil)
	if err != nil {
		if is404StorageError(err) {
			return nil
		}
		return err
	}

	acc := properties[accountName].StringValue()

	siteProps := &service.SetPropertiesOptions{
		StaticWebsite: &service.StaticWebsite{
			Enabled: pulumi.BoolRef(false),
		},
	}

	if _, err := accountClient.ServiceClient().SetProperties(ctx, siteProps); err != nil {
		return errors.Wrapf(err, "error updating storage account %q service properties", acc)
	}

	return nil
}

// newBlob_azidentity creates a custom resource for a Storage Blob.
func newBlob_azidentity(env cloud.Configuration, creds azcore.TokenCredential) *CustomResource {
	r := blob_azidentity{
		creds: creds,
		env:   env,
	}
	return &CustomResource{
		path:   blobPath,
		tok:    "azure-native:storage:Blob",
		Create: r.create,
		Update: r.update,
		Delete: r.delete,
		Read:   r.read,
		Types: map[string]schema.ComplexTypeSpec{
			"azure-native:storage:BlobAccessTier": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "The access tier of a storage blob.",
					Type:        "string",
				},
				Enum: []schema.EnumValueSpec{
					{
						Value:       "Hot",
						Description: "Optimized for storing data that is accessed frequently.",
					},
					{
						Value:       "Cool",
						Description: "Optimized for storing data that is infrequently accessed and stored for at least 30 days.",
					},
					{
						Value:       "Archive",
						Description: "Optimized for storing data that is rarely accessed and stored for at least 180 days with flexible latency requirements, on the order of hours.",
					},
				},
			},
			"azure-native:storage:BlobType": {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "The type of a storage blob to be created.",
					Type:        "string",
				},
				Enum: []schema.EnumValueSpec{
					{
						Value:       "Block",
						Description: "Block blobs store text and binary data. Block blobs are made up of blocks of data that can be managed individually.",
					},
					{
						Value:       "Append",
						Description: "Append blobs are made up of blocks like block blobs, but are optimized for append operations.",
					},
				},
			},
		},
		LegacySchema: &schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "Manages a Blob within a Storage Container. For the supported combinations of properties and features please see [here](https://learn.microsoft.com/en-us/azure/storage/blobs/storage-feature-support-in-storage-accounts).",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					accessTier: {
						Description: "The access tier of the storage blob. Only supported for standard storage accounts, not premium.",
						TypeSpec:    schema.TypeSpec{Ref: "#/types/azure-native:storage:BlobAccessTier"},
					},
					contentMd5: {
						Description: "The MD5 sum of the blob contents.",
						TypeSpec:    schema.TypeSpec{Type: "string"},
					},
					contentType: {
						Description: "The content type of the storage blob.",
						TypeSpec:    schema.TypeSpec{Type: "string"},
					},
					metadata: {
						Description: "A map of custom blob metadata.",
						TypeSpec:    schema.TypeSpec{Type: "object", AdditionalProperties: &schema.TypeSpec{Type: "string"}},
					},
					nameProp: {
						Description: "The name of the storage blob.",
						TypeSpec:    schema.TypeSpec{Type: "string"},
					},
					typeProp: {
						Description: "The type of the storage blob to be created.",
						TypeSpec:    schema.TypeSpec{Ref: "#/types/azure-native:storage:BlobType"},
					},
					url: {
						Description: "The URL of the blob.",
						TypeSpec:    schema.TypeSpec{Type: "string"},
					},
				},
				Required: []string{metadata, nameProp, typeProp, url},
			},
			InputProperties: map[string]schema.PropertySpec{
				accessTier: {
					Description: "The access tier of the storage blob. Only supported for standard storage accounts, not premium.",
					TypeSpec:    schema.TypeSpec{Ref: "#/types/azure-native:storage:BlobAccessTier"},
				},
				accountName: {
					Description:          "Specifies the storage account in which to create the storage container.",
					TypeSpec:             schema.TypeSpec{Type: "string"},
					WillReplaceOnChanges: true,
				},
				blobName: {
					Description:          "The name of the storage blob. Must be unique within the storage container the blob is located. If this property is not specified it will be set to the name of the resource.",
					TypeSpec:             schema.TypeSpec{Type: "string"},
					WillReplaceOnChanges: true,
				},
				containerName: {
					Description:          "The name of the storage container in which this blob should be created.",
					TypeSpec:             schema.TypeSpec{Type: "string"},
					WillReplaceOnChanges: true,
				},
				contentMd5: {
					Description:          "The MD5 sum of the blob contents. Cannot be defined if blob type is Append.",
					TypeSpec:             schema.TypeSpec{Type: "string"},
					WillReplaceOnChanges: true,
				},
				contentType: {
					Description: "The content type of the storage blob. Defaults to `application/octet-stream`.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
				metadata: {
					Description: "A map of custom blob metadata.",
					TypeSpec:    schema.TypeSpec{Type: "object", AdditionalProperties: &schema.TypeSpec{Type: "string"}},
				},
				resourceGroupName: {
					Description:          "The name of the resource group within the user's subscription.",
					TypeSpec:             schema.TypeSpec{Type: "string"},
					WillReplaceOnChanges: true,
				},
				source: {
					Description:          "An asset to copy to the blob contents. This field cannot be specified for Append blobs.",
					TypeSpec:             schema.TypeSpec{Ref: "pulumi.json#/Asset"},
					WillReplaceOnChanges: true,
				},
				typeProp: {
					Description:          "The type of the storage blob to be created. Defaults to 'Block'.",
					TypeSpec:             schema.TypeSpec{Ref: "#/types/azure-native:storage:BlobType"},
					Default:              "Block",
					WillReplaceOnChanges: true,
				},
			},
			RequiredInputs: []string{resourceGroupName, accountName, containerName},
		},
		Meta: &resources.AzureAPIResource{
			Path: blobPath,
			PutParameters: []resources.AzureAPIParameter{
				{Name: subscriptionId, Location: "path", IsRequired: true, Value: &resources.AzureAPIProperty{Type: "string"}},
				{Name: resourceGroupName, Location: "path", IsRequired: true, Value: &resources.AzureAPIProperty{Type: "string"}},
				{Name: accountName, Location: "path", IsRequired: true, Value: &resources.AzureAPIProperty{Type: "string"}},
				{Name: containerName, Location: "path", IsRequired: true, Value: &resources.AzureAPIProperty{Type: "string"}},
				{Name: blobName, Location: "path", IsRequired: true, Value: &resources.AzureAPIProperty{Type: "string", AutoName: "copy"}},
				{
					Name:     "properties",
					Location: "body",
					Body: &resources.AzureAPIType{
						Properties: map[string]resources.AzureAPIProperty{
							accessTier:  {Type: "string"},
							contentMd5:  {Type: "string", ForceNew: true},
							contentType: {Type: "string"},
							metadata:    {Type: "object", AdditionalProperties: &resources.AzureAPIProperty{Type: "string"}},
							source:      {Ref: "pulumi.json#/Asset", ForceNew: true},
							typeProp:    {Type: "string", ForceNew: true},
						},
						RequiredProperties: []string{resourceGroupName, accountName, containerName, blobName, typeProp},
					},
				},
			},
		},
	}
}

type blob_azidentity struct {
	creds azcore.TokenCredential
	env   cloud.Configuration
}

func (r *blob_azidentity) newBlobClient(properties resource.PropertyMap) (*azureblob.Client, error) {
	return newBlobClient(properties, r.env, r.creds)
}

func (r *blob_azidentity) newAppendBlobClient(properties resource.PropertyMap) (*appendblob.Client, error) {
	return newAppendBlobClient(properties, r.env, r.creds)
}

func (r *blob_azidentity) newBlockBlobClient(properties resource.PropertyMap) (*blockblob.Client, error) {
	return newBlockBlobClient(properties, r.env, r.creds)
}

func newBlobClient(properties resource.PropertyMap, env cloud.Configuration, creds azcore.TokenCredential) (*azureblob.Client, error) {
	blobUrl, err := getBlobURL(properties, env)
	if err != nil {
		return nil, err
	}
	return azureblob.NewClient(blobUrl, creds, nil)
}

func newAppendBlobClient(properties resource.PropertyMap, env cloud.Configuration, creds azcore.TokenCredential) (*appendblob.Client, error) {
	blobUrl, err := getBlobURL(properties, env)
	if err != nil {
		return nil, err
	}
	return appendblob.NewClient(blobUrl, creds, nil)
}

func newBlockBlobClient(properties resource.PropertyMap, env cloud.Configuration, creds azcore.TokenCredential) (*blockblob.Client, error) {
	blobUrl, err := getBlobURL(properties, env)
	if err != nil {
		return nil, err
	}
	return blockblob.NewClient(blobUrl, creds, nil)
}

func getBlobURL(properties resource.PropertyMap, env cloud.Configuration) (string, error) {
	acc := properties[accountName]
	if !acc.HasValue() || !acc.IsString() {
		return "", errors.Errorf("%q not found in resource state", accountName)
	}
	container := properties[containerName]
	if !container.HasValue() || !container.IsString() {
		return "", errors.Errorf("%q not found in resource state", containerName)
	}
	blobName := properties[blobName]
	if !blobName.HasValue() || !blobName.IsString() {
		return "", errors.Errorf("%q not found in resource state", blobName)
	}

	saUrl, err := getStorageAccountURL(acc.StringValue(), env)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s/%s", saUrl, container.StringValue(), blobName.StringValue()), nil
}

func populateAzureBlobMetadata(properties resource.PropertyMap) map[string]*string {
	metaData := make(map[string]*string)
	if properties[metadata].HasValue() {
		metadataRaw := properties[metadata]
		if !metadataRaw.IsObject() {
			logging.V(5).Infof("Warning: property 'metadata' is not an object, skipping: %v", metadataRaw)
			return metaData
		}
		for k, v := range metadataRaw.ObjectValue() {
			if v.IsString() {
				metaData[string(k)] = pulumi.StringRef(v.StringValue())
			} else if v.IsNumber() || v.IsBool() {
				metaData[string(k)] = pulumi.StringRef(fmt.Sprintf("%v", v.NumberValue()))
			} else {
				logging.V(5).Infof("Warning: metadata for key '%q' is not a string, skipping: %v", k, v)
			}
		}
	}
	return metaData
}

func (r *blob_azidentity) create(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, error) {
	blobClient, err := r.newBlobClient(properties)
	if err != nil {
		return nil, err
	}

	acc := properties[accountName].StringValue()
	container := properties[containerName].StringValue()
	name := properties[blobName].StringValue()
	metaData := populateAzureBlobMetadata(properties)

	var ct string
	if properties[contentType].HasValue() {
		ct = properties[contentType].StringValue()
	}

	blobType := strings.ToLower(properties[typeProp].StringValue())
	switch blobType {
	case "append":
		if properties[source].HasValue() {
			return nil, errors.New("a source cannot be specified for an Append blob")
		}

		appendClient, err := r.newAppendBlobClient(properties)
		if err != nil {
			return nil, err
		}

		if _, err := appendClient.Create(ctx, &appendblob.CreateOptions{
			Metadata: metaData,
			HTTPHeaders: &azureblob.HTTPHeaders{
				BlobContentType: &ct,
			},
		}); err != nil {
			return nil, err
		}
	case "block":
		blockClient, err := r.newBlockBlobClient(properties)
		if err != nil {
			return nil, err
		}

		opts := blockblob.UploadBufferOptions{
			HTTPHeaders: &azureblob.HTTPHeaders{
				BlobContentType: &ct,
			},
			Metadata: metaData,
		}

		if properties[contentMd5].HasValue() {
			md5 := properties[contentMd5].StringValue()
			opts.HTTPHeaders.BlobContentMD5 = []byte(md5)
		}

		input := []byte{}
		if properties[source].HasValue() {
			bytes, err := readAssetBytes(properties[source])
			if err != nil {
				return nil, err
			}
			input = bytes
		}

		if _, err := blockClient.UploadBuffer(ctx, input, &opts); err != nil {
			return nil, err
		}
	default:
		return nil, errors.Errorf("unsupported blob type: %q", blobType)
	}

	if properties[accessTier].HasValue() {
		tier := azureblob.AccessTier(properties[accessTier].StringValue())
		if _, err := blobClient.SetTier(ctx, tier, nil); err != nil {
			return nil, errors.Wrapf(err, "updating access tier for blob %q (container %q / account %q)", name, container, acc)
		}
	}

	state, found, err := r.read(ctx, "", properties)
	if !found {
		return nil, errors.New("newly created blob is not found")
	}
	return state, err
}

func (r *blob_azidentity) update(ctx context.Context, id string, properties, oldState resource.PropertyMap) (map[string]interface{}, error) {
	dataClient, err := r.newBlobClient(properties)
	if err != nil {
		return nil, err
	}

	acc := properties[accountName].StringValue()
	container := properties[containerName].StringValue()
	name := properties[blobName].StringValue()

	if properties[contentType].HasValue() {
		ct := properties[contentType].StringValue()
		if _, err := dataClient.SetHTTPHeaders(ctx, azureblob.HTTPHeaders{
			BlobContentType: &ct,
		}, nil); err != nil {
			return nil, errors.Wrapf(err, "updating properties for blob %q (container %q / account %q)", name, container, acc)
		}
	}

	meta := populateAzureBlobMetadata(properties)
	if len(meta) > 0 {
		if _, err := dataClient.SetMetadata(ctx, meta, nil); err != nil {
			return nil, errors.Wrapf(err, "updating metadata for blob %q (container %q / account %q)", name, container, acc)
		}
	}

	if properties[accessTier].HasValue() {
		tier := azureblob.AccessTier(properties[accessTier].StringValue())
		if _, err := dataClient.SetTier(ctx, tier, nil); err != nil {
			return nil, errors.Wrapf(err, "updating access tier for blob %q (container %q / account %q)", name, container, acc)
		}
	}

	state, found, err := r.read(ctx, "", properties)
	if !found {
		return nil, errors.New("newly created blob is not found")
	}
	return state, err
}

func (r *blob_azidentity) delete(ctx context.Context, id string, properties resource.PropertyMap) error {
	blobsClient, err := r.newBlobClient(properties)
	if err != nil {
		return err
	}

	ref := func(s azureblob.DeleteSnapshotsOptionType) *azureblob.DeleteSnapshotsOptionType { return &s }

	if _, err := blobsClient.Delete(ctx, &azureblob.DeleteOptions{
		DeleteSnapshots: ref(azureblob.DeleteSnapshotsOptionTypeInclude),
	}); err != nil {
		// Storage Account does not exist => delete is a no-op.
		if is404StorageError(err) {
			return nil
		}

		acc := properties[accountName].StringValue()
		container := properties[containerName].StringValue()
		name := properties[blobName].StringValue()
		return errors.Wrapf(err, "deleting blob %q (container %q / account %q)", name, container, acc)
	}

	return nil
}

func (r *blob_azidentity) read(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, bool, error) {
	if len(properties) == 0 && id != "" {
		if idProps, ok := parseBlobIdProperties(id); ok {
			properties = idProps
		}
	}

	blobsClient, err := r.newBlobClient(properties)
	if err != nil {
		return nil, false, err
	}

	acc := properties[accountName].StringValue()
	container := properties[containerName].StringValue()
	name := properties[blobName].StringValue()
	props, err := blobsClient.GetProperties(ctx, nil)
	if err != nil {
		if is404StorageError(err) {
			return nil, false, nil
		}

		return nil, false, errors.Wrapf(err, "retrieving blob properties %q (container %q / account %q)", name, container, acc)
	}

	blobURL, err := getBlobURL(properties, r.env)
	if err != nil {
		return nil, false, err
	}

	return azblobToPulumiProperties(name,
			properties[resourceGroupName].StringValue(),
			acc,
			container,
			// The previous implementation's id from `blobsClient.GetResourceID(acc, container, name)`
			// was also the blob URL, so the id in state should be identical. See
			// https://github.com/tombuildsstuff/giovanni/blob/v0.15.1/storage/2018-11-09/blob/blobs/resource_id.go#L13
			blobURL,
			props),
		true, nil
}

func azblobToPulumiProperties(name, rg, account, container, azureResourceId string, props azureblob.GetPropertiesResponse) map[string]any {
	result := map[string]any{
		resourceGroupName: rg,
		accountName:       account,
		containerName:     container,
		blobName:          name,
		nameProp:          name,
		contentMd5:        fmt.Sprintf("%x", props.ContentMD5), // the binary hash needs to be hex-encoded
		contentType:       *props.ContentType,
		metadata:          props.Metadata,
		typeProp:          strings.TrimSuffix(string(*props.BlobType), "Blob"),
		url:               azureResourceId,
	}
	// We can't serialize an empty string as that would be an invalid enum value.
	at := *props.AccessTier
	if at != "" {
		result[accessTier] = at
	}
	return result
}
