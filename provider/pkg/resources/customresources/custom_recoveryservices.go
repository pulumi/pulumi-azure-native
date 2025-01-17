// Copyright 2024, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	recovery "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

const (
	protectedItemPath                  = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}"
	recoveryServicesAzureStorageFilter = "backupManagementType eq 'AzureStorage'"
)

type protectedItem struct {
	protectedItemType string
	policyID          string
	// For file shares, this is the storage account ID.
	sourceResourceId string
}

type protectedItemProperties struct {
	protectedItem
	// The internal name of the protected item, also called system name. It looks like
	// "azurefileshare;339f98592ed6329dc5f83bdbb8675bd3528bf58d2246d5e172615186febdc45c". It's determined by
	// readSystemNameFromProtectableItem.
	name              string
	resourceGroupName string
	recoveryVaultName string
	containerName     string
	itemName          string
	fabricName        string
}

// A custom resource for Microsoft.RecoveryServices ProtectedItem, specifically the file share protected item. It looks
// up the magic "system name" of the file share and adds it to the inputs. For other types of protected items, it does
// nothing. #1420
func recoveryServicesProtectedItem(subscription string, cred azcore.TokenCredential) (*CustomResource, error) {
	clientFactory, err := recovery.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create recovery services client factory: %w", err)
	}

	reader := &systemNameReaderImpl{
		protectableItemsClient:     clientFactory.NewBackupProtectableItemsClient(),
		protectionContainersClient: clientFactory.NewProtectionContainersClient(),
	}

	return &CustomResource{
		path: protectedItemPath,
		tok:  "azure-native:recoveryservices:ProtectedItem",
		GetIdAndQuery: func(ctx context.Context, inputs resource.PropertyMap, crudClient crud.ResourceCrudClient) (string, map[string]any, error) {
			return getIdAndQuery(ctx, inputs, crudClient, reader)
		},
	}, nil
}

// getIdAndQuery replaces the default implementation of crud.ResourceCrudClient.PrepareAzureRESTIdAndQuery. It doesn't
// change queryParams, only the id, to replace the file share's friendly name with the system name.
func getIdAndQuery(ctx context.Context, inputs resource.PropertyMap, crudClient crud.ResourceCrudClient, reader systemNameReader) (string, map[string]any, error) {
	systemName, err := retrieveSystemName(ctx, inputs, reader, 10, 30*time.Second)
	if err != nil {
		return "", nil, fmt.Errorf("failed to retrieve system name: %w", err)
	}

	inputsCopy := resource.NewPropertyMapFromMap(inputs.Mappable()) // deep copy
	if systemName != "" {
		inputsCopy["protectedItemName"] = resource.NewStringProperty(systemName)
	}

	origId, queryParams, err := crudClient.PrepareAzureRESTIdAndQuery(inputsCopy)
	if err != nil {
		return "", nil, fmt.Errorf("failed to prepare Azure REST ID and query: %w", err)
	}

	return origId, queryParams, nil
}

// retrieveSystemName looks up the system name of a file share protected item in Azure.
func retrieveSystemName(ctx context.Context, input resource.PropertyMap, reader systemNameReader, maxAttempts int, waitBetweenAttempts time.Duration) (string, error) {
	item, err := extractProtectedItemProperties(input)
	if err != nil {
		return "", fmt.Errorf("failed to extract protected item properties from input: %w", err)
	}

	if item.protectedItemType != "AzureFileShareProtectedItem" {
		logging.V(9).Infof("not modifying protected item of type %s", item.protectedItemType)
		return "", nil
	}

	err = reader.inquireContainer(ctx, item)
	if err != nil {
		return "", fmt.Errorf("failed to inquire protection container %s: %w", item.containerName, err)
	}

	logging.V(9).Infof("looking up system name for %s", item.itemName)
	// Based on observations during testing, the system name is usually, but not always immediately available.
	// The /inquire request above should be awaited according to the docs, but the SDK doesn't actually support that.
	for i := 1; i <= maxAttempts; i++ {
		systemName, err := reader.readSystemNameFromProtectableItem(ctx, item)
		if err != nil {
			return "", fmt.Errorf("failed to read system name for protectable item: %w", err)
		}
		if systemName != "" {
			logging.V(9).Infof("found system name %s for %s in attempt %d", systemName, item.itemName, i)
			return systemName, nil
		}
		if i < maxAttempts {
			time.Sleep(waitBetweenAttempts)
		}
	}
	return "", fmt.Errorf("failed to retrieve system name after %d attempts", maxAttempts)
}

// systemNameReader is an interface for getting the Azure system name of a protected item.
// The system name looks like "azurefileshare;339f98592ed6329dc5f83bdbb8675bd3528bf58d2246d5e172615186febdc45c" and
// needs to be determined by iterating over the protected items in scope and matching by the human-readable name.
// Abstracted into an interface to allow for testing.
type systemNameReader interface {
	// inquireContainer makes a request to /inquire, a special recovery services API that triggers a background job to
	// update the protectable items in scope.
	inquireContainer(ctx context.Context, input *protectedItemProperties) error
	readSystemNameFromProtectableItem(ctx context.Context, input *protectedItemProperties) (string, error)
}

