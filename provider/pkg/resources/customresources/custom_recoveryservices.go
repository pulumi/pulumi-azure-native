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

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
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
		return nil, err
	}

	reader := &systemNameReaderImpl{
		protectableItemsClient: clientFactory.NewBackupProtectableItemsClient(),
	}

	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}",
		tok:  "azure-native:recoveryservices:ProtectedItem",
		PreProcessInputs: func(ctx context.Context, input resource.PropertyMap) (resource.PropertyMap, error) {
			err := updateInputWithFileShareSystemName(ctx, input, reader)
			if err != nil {
				return nil, err
			}
			return input, nil
		},
	}, nil
}

// updateInputWithFileShareSystemName updates the "protectedItemName" from the input with the system name of the file
// share protected item, looked up via `reader`
func updateInputWithFileShareSystemName(ctx context.Context, input resource.PropertyMap, reader systemNameReader) error {
	item, err := extractProtectedItemProperties(input)
	if err != nil {
		return err
	}

	if item.protectedItemType != "AzureFileShareProtectedItem" {
		logging.V(9).Infof("not modifying protected item of type %s", item.protectedItemType)
		return nil
	}

	logging.V(9).Infof("looking up system name for %s", item.itemName)
	systemName, err := reader.readSystemNameFromProtectableItem(ctx, item)
	if err != nil {
		return err
	}
	// Based on observations during testing, the system name is usually, but not always immediately available.
	if systemName == "" {
		time.Sleep(30 * time.Second)
		systemName, err = reader.readSystemNameFromProtectableItem(ctx, item)
		if err != nil {
			return err
		}
	}

	if systemName != "" {
		input["__friendlyProtectedItemName"] = input["protectedItemName"]
		input["protectedItemName"] = resource.NewStringProperty(systemName)
	} else {
		logging.V(5).Infof("no system name found for %s", input["protectedItemName"])
	}
	return nil
}

// systemNameReader is an interface for getting the Azure system name of a protected item.
// The system name looks like "azurefileshare;339f98592ed6329dc5f83bdbb8675bd3528bf58d2246d5e172615186febdc45c" and
// needs to be determined by iterating over the protected items in scope and matching by the human-readable name.
// Abstracted into an interface to allow for testing.
type systemNameReader interface {
	readSystemNameFromProtectableItem(ctx context.Context, input *protectedItemProperties) (string, error)
}

type systemNameReaderImpl struct {
	protectableItemsClient *recovery.BackupProtectableItemsClient
}

func (s *systemNameReaderImpl) readSystemNameFromProtectableItem(ctx context.Context, input *protectedItemProperties) (string, error) {
	segments := strings.Split(input.sourceResourceId, "/")
	storageAccountName := segments[len(segments)-1]

	protectablePager := s.protectableItemsClient.NewListPager(input.recoveryVaultName, input.resourceGroupName, &recovery.BackupProtectableItemsClientListOptions{
		Filter: to.Ptr("backupManagementType eq 'AzureStorage'"),
	})
	for protectablePager.More() {
		page, err := protectablePager.NextPage(ctx)
		if err != nil {
			return "", err
		}
		systemName, err := findSystemName(page.Value, input.itemName, storageAccountName)
		if err != nil {
			return "", err
		}
		if systemName != "" {
			return systemName, nil
		}
	}
	return "", nil
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
		return nil, fmt.Errorf("properties not found")
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
