// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"
	"net/http"
	neturl "net/url"
	"strings"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/tombuildsstuff/giovanni/storage/2018-11-09/blob/blobs"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// newStorageAccountStaticWebsite creates a custom resource to mark Azure Storage Account as a static website.
func newStorageAccountStaticWebsite_azidentity(env *cloud.Configuration, creds azcore.TokenCredential) *CustomResource {
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

func getStorageAccountURL(saName string, env *cloud.Configuration) (string, error) {
	// TODO,tkappler use env.Services["storage"].Endpoint
	urlStr := fmt.Sprintf("https://%s.blob.core.windows.net", saName)
	_, err := neturl.Parse(urlStr)
	if err != nil {
		return "", errors.Errorf("invalid storage account URL: %v", err)
	}
	return urlStr, nil
}

func (r *staticWebsite_azidentity) newClient(properties resource.PropertyMap) (*azblob.Client, bool, error) {
	return newBlobClient(properties, r.env, r.creds)
}

func newBlobClient(properties resource.PropertyMap, env *cloud.Configuration, creds azcore.TokenCredential) (*azblob.Client, bool, error) {
	acc := properties[accountName]
	if !acc.HasValue() || !acc.IsString() {
		return nil, false, errors.Errorf("%q not found in resource state", accountName)
	}
	accName := acc.StringValue()

	saUrl, err := getStorageAccountURL(accName, env)
	if err != nil {
		return nil, false, err
	}

	client, err := azblob.NewClient(saUrl, creds, nil)
	// TODO,tkappler does this error if the account doesn't exist? If so, catch that here and return nil, false, nil
	return client, true, err
}

// // Resource property names.
// const (
// 	staticWebsitePath = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/staticWebsite"
// 	subscriptionId    = "subscriptionId"
// 	resourceGroupName = "resourceGroupName"
// 	accountName       = "accountName"
// 	containerName     = "containerName"
// 	indexDocument     = "indexDocument"
// 	error404Document  = "error404Document"
// )

type staticWebsite_azidentity struct {
	env   *cloud.Configuration
	creds azcore.TokenCredential
}

func (r *staticWebsite_azidentity) update(ctx context.Context, id string, properties, oldState resource.PropertyMap) (map[string]interface{}, error) {
	return r.createOrUpdate(ctx, id, properties)
}

func (r *staticWebsite_azidentity) createOrUpdate(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, error) {
	dataClient, found, err := r.newClient(properties)
	if err != nil {
		return nil, err
	}

	acc := properties[accountName].StringValue()
	if !found {
		return nil, errors.Errorf("storage account %q not found", acc)
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

	if _, err := dataClient.ServiceClient().SetProperties(ctx, siteProps); err != nil {
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
	dataClient, found, err := r.newClient(properties)
	if err != nil {
		return nil, false, err
	}

	if !found {
		// Storage Account not found, return empty state.
		return nil, false, nil
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
	dataClient, found, err := r.newClient(properties)
	if err != nil {
		return err
	}

	// Storage Account does not exist => delete is a no-op.
	if !found {
		return nil
	}

	acc := properties[accountName].StringValue()

	siteProps := &service.SetPropertiesOptions{
		StaticWebsite: &service.StaticWebsite{
			Enabled: pulumi.BoolRef(false),
		},
	}

	if _, err := dataClient.ServiceClient().SetProperties(ctx, siteProps); err != nil {
		return errors.Wrapf(err, "error updating storage account %q service properties", acc)
	}

	return nil
}

// newBlob_azidentity creates a custom resource for a Storage Blob.
func newBlob_azidentity(env *cloud.Configuration, creds azcore.TokenCredential) *CustomResource {
	r := blob_azidentity{
		env: env,
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
	env   *cloud.Configuration
}

func (r *blob_azidentity) newClient(properties resource.PropertyMap) (*azblob.Client, bool, error) {
	return newBlobClient(properties, r.env, r.creds)
}

func (r *blob_azidentity) create(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, error) {
	dataClient, found, err := r.newClient(properties)
	if err != nil {
		return nil, err
	}

	// If a storage account does not exist, we won't get an error but do get found=false.
	acc := properties[accountName].StringValue()
	if !found {
		return nil, errors.Errorf("storage account %q not found", acc)
	}

	container := properties[containerName].StringValue()
	name := properties[blobName].StringValue()

	metaData := make(map[string]string)
	if properties[metadata].HasValue() {
		for k, v := range properties[metadata].ObjectValue().Mappable() {
			metaData[k] = v.(string)
		}
	}

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

		input := blobs.PutAppendBlobInput{
			ContentType: &ct,
			MetaData:    metaData,
		}
		if _, err := dataClient.PutAppendBlob(ctx, acc, container, name, input); err != nil {
			return nil, err
		}
	case "block":
		input := blobs.PutBlockBlobInput{
			ContentType: &ct,
			MetaData:    metaData,
		}
		if properties[contentMd5].HasValue() {
			md5 := properties[contentMd5].StringValue()
			input.ContentMD5 = &md5
		}
		if properties[source].HasValue() {
			bytes, err := readAssetBytes(properties[source])
			if err != nil {
				return nil, err
			}
			input.Content = &bytes
		}
		if _, err := dataClient.PutBlockBlob(ctx, acc, container, name, input); err != nil {
			return nil, err
		}
	default:
		return nil, errors.Errorf("unsupported blob type: %q", blobType)
	}

	if properties[accessTier].HasValue() {
		tier := blobs.AccessTier(properties[accessTier].StringValue())
		if _, err := dataClient.SetTier(ctx, acc, container, name, tier); err != nil {
			return nil, errors.Wrapf(err, "updating access tier for blob %q (container %q / account %q)", name, container, acc)
		}
	}

	state, found, err := r.read(ctx, "", properties)
	if !found {
		return nil, errors.New("newly created blob is not found")
	}
	return state, err
}

// func (r *blob_azidentity) update(ctx context.Context, id string, properties, oldState resource.PropertyMap) (map[string]interface{}, error) {
// 	dataClient, found, err := r.newDataClient(ctx, properties)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// If a storage account does not exist, we won't get an error but do get found=false.
// 	acc := properties[accountName].StringValue()
// 	if !found {
// 		return nil, errors.Errorf("storage account %q not found", acc)
// 	}

// 	container := properties[containerName].StringValue()
// 	name := properties[blobName].StringValue()

// 	if properties[contentType].HasValue() {
// 		ct := properties[contentType].StringValue()
// 		input := blobs.SetPropertiesInput{
// 			ContentType: &ct,
// 		}
// 		if _, err := dataClient.SetProperties(ctx, acc, container, name, input); err != nil {
// 			return nil, errors.Wrapf(err, "updating properties for blob %q (container %q / account %q)", name, container, acc)
// 		}
// 	}
// 	if properties[metadata].HasValue() {
// 		metaData := make(map[string]string)
// 		for k, v := range properties[metadata].ObjectValue().Mappable() {
// 			metaData[k] = v.(string)
// 		}
// 		input := blobs.SetMetaDataInput{
// 			MetaData: metaData,
// 		}
// 		if _, err := dataClient.SetMetaData(ctx, acc, container, name, input); err != nil {
// 			return nil, errors.Wrapf(err, "updating metadata for blob %q (container %q / account %q)", name, container, acc)
// 		}
// 	}
// 	if properties[accessTier].HasValue() {
// 		tier := blobs.AccessTier(properties[accessTier].StringValue())
// 		if _, err := dataClient.SetTier(ctx, acc, container, name, tier); err != nil {
// 			return nil, errors.Wrapf(err, "updating access tier for blob %q (container %q / account %q)", name, container, acc)
// 		}
// 	}

// 	state, found, err := r.read(ctx, "", properties)
// 	if !found {
// 		return nil, errors.New("newly created blob is not found")
// 	}
// 	return state, err
// }

// func (r *blob) delete(ctx context.Context, id string, properties resource.PropertyMap) error {
// 	blobsClient, found, err := r.newDataClient(ctx, properties)
// 	if err != nil {
// 		return err
// 	}

// 	// Storage Account does not exist => delete is a no-op.
// 	if !found {
// 		return nil
// 	}

// 	acc := properties[accountName].StringValue()
// 	container := properties[containerName].StringValue()
// 	name := properties[blobName].StringValue()
// 	input := blobs.DeleteInput{
// 		DeleteSnapshots: true,
// 	}
// 	if _, err := blobsClient.Delete(ctx, acc, container, name, input); err != nil {
// 		return errors.Wrapf(err, "deleting blob %q (container %q / account %q)", name, container, acc)
// 	}

// 	return nil
// }

// func (r *blob) read(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, bool, error) {
// 	if len(properties) == 0 && id != "" {
// 		if idProps, ok := parseBlobIdProperties(id); ok {
// 			properties = idProps
// 		}
// 	}

// 	blobsClient, found, err := r.newDataClient(ctx, properties)
// 	if err != nil {
// 		return nil, false, err
// 	}

// 	// Storage Account does not exist
// 	if !found {
// 		return nil, false, nil
// 	}

// 	acc := properties[accountName].StringValue()
// 	container := properties[containerName].StringValue()
// 	name := properties[blobName].StringValue()
// 	input := blobs.GetPropertiesInput{}
// 	props, err := blobsClient.GetProperties(ctx, acc, container, name, input)
// 	if err != nil {
// 		if props.Response.Response != nil && props.Response.Response.StatusCode == http.StatusNotFound {
// 			return nil, false, nil
// 		}

// 		return nil, false, errors.Wrapf(err, "retrieving blob properties %q (container %q / account %q)", name, container, acc)
// 	}

// 	return sdkBlobToPulumiProperties(name,
// 			properties[resourceGroupName].StringValue(),
// 			acc,
// 			container,
// 			blobsClient.GetResourceID(acc, container, name),
// 			props),
// 		true, nil
// }

// func sdkBlobToPulumiProperties(name, rg, account, container, azureResourceId string, props blobs.GetPropertiesResult) map[string]any {
// 	result := map[string]interface{}{
// 		resourceGroupName: rg,
// 		accountName:       account,
// 		containerName:     container,
// 		blobName:          name,
// 		nameProp:          name,
// 		contentMd5:        props.ContentMD5,
// 		contentType:       props.ContentType,
// 		metadata:          props.MetaData,
// 		typeProp:          strings.TrimSuffix(string(props.BlobType), "Blob"),
// 		url:               azureResourceId,
// 	}
// 	// We can't serialize an empty string as that would be an invalid enum value.
// 	if props.AccessTier != "" {
// 		result[accessTier] = string(props.AccessTier)
// 	}
// 	return result
// }

// var blobIDPattern = regexp.MustCompile(`(?i)^/subscriptions/(.+)/resourceGroups/(.+)/providers/Microsoft.Storage/storageAccounts/(.+)/blobServices/default/containers/(.+)/blobs/(.+)$`)

// // parseBlobIdProperties parses an ID of a Blob resource to its identified properties.
// // For instance, it will convert
// // /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Storage/storageAccounts/mysa/blobServices/default/containers/myc/blobs/log.txt
// // to a map of
// // resourceGroupName=myrg,accountName=mysa,containerName=myc,blobName=log.txt.
// func parseBlobIdProperties(id string) (resource.PropertyMap, bool) {
// 	match := blobIDPattern.FindStringSubmatch(id)
// 	if len(match) != 6 {
// 		return nil, false
// 	}

// 	clientProperties := resource.PropertyMap{}
// 	clientProperties[resourceGroupName] = resource.NewStringProperty(match[2])
// 	clientProperties[accountName] = resource.NewStringProperty(match[3])
// 	clientProperties[containerName] = resource.NewStringProperty(match[4])
// 	clientProperties[blobName] = resource.NewStringProperty(match[5])
// 	return clientProperties, true
// }

// func newAccountAuthorizer(ctx context.Context, accountsClient *storage.AccountsClient, properties resource.PropertyMap) (*autorest.SharedKeyAuthorizer, bool, error) {
// 	rg := properties[resourceGroupName]
// 	if !rg.HasValue() || !rg.IsString() {
// 		return nil, false, errors.Errorf("%q not found in resource state", resourceGroupName)
// 	}

// 	acc := properties[accountName]
// 	if !acc.HasValue() || !acc.IsString() {
// 		return nil, false, errors.Errorf("%q not found in resource state", accountName)
// 	}

// 	accName := acc.StringValue()
// 	rgName := rg.StringValue()

// 	props, err := accountsClient.ListKeys(ctx, rgName, accName, storage.Kerb)
// 	if err != nil {
// 		if detailed, ok := err.(autorest.DetailedError); ok && detailed.StatusCode == http.StatusNotFound {
// 			return nil, false, nil
// 		}
// 		return nil, false, errors.Wrapf(err, "listing keys for storage account %q (resource group %q)", accName, rgName)
// 	}

// 	if props.Keys == nil || len(*props.Keys) == 0 || (*props.Keys)[0].Value == nil {
// 		return nil, false, errors.Errorf("keys were nil for storage account %q (resource group %q)", accName, rgName)
// 	}

// 	accountKey := (*props.Keys)[0].Value
// 	storageAuth, err := autorest.NewSharedKeyAuthorizer(accName, *accountKey, autorest.SharedKey)
// 	if err != nil {
// 		return nil, false, errors.Wrapf(err, "building shared key authorizer")
// 	}

// 	return storageAuth, true, nil
// }

// func readAssetBytes(s resource.PropertyValue) ([]byte, error) {
// 	switch {
// 	case s.IsAsset():
// 		asset := s.AssetValue()
// 		assetBytes, err := asset.Bytes()
// 		if err != nil {
// 			return nil, errors.Wrap(err, "reading asset bytes")
// 		}
// 		return assetBytes, nil
// 	case s.IsArchive():
// 		archive := s.ArchiveValue()
// 		archiveBytes, err := archive.Bytes(resource.ZIPArchive)
// 		if err != nil {
// 			return nil, errors.Wrap(err, "reading archive bytes")
// 		}
// 		return archiveBytes, nil
// 	default:
// 		return nil, errors.Errorf("unknown asset type %q", s.TypeString())
// 	}
// }
