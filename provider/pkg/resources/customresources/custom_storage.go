// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/tombuildsstuff/giovanni/storage/2018-11-09/blob/accounts"
	"github.com/tombuildsstuff/giovanni/storage/2018-11-09/blob/blobs"
)

// newStorageAccountStaticWebsite creates a custom resource to mark Azure Storage Account as a static website.
func newStorageAccountStaticWebsite(env *azure.Environment, accountsClient *storage.AccountsClient) *CustomResource {
	r := staticWebsite{
		env:            env,
		accountsClient: accountsClient,
	}
	return &CustomResource{
		path:         staticWebsitePath,
		tok:          "azure-native:storage:StorageAccountStaticWebsite",
		Create:       r.createOrUpdate,
		Update:       r.update,
		Read:         r.read,
		Delete:       r.delete,
		LegacySchema: getStorageAccountStaticWebsiteSchema(),
		Meta:         getStorageAccountStaticWebsiteMetadata(),
	}
}

// Resource property names.
const (
	staticWebsitePath = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/staticWebsite"
	subscriptionId    = "subscriptionId"
	resourceGroupName = "resourceGroupName"
	accountName       = "accountName"
	containerName     = "containerName"
	indexDocument     = "indexDocument"
	error404Document  = "error404Document"
)

type staticWebsite struct {
	accountsClient *storage.AccountsClient
	env            *azure.Environment
}

// newDataClient creates a new data-plane client given a property map with a resource group and a storage account name.
// Returns false as the second result if the storage account does not exist.
func (r *staticWebsite) newDataClient(ctx context.Context, properties resource.PropertyMap) (*accounts.Client, bool, error) {
	storageAuth, exists, err := newAccountAuthorizer(ctx, r.accountsClient, properties)
	if err != nil || !exists {
		return nil, exists, err
	}

	dataClient := accounts.NewWithEnvironment(*r.env)
	dataClient.Client.Authorizer = storageAuth
	return &dataClient, true, nil
}

func (r *staticWebsite) update(ctx context.Context, id string, properties, oldState resource.PropertyMap) (map[string]interface{}, error) {
	return r.createOrUpdate(ctx, id, properties)
}

func (r *staticWebsite) createOrUpdate(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, error) {
	dataClient, found, err := r.newDataClient(ctx, properties)
	if err != nil {
		return nil, err
	}

	acc := properties[accountName].StringValue()
	if !found {
		return nil, errors.Errorf("storage account %q not found", acc)
	}

	siteProps := accounts.StorageServiceProperties{
		StaticWebsite: &accounts.StaticWebsite{
			Enabled: true,
		},
	}
	if p := properties[indexDocument]; p.HasValue() && p.IsString() {
		siteProps.StaticWebsite.IndexDocument = p.StringValue()
	}
	if p := properties[error404Document]; p.HasValue() && p.IsString() {
		siteProps.StaticWebsite.ErrorDocument404Path = p.StringValue()
	}

	if _, err := dataClient.SetServiceProperties(ctx, acc, siteProps); err != nil {
		return nil, errors.Wrapf(err, "error updating storage account %q service properties", acc)
	}

	outputs := properties.Mappable()
	outputs[containerName] = "$web"
	return outputs, nil
}

func (r *staticWebsite) read(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, bool, error) {
	dataClient, found, err := r.newDataClient(ctx, properties)
	if err != nil {
		return nil, false, err
	}

	// Storage Account not found, return empty state.
	if !found {
		return nil, false, nil
	}

	acc := properties[accountName].StringValue()
	resp, err := dataClient.GetServiceProperties(ctx, acc)
	if err != nil {
		return nil, false, errors.Wrapf(err, "error reading storage account %q service properties", acc)
	}

	// Static Website not enabled, return empty state.
	if !resp.StorageServiceProperties.StaticWebsite.Enabled {
		return nil, false, nil
	}

	outputs := properties.Mappable()
	outputs[containerName] = "$web"
	outputs[indexDocument] = resp.StorageServiceProperties.StaticWebsite.IndexDocument
	outputs[error404Document] = resp.StorageServiceProperties.StaticWebsite.ErrorDocument404Path
	return outputs, true, nil
}

func (r *staticWebsite) delete(ctx context.Context, id string, properties resource.PropertyMap) error {
	dataClient, found, err := r.newDataClient(ctx, properties)
	if err != nil {
		return err
	}

	// Storage Account does not exist => delete is a no-op.
	if !found {
		return nil
	}

	acc := properties[accountName].StringValue()
	if _, err := dataClient.SetServiceProperties(ctx, acc, accounts.StorageServiceProperties{
		StaticWebsite: &accounts.StaticWebsite{
			Enabled: false,
		},
	}); err != nil {
		return errors.Wrapf(err, "error updating storage account %q service properties", acc)
	}

	return nil
}

// Resource property names.
const (
	blobPath    = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/blobs/{blobName}"
	accessTier  = "accessTier"
	blobName    = "blobName"
	contentMd5  = "contentMd5"
	contentType = "contentType"
	metadata    = "metadata"
	nameProp    = "name"
	source      = "source"
	typeProp    = "type"
	url         = "url"
)

// newBlob creates a custom resource for a Storage Blob.
func newBlob(env *azure.Environment, accountsClient *storage.AccountsClient) *CustomResource {
	r := blob{
		env:            env,
		accountsClient: accountsClient,
	}
	return &CustomResource{
		path:         blobPath,
		tok:          "azure-native:storage:Blob",
		Create:       r.create,
		Update:       r.update,
		Delete:       r.delete,
		Read:         r.read,
		Types:        getBlobTypes(),
		LegacySchema: getBlobSchema(),
		Meta:         getBlobMetadata(),
	}
}

type blob struct {
	accountsClient *storage.AccountsClient
	env            *azure.Environment
}

// newDataClient creates a new blob client given a property map with a resource group and a storage account name.
// Returns false as the second result if the storage account does not exist (the error is nil in this case).
func (r *blob) newDataClient(ctx context.Context, properties resource.PropertyMap) (*blobs.Client, bool, error) {
	storageAuth, exists, err := newAccountAuthorizer(ctx, r.accountsClient, properties)
	if err != nil || !exists {
		return nil, exists, err
	}

	dataClient := blobs.NewWithEnvironment(*r.env)
	dataClient.Client.Authorizer = storageAuth
	return &dataClient, true, nil
}

