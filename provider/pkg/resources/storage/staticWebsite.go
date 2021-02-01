package storage

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/tombuildsstuff/giovanni/storage/2018-11-09/blob/accounts"
	"net/http"
)

// Resource property names.
const (
	resourceGroupName = "resourceGroupName"
	accountName       = "accountName"
	indexDocument     = "indexDocument"
	error404Document  = "error404Document"
)

type StaticWebsite struct {
	accountsClient *storage.AccountsClient
	env            *azure.Environment
}

func (r *StaticWebsite) Schema() *schema.ResourceSpec {
	return &schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Enables the static website feature of a storage account.",
			Type:        "object",
			Properties: map[string]schema.PropertySpec{
				indexDocument: {
					Description: "The webpage that Azure Storage serves for requests to the root of a website or any subfolder. For example, 'index.html'. The value is case-sensitive.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
				error404Document: {
					Description: "The absolute path to a custom webpage that should be used when a request is made which does not correspond to an existing file.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
			},
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
				Description: "The webpage that Azure Storage serves for requests to the root of a website or any subfolder. For example, 'index.html'. The value is case-sensitive.",
				TypeSpec:    schema.TypeSpec{Type: "string"},
			},
			error404Document: {
				Description: "The absolute path to a custom webpage that should be used when a request is made which does not correspond to an existing file.",
				TypeSpec:    schema.TypeSpec{Type: "string"},
			},
		},
		RequiredInputs: []string{resourceGroupName, accountName},
	}
}

// newDataClient creates a new data-plane client given a property map with a resource group and a storage account name.
// Returns false as the second result if the storage account does not exist.
func (r *StaticWebsite) newDataClient(ctx context.Context, properties resource.PropertyMap) (*accounts.Client, bool, error) {
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

	props, err := r.accountsClient.ListKeys(ctx, rgName, accName, storage.Kerb)
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

	dataClient := accounts.NewWithEnvironment(*r.env)
	dataClient.Client.Authorizer = storageAuth
	return &dataClient, true, nil
}

func (r *StaticWebsite) CreateOrUpdate(ctx context.Context, properties resource.PropertyMap) (map[string]interface{}, error) {
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

	return properties.Mappable(), nil
}

func (r *StaticWebsite) Read(ctx context.Context, properties resource.PropertyMap) (map[string]interface{}, bool, error) {
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
	outputs[indexDocument] = resp.StorageServiceProperties.StaticWebsite.IndexDocument
	outputs[error404Document] = resp.StorageServiceProperties.StaticWebsite.ErrorDocument404Path
	return outputs, true, nil
}

func (r *StaticWebsite) Delete(ctx context.Context, properties resource.PropertyMap) error {
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

