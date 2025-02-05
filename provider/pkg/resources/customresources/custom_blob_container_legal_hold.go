package customresources

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	. "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/versionLookup"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

const lhPath = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/legalHold"
const tagsProp = "tags"
const allowProtectedAppendWritesAllProp = "allowProtectedAppendWritesAll"

var legalHoldProperties = map[string]schema.PropertySpec{
	resourceGroupName: {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "Name of the resource group that contains the storage account.",
	},
	accountName: {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "Name of the Storage Account.",
	},
	containerName: {
		TypeSpec:    schema.TypeSpec{Type: "string"},
		Description: "Name of the Blob Container.",
	},
	tagsProp: {
		TypeSpec:    schema.TypeSpec{Type: "array", Items: &schema.TypeSpec{Type: "string"}},
		Description: "List of legal hold tags. Each tag should be 3 to 23 alphanumeric characters and is normalized to lower case at SRP.",
	},
	allowProtectedAppendWritesAllProp: {
		TypeSpec:    schema.TypeSpec{Type: "boolean"},
		Description: "When enabled, new blocks can be written to both 'Append and Bock Blobs' while maintaining legal hold protection and compliance. Only new blocks can be added and any existing blocks cannot be modified or deleted.",
	},
}

func blobContainerLegalHold(azureClient azure.AzureClient) *CustomResource {
	apiVersion, ok := versionLookup.GetDefaultApiVersionForResource("Storage", "BlobContainer")
	if !ok {
		// default as of 2024-02-16
		apiVersion = "2022-09-01"
		logging.V(3).Infof("Warning: could not find default API version for storage:blobContainer. Using %s", apiVersion)
	}

	return &CustomResource{
		tok:  "azure-native:storage:BlobContainerLegalHold",
		path: lhPath,
		LegacySchema: &schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: ".",
				Type:        "object",
				Properties:  legalHoldProperties,
			},
			InputProperties: legalHoldProperties,
			RequiredInputs:  []string{resourceGroupName, accountName, containerName, tagsProp},
		},
		Meta: &AzureAPIResource{
			Path: lhPath,
			PutParameters: []AzureAPIParameter{
				{Name: subscriptionId, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
				{Name: resourceGroupName, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
				{Name: accountName, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
				{Name: containerName, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
				{
					Name:     "properties",
					Location: "body",
					Body: &AzureAPIType{
						Properties: map[string]AzureAPIProperty{
							tagsProp: {
								Type: "array",
								Items: &AzureAPIProperty{
									Type: "string",
								},
							},
							allowProtectedAppendWritesAllProp: {
								Type: "boolean",
							},
						},
						RequiredProperties: []string{resourceGroupName, accountName, containerName, tagsProp},
					},
				},
			},
		},
		Create: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]any, error) {
			createPath := strings.TrimSuffix(id, "/legalHold") + "/setLegalHold"

			p := properties.Mappable()
			tagsVal, ok := p[tagsProp]
			if !ok {
				return nil, errors.New("missing required property 'tags'")
			}
			tags, ok := tagsVal.([]any)
			if !ok {
				return nil, errors.New("property 'tags' must be an array")
			}
			if len(tags) == 0 {
				return nil, errors.New("property 'tags' must not be empty")
			}

			body := map[string]any{
				tagsProp: tags,
			}

			if allowProtectedAppendWritesAllVal, ok := p[allowProtectedAppendWritesAllProp]; ok {
				if allowProtectedAppendWritesAll, ok := allowProtectedAppendWritesAllVal.(bool); ok {
					body[allowProtectedAppendWritesAllProp] = allowProtectedAppendWritesAll
				}
			}

			return post(ctx, createPath, body, apiVersion, azureClient)
		},
		Read: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]any, bool, error) {
			containerId := strings.TrimSuffix(id, "/legalHold")
			container, err := azureClient.Get(ctx, containerId, "2022-09-01", nil)
			if err != nil {
				if azure.IsNotFound(err) {
					return nil, false, nil
				}
				return nil, false, err
			}

			legalHoldProp, ok := util.GetInnerMap(container.(map[string]any), "properties", "legalHold")
			if !ok {
				return nil, false, nil
			}
			if hasLegalHold, ok := legalHoldProp["hasLegalHold"]; ok && hasLegalHold.(bool) {
				// Annoyingly, when reading the blob container, the legal hold data is returned in
				// a different format than what 'set' and 'clear' accept and return, so we convert.
				// From: { "hasLegalHold": true,
				//         "tags": [ { "tag": "tag1" }, { "tag": "tag2" } ],
				//         "protectedAppendWritesHistory": { "allowProtectedAppendWritesAll": true } }
				// To:   { "tags": [ "tag1", "tag2" ],
				//         "allowProtectedAppendWritesAll": true }
				result := map[string]any{}
				if tags, ok := legalHoldProp[tagsProp]; ok {
					tagsResult := []string{}
					if tagsArr, ok := tags.([]any); ok {
						for _, tagObj := range tagsArr {
							if tagObjMap, ok := tagObj.(map[string]any); ok {
								if tag, ok := tagObjMap["tag"]; ok {
									tagsResult = append(tagsResult, tag.(string))
								}
							}
						}
					}
					result[tagsProp] = tagsResult
				}
				if hist, ok := util.GetInnerMap(legalHoldProp, "protectedAppendWritesHistory"); ok {
					if allow, ok := hist["allowProtectedAppendWritesAll"]; ok {
						result[allowProtectedAppendWritesAllProp] = allow.(bool)
					}
				}
				return result, true, nil
			}

			return nil, false, nil
		},
		Update: func(ctx context.Context, id string, properties, oldState resource.PropertyMap) (map[string]any, error) {
			basePath := strings.TrimSuffix(id, "/legalHold")
			setPath := basePath + "/setLegalHold"
			clearPath := basePath + "/clearLegalHold"

			input := properties.Mappable()
			olds := oldState.Mappable()

			// We need to handle added and removed tags separately.
			tags, err := readTags(input)
			if err != nil {
				return nil, err
			}
			oldTags, err := readTags(olds)
			if err != nil {
				return nil, err
			}

			addedTags := tags.Subtract(oldTags)
			removedTags := oldTags.Subtract(tags)

			// not set -> set     => true
			// set     -> not set => false
			// set     -> set     => true  or no-op
			// not set -> not set => false or no-op
			allowProtected := false
			if allowProtectedAppendWritesAllVal, ok := input[allowProtectedAppendWritesAllProp]; ok {
				if allowProtectedAppendWritesAll, ok := allowProtectedAppendWritesAllVal.(bool); ok {
					allowProtected = allowProtectedAppendWritesAll
				}
			}
			allowProtectedPrev := false
			if allowProtectedAppendWritesAllValPrev, ok := olds[allowProtectedAppendWritesAllProp]; ok {
				if allowProtectedAppendWritesAllPrev, ok := allowProtectedAppendWritesAllValPrev.(bool); ok {
					allowProtectedPrev = allowProtectedAppendWritesAllPrev
				}
			}
			allowProtectedChanged := allowProtected != allowProtectedPrev

			if addedTags.Any() || allowProtectedChanged {
				body := map[string]any{
					tagsProp:                          addedTags.SortedValues(),
					allowProtectedAppendWritesAllProp: allowProtected,
				}
				_, err = post(ctx, setPath, body, apiVersion, azureClient)
				if err != nil {
					return nil, err
				}
			}
			if removedTags.Any() {
				body := map[string]any{
					tagsProp: removedTags.SortedValues(),
				}
				_, err = post(ctx, clearPath, body, apiVersion, azureClient)
				if err != nil {
					return nil, err
				}
			}

			return input, nil
		},
		Delete: func(ctx context.Context, id string, properties, state resource.PropertyMap) error {
			path := strings.TrimSuffix(id, "/legalHold") + "/clearLegalHold"

			tags, err := readTags(properties.Mappable())
			if err != nil {
				return err
			}

			body := map[string]any{
				tagsProp: tags.SortedValues(),
			}

			_, err = post(ctx, path, body, apiVersion, azureClient)
			return err
		},
	}
}

func post(ctx context.Context, path string, body map[string]any, apiVersion string, azureClient azure.AzureClient) (map[string]any, error) {
	queryParams := map[string]any{
		"api-version": apiVersion,
	}
	resp, err := azureClient.Post(ctx, path, body, queryParams)
	if err != nil {
		return nil, err
	}
	return resp.(map[string]any), nil
}

func readTags(p map[string]any) (codegen.StringSet, error) {
	tagsVal, ok := p[tagsProp]
	if !ok {
		return nil, errors.New("missing required property 'tags'")
	}
	tags, ok := tagsVal.([]any)
	if !ok {
		return nil, errors.New("property 'tags' must be an array")
	}
	strTags := codegen.NewStringSet()
	for _, tag := range tags {
		strTags.Add(tag.(string))
	}
	return strTags, nil
}
