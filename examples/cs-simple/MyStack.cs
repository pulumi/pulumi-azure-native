// Copyright 2021, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.Azure.Management.Storage;
using EncryptionScope = Microsoft.Azure.Management.Storage.Models.EncryptionScope;
using Microsoft.Rest;
using Pulumi;
using Pulumi.AzureNative.Authorization;
using Pulumi.AzureNative.Resources;
using Deployment = Pulumi.AzureNative.Resources.Deployment;
using Pulumi.AzureNative.Resources.Inputs;
using Pulumi.AzureNative.Storage;
using SkuArgs = Pulumi.AzureNative.Storage.Inputs.SkuArgs;

class MyStack : Stack
{
    public MyStack()
    {
        var resourceGroup = new ResourceGroup("resourceGroup");

        var tags = new TagAtScope("tagAtScope", new TagAtScopeArgs
        {
            Scope = resourceGroup.Id,
            Properties = new TagsArgs
            {
                Tags =
                {
                    ["foo"] = "bar",
                }
            }
        });

        var storageAccount = new StorageAccount("sa", new StorageAccountArgs
        {
            ResourceGroupName = resourceGroup.Name,
            Location = resourceGroup.Location,
            Sku = new SkuArgs
            {
                Name = SkuName.Standard_LRS
            },
            Kind = Kind.StorageV2
        });

        new Deployment("template", new DeploymentArgs
        {
            ResourceGroupName = resourceGroup.Name,
            Properties = new DeploymentPropertiesArgs
            {
                Mode = DeploymentMode.Incremental,
                Template = new Dictionary<string, object>
                {
                    { "$schema", "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#" },
                    { "contentVersion", "1.0.0.0" },
                    { "resources", new object[0] }
                }
            }
        });

        this.PrimaryStorageKey = Output.Tuple(resourceGroup.Name, storageAccount.Name).Apply(names =>
            Output.CreateSecret(GetStorageAccountPrimaryKey(names.Item1, names.Item2)));
        this.EncryptionScopeState = Output.Tuple(resourceGroup.Name, storageAccount.Name).Apply(names =>
            CreateEncryptionScope(names.Item1, names.Item2, "test"));
    }

    [Output]
    public Output<string> PrimaryStorageKey { get; set; }
    
    [Output]
    public Output<string> EncryptionScopeState { get; set; }


    private static async Task<string> GetStorageAccountPrimaryKey(string resourceGroupName, string accountName)
    {
        var accountKeys = await ListStorageAccountKeys.InvokeAsync(new ListStorageAccountKeysArgs
        {
            ResourceGroupName = resourceGroupName,
            AccountName = accountName
        });
        return accountKeys.Keys[0].Value;
    }
    
    private static async Task<string> CreateEncryptionScope(string resourceGroupName, string accountName, string scopeName)
    {
        // Example of using the Azure SDK in Pulumi programs
        var config = await GetClientConfig.InvokeAsync();
        var token = await GetClientToken.InvokeAsync();
        var client = new StorageManagementClient(new TokenCredentials(new StringTokenProvider(token.Token, "Bearer")))
        {
            SubscriptionId = config.SubscriptionId
        };

        var scope = await client.EncryptionScopes.PutAsync(
            resourceGroupName, accountName, scopeName, new EncryptionScope());
        return scope.State;
    }
}
