using Pulumi;
using Pulumi.AzureNextGen.Network.Latest;
using Pulumi.AzureNextGen.Network.Latest.Inputs;
using SubnetArgs = Pulumi.AzureNextGen.Network.Latest.Inputs.SubnetArgs;
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
            ResourceGroupName = randomString.Result
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
        
        var sqlFwRuleClientIP = new FirewallRule("sqlFwRuleClientIP", new FirewallRuleArgs
        {
            ResourceGroupName = resourceGroup.Name,
            ServerName = server.Name,
            FirewallRuleName = "ClientIPAddress",
            StartIpAddress = "222.222.222.222",
            EndIpAddress = "222.222.222.222"
        });
        
        var vnet = new VirtualNetwork("vnet", new VirtualNetworkArgs
        {
            ResourceGroupName = resourceGroup.Name,
            VirtualNetworkName = randomString.Result,
            Location = resourceGroup.Location,
            AddressSpace = new AddressSpaceArgs { AddressPrefixes = {"10.1.0.0/16"} },
            Subnets = 
            {
                new SubnetArgs
                {
                    Name = "default",
                    AddressPrefix = "10.1.0.0/24",
                    PrivateEndpointNetworkPolicies = "Disabled"
                }
            }
        });

        var privateEndpoint = new PrivateEndpoint("endpoint", new PrivateEndpointArgs
        {
            ResourceGroupName = resourceGroup.Name,
            PrivateEndpointName = randomString.Result,
            Location = resourceGroup.Location,
            PrivateLinkServiceConnections =
            {
                new PrivateLinkServiceConnectionArgs
                {
                    GroupIds = {"sqlServer"},
                    PrivateLinkServiceId = server.Id,
                    Name = "conn-sql",
                }
            },
            Subnet = new SubnetArgs { Id = vnet.Subnets.GetAt(0).Apply(v => v.Id) }
        });

        var privateDnsZone = new PrivateDnsZoneGroup("zoneGroup", new PrivateDnsZoneGroupArgs
        {
            ResourceGroupName = resourceGroup.Name,
            PrivateDnsZoneGroupName = randomString.Result,
            PrivateEndpointName = privateEndpoint.Name,
        });
    }
}