type systemNameReaderImpl struct {
	protectableItemsClient     *recovery.BackupProtectableItemsClient
	protectionContainersClient *recovery.ProtectionContainersClient
}

func (s *systemNameReaderImpl) readSystemNameFromProtectableItem(ctx context.Context, input *protectedItemProperties) (string, error) {
	segments := strings.Split(input.sourceResourceId, "/")
	storageAccountName := segments[len(segments)-1]

	protectablePager := s.protectableItemsClient.NewListPager(input.recoveryVaultName, input.resourceGroupName, &recovery.BackupProtectableItemsClientListOptions{
		Filter: to.Ptr(recoveryServicesAzureStorageFilter),
	})
	for protectablePager.More() {
		page, err := protectablePager.NextPage(ctx)
		if err != nil {
			return "", fmt.Errorf("failed to get next page of protectable items: %w", err)
		}
		systemName, err := findSystemName(page.Value, input.itemName, storageAccountName)
		if err != nil {
			return "", fmt.Errorf("failed to find system name for protectable item: %w", err)
		}
		if systemName != "" {
			return systemName, nil
		}
	}
	return "", nil
}

func (s *systemNameReaderImpl) inquireContainer(ctx context.Context, item *protectedItemProperties) error {
	// the first return value is an empty struct, so we can ignore it
	_, err := s.protectionContainersClient.Inquire(ctx, item.recoveryVaultName, item.resourceGroupName, item.fabricName, item.containerName,
		&recovery.ProtectionContainersClientInquireOptions{
			Filter: to.Ptr(recoveryServicesAzureStorageFilter),
		})
	return err
}

// findSystemName finds the system name of the file share protected item by looking through the given protectable items
// for one that matches the target friendly name and storage account name.
func findSystemName(protectableItems []*recovery.WorkloadProtectableItemResource, targetFriendlyName, storageAccountName string) (string, error) {
	for _, item := range protectableItems {
		protectableFileShare, ok := item.Properties.(*recovery.AzureFileShareProtectableItem)
		if !ok {
			itemType := "unknown"
			if item.Type != nil {
				itemType = *item.Type
			}
			return "", fmt.Errorf("protectable item of type %s is not a file share", itemType)
		}
		friendlyName := *protectableFileShare.FriendlyName
		containerName := *protectableFileShare.ParentContainerFriendlyName
		if friendlyName == targetFriendlyName && containerName == storageAccountName {
			return *item.Name, nil
		}
	}
	return "", nil
}

func extractProtectedItemProperties(properties resource.PropertyMap) (*protectedItemProperties, error) {
	rg, err := getRequiredStringProperty(properties, resourceGroupName)
	if err != nil {
		return nil, err
	}

	// not set for new resources
	var name string
	if systemName, ok := properties["name"]; ok {
		name = systemName.StringValue()
	}

	recoveryVaultName, err := getRequiredStringProperty(properties, "vaultName")
	if err != nil {
		return nil, err
	}

	containerName, err := getRequiredStringProperty(properties, "containerName")
	if err != nil {
		return nil, err
	}

	itemName, err := getRequiredStringProperty(properties, "protectedItemName")
	if err != nil {
		return nil, err
	}

	fabricName, err := getRequiredStringProperty(properties, "fabricName")
	if err != nil {
		return nil, err
	}

	if !properties.HasValue("properties") {
		return nil, fmt.Errorf("'properties' not found in input")
	}
	itemProperties := properties["properties"].ObjectValue()
	itemType, err := getRequiredStringProperty(itemProperties, "protectedItemType")
	if err != nil {
		return nil, err
	}
	policyID, err := getRequiredStringProperty(itemProperties, "policyId")
	if err != nil {
		return nil, err
	}
	sourceResourceId, err := getRequiredStringProperty(itemProperties, "sourceResourceId")
	if err != nil {
		return nil, err
	}

	return &protectedItemProperties{
		name: name,
		protectedItem: protectedItem{
			protectedItemType: itemType,
			policyID:          policyID,
			sourceResourceId:  sourceResourceId,
		},
		resourceGroupName: rg,
		recoveryVaultName: recoveryVaultName,
		containerName:     containerName,
		itemName:          itemName,
		fabricName:        fabricName,
	}, nil
}

func getRequiredStringProperty(properties resource.PropertyMap, key resource.PropertyKey) (string, error) {
	prop, ok := properties[key]
	if !ok {
		return "", fmt.Errorf("%s not found", key)
	}
	if !prop.IsString() {
		return "", fmt.Errorf("%s is not a string", key)
	}
	return prop.StringValue(), nil
}
