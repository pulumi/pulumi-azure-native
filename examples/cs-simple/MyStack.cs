using System.Threading.Tasks;
using Pulumi;
using Pulumi.AzureNextGen.Resources.Latest;
using Pulumi.AzureNextGen.Storage.Latest;
using Pulumi.AzureNextGen.Storage.Latest.Inputs;
using Pulumi.Random;

class MyStack : Stack
{
    public MyStack()
    {
        var randomString = new RandomString("random", new RandomStringArgs
        {
            Length = 12,
            Special = false,
            Upper = false,
        });

        var resourceGroup = new ResourceGroup("resourceGroup", new ResourceGroupArgs
        {
            ResourceGroupName = randomString.Result,
            Location = "WestUS"
        });

        var storageAccount = new StorageAccount("sa", new StorageAccountArgs
        {
            ResourceGroupName = resourceGroup.Name,
            AccountName = randomString.Result,
            Location = resourceGroup.Location,
            Sku = new SkuArgs
            {
                Name = "Standard_LRS",
                Tier = "Standard"
            },
            Kind = "StorageV2"
        });

        this.PrimaryStorageKey = Output.Tuple(resourceGroup.Name, storageAccount.Name).Apply(names =>
            Output.CreateSecret(GetStorageAccountPrimaryKey(names.Item1, names.Item2)));
    }

    [Output]
    public Output<string> PrimaryStorageKey { get; set; }

    private static async Task<string> GetStorageAccountPrimaryKey(string resourceGroupName, string accountName)
    {
        var accountKeys = await ListStorageAccountKeys.InvokeAsync(new ListStorageAccountKeysArgs
        {
            ResourceGroupName = resourceGroupName,
            AccountName = accountName
        });
        return accountKeys.Keys[0].Value;
    }
}
