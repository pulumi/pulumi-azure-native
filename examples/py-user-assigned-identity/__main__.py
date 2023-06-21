# Copyright 2023, Pulumi Corporation.  All rights reserved.

import pulumi
from pulumi_azure_native import storage
from pulumi_azure_native import managedidentity
from pulumi_azure_native import resources

# Create an Azure Resource Group
resource_group = resources.ResourceGroup('resource_group')

# Create a user-assigned managed identity
user_assigned_identity = managedidentity.UserAssignedIdentity("my-user-assigned-identity",
    resource_group_name=resource_group.name)

# Create a storage account that references the user-assigned managed identity
storage_account = storage.StorageAccount("mystorageaccount",
    resource_group_name=resource_group.name,
    kind="StorageV2",
    location=resource_group.location,
    identity=storage.IdentityArgs(
        type="UserAssigned",
        user_assigned_identities=[user_assigned_identity.id],
        ),
    sku=storage.SkuArgs(
        name="Standard_LRS",
    ))