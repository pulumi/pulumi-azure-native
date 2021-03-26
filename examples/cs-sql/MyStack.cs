// Copyright 2021, Pulumi Corporation.  All rights reserved.

using Pulumi;
using Pulumi.AzureNative.Network;
using Pulumi.AzureNative.Network.Inputs;
using SubnetArgs = Pulumi.AzureNative.Network.Inputs.SubnetArgs;
using Pulumi.AzureNative.Resources;
using Pulumi.AzureNative.Sql;

class MyStack : Stack
{
    public MyStack()
    {
        var resourceGroup = new ResourceGroup("resourceGroup");

        var server = new Server("server", new ServerArgs
        {
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
            PrivateEndpointName = privateEndpoint.Name,
        });
    }
}
