using Pulumi;
using Pulumi.AzureRM;

class MyStack : Stack
{
    public MyStack()
    {
        // Create an Azure Resource Group
        var resourceGroup = new ResourceGroup("resourceGroup", new ResourceGroupArgs
        {
            Name = "azurerm-dotnet",
            Location = "WestUS"
        });
    }
}
