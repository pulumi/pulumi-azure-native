using Pulumi;
using Pulumi.AzureNextGen.Resources.Latest;
using Pulumi.AzureNextGen.Sql.Latest;
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

        var server = new Server("server", new ServerArgs
        {
            ServerName = randomString.Result,
            ResourceGroupName = resourceGroup.Name,
            Location = resourceGroup.Location,
            AdministratorLogin = "dummylogin",
            AdministratorLoginPassword = "Un53cuRE!",
            Version = "12.0"
        });
        
        var sqlFwRuleClientIP2 = new FirewallRule("sqlFwRuleClientIP", new FirewallRuleArgs
        {
            ResourceGroupName = resourceGroup.Name,
            ServerName = server.Name,
            FirewallRuleName = "ClientIPAddress",
            StartIpAddress = "222.222.222.222",
            EndIpAddress = "222.222.222.222"
        });
    }
}
