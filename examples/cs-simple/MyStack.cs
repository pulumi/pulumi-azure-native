using Pulumi;
using Pulumi.AzureRM.Resources.V20200601;
using Pulumi.AzureRM.Storage.V20190601;
using Pulumi.AzureRM.Storage.V20190601.Inputs;

class MyStack : Stack
{
    public MyStack()
    {
        var resourceGroup = new ResourceGroup("resourceGroup", new ResourceGroupArgs
        {
            ResourceGroupName = "azurerm-dotnet",
            Location = "WestUS"
        });

        var storageAccount = new StorageAccount("sa", new StorageAccountArgs
        {
            ResourceGroupName = resourceGroup.Name,
            AccountName = "pulumiassacs",
            Location = "WestUS",
            Sku = new SkuArgs
            {
                Name = "Standard_LRS",
                Tier = "Standard"
            },
            Kind = "StorageV2"
        });
    }
}
