// Copyright 2024, Pulumi Corporation.  All rights reserved.

using Pulumi;
using Pulumi.AzureNative.Resources;
using Pulumi.AzureNative.DBforPostgreSQL;
using Pulumi.AzureNative.DBforPostgreSQL.Inputs;

class MyStack : Stack
{
    public MyStack()
    {
        var resourceGroup = new ResourceGroup("resourceGroup");

        var server = new Server("server", new ServerArgs
        {
            ResourceGroupName = resourceGroup.Name,
            Location = resourceGroup.Location,
            Version = "14",
            Sku = new SkuArgs
            {
                Name = "Standard_B1ms",
                Tier = SkuTier.Burstable
            },
            Storage = new StorageArgs
            {
                StorageSizeGB = 32
            },
            AdministratorLogin = "pulumiadminuser",
            AdministratorLoginPassword = "pulumiadminpassword"
        });
    }
}