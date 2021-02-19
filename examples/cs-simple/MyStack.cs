using System.Threading.Tasks;
using Microsoft.Azure.Management.Storage;
using EncryptionScope = Microsoft.Azure.Management.Storage.Models.EncryptionScope;
using Microsoft.Rest;
using Pulumi;
using Pulumi.AzureNextGen.Authorization;
using Pulumi.AzureNextGen.Resources;
using Pulumi.AzureNextGen.Storage;
using Pulumi.AzureNextGen.Storage.Inputs;

class MyStack : Stack
{
    public MyStack()
    {
        var resourceGroup = new ResourceGroup("resourceGroup");

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