func (r *blob) create(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, error) {
	dataClient, found, err := r.newDataClient(ctx, properties)
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

func (r *blob) update(ctx context.Context, id string, properties, oldState resource.PropertyMap) (map[string]interface{}, error) {
	dataClient, found, err := r.newDataClient(ctx, properties)
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

	if properties[contentType].HasValue() {
		ct := properties[contentType].StringValue()
		input := blobs.SetPropertiesInput{
			ContentType: &ct,
		}
		if _, err := dataClient.SetProperties(ctx, acc, container, name, input); err != nil {
			return nil, errors.Wrapf(err, "updating properties for blob %q (container %q / account %q)", name, container, acc)
		}
	}
	if properties[metadata].HasValue() {
		metaData := make(map[string]string)
		for k, v := range properties[metadata].ObjectValue().Mappable() {
			metaData[k] = v.(string)
		}
		input := blobs.SetMetaDataInput{
			MetaData: metaData,
		}
		if _, err := dataClient.SetMetaData(ctx, acc, container, name, input); err != nil {
			return nil, errors.Wrapf(err, "updating metadata for blob %q (container %q / account %q)", name, container, acc)
		}
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

func (r *blob) delete(ctx context.Context, id string, properties resource.PropertyMap) error {
	blobsClient, found, err := r.newDataClient(ctx, properties)
	if err != nil {
		return err
	}

	// Storage Account does not exist => delete is a no-op.
	if !found {
		return nil
	}

	acc := properties[accountName].StringValue()
	container := properties[containerName].StringValue()
	name := properties[blobName].StringValue()
	input := blobs.DeleteInput{
		DeleteSnapshots: true,
	}
	if _, err := blobsClient.Delete(ctx, acc, container, name, input); err != nil {
		return errors.Wrapf(err, "deleting blob %q (container %q / account %q)", name, container, acc)
	}

	return nil
}

func (r *blob) read(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, bool, error) {
	if len(properties) == 0 && id != "" {
		if idProps, ok := parseBlobIdProperties(id); ok {
			properties = idProps
		}
	}

	blobsClient, found, err := r.newDataClient(ctx, properties)
	if err != nil {
		return nil, false, err
	}

	// Storage Account does not exist
	if !found {
		return nil, false, nil
	}

	acc := properties[accountName].StringValue()
	container := properties[containerName].StringValue()
	name := properties[blobName].StringValue()
	input := blobs.GetPropertiesInput{}
	props, err := blobsClient.GetProperties(ctx, acc, container, name, input)
	if err != nil {
		if props.Response.Response != nil && props.Response.Response.StatusCode == http.StatusNotFound {
			return nil, false, nil
		}

		return nil, false, errors.Wrapf(err, "retrieving blob properties %q (container %q / account %q)", name, container, acc)
	}

	return sdkBlobToPulumiProperties(name,
			properties[resourceGroupName].StringValue(),
			acc,
			container,
			blobsClient.GetResourceID(acc, container, name),
			props),
		true, nil
}

func sdkBlobToPulumiProperties(name, rg, account, container, blobUrl string, props blobs.GetPropertiesResult) map[string]any {
	result := map[string]interface{}{
		resourceGroupName: rg,
		accountName:       account,
		containerName:     container,
		blobName:          name,
		nameProp:          name,
		contentMd5:        props.ContentMD5,
		contentType:       props.ContentType,
		metadata:          props.MetaData,
		typeProp:          strings.TrimSuffix(string(props.BlobType), "Blob"),
		url:               blobUrl,
	}
	// We can't serialize an empty string as that would be an invalid enum value.
	if props.AccessTier != "" {
		result[accessTier] = string(props.AccessTier)
	}
	return result
}

var blobIDPattern = regexp.MustCompile(`(?i)^/subscriptions/(.+)/resourceGroups/(.+)/providers/Microsoft.Storage/storageAccounts/(.+)/blobServices/default/containers/(.+)/blobs/(.+)$`)

// parseBlobIdProperties parses an ID of a Blob resource to its identified properties.
// For instance, it will convert
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Storage/storageAccounts/mysa/blobServices/default/containers/myc/blobs/log.txt
// to a map of
// resourceGroupName=myrg,accountName=mysa,containerName=myc,blobName=log.txt.
func parseBlobIdProperties(id string) (resource.PropertyMap, bool) {
	match := blobIDPattern.FindStringSubmatch(id)
	if len(match) != 6 {
		return nil, false
	}

	clientProperties := resource.PropertyMap{}
	clientProperties[resourceGroupName] = resource.NewStringProperty(match[2])
	clientProperties[accountName] = resource.NewStringProperty(match[3])
	clientProperties[containerName] = resource.NewStringProperty(match[4])
	clientProperties[blobName] = resource.NewStringProperty(match[5])
	return clientProperties, true
}

func newAccountAuthorizer(ctx context.Context, accountsClient *storage.AccountsClient, properties resource.PropertyMap) (*autorest.SharedKeyAuthorizer, bool, error) {
	rg := properties[resourceGroupName]
	if !rg.HasValue() || !rg.IsString() {
		return nil, false, errors.Errorf("%q not found in resource state", resourceGroupName)
	}

	acc := properties[accountName]
	if !acc.HasValue() || !acc.IsString() {
		return nil, false, errors.Errorf("%q not found in resource state", accountName)
	}

	accName := acc.StringValue()
	rgName := rg.StringValue()

	props, err := accountsClient.ListKeys(ctx, rgName, accName, storage.Kerb)
	if err != nil {
		if detailed, ok := err.(autorest.DetailedError); ok && detailed.StatusCode == http.StatusNotFound {
			return nil, false, nil
		}
		return nil, false, errors.Wrapf(err, "listing keys for storage account %q (resource group %q)", accName, rgName)
	}

	if props.Keys == nil || len(*props.Keys) == 0 || (*props.Keys)[0].Value == nil {
		return nil, false, errors.Errorf("keys were nil for storage account %q (resource group %q)", accName, rgName)
	}

	accountKey := (*props.Keys)[0].Value
	storageAuth, err := autorest.NewSharedKeyAuthorizer(accName, *accountKey, autorest.SharedKey)
	if err != nil {
		return nil, false, errors.Wrapf(err, "building shared key authorizer")
	}

	return storageAuth, true, nil
}

func readAssetBytes(s resource.PropertyValue) ([]byte, error) {
	switch {
	case s.IsAsset():
		asset := s.AssetValue()
		assetBytes, err := asset.Bytes()
		if err != nil {
			return nil, errors.Wrap(err, "reading asset bytes")
		}
		return assetBytes, nil
	case s.IsArchive():
		archive := s.ArchiveValue()
		archiveBytes, err := archive.Bytes(resource.ZIPArchive)
		if err != nil {
			return nil, errors.Wrap(err, "reading archive bytes")
		}
		return archiveBytes, nil
	default:
		return nil, errors.Errorf("unknown asset type %q", s.TypeString())
	}
}
