// Copyright 2021, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
using Pulumi.AzureNative.Authorization;
using Pulumi.AzureNative.Network;
using Pulumi.AzureNative.Network.Inputs;
using SubnetArgs = Pulumi.AzureNative.Network.Inputs.SubnetArgs;
using Pulumi.AzureNative.Resources;
using Pulumi.AzureNative.Sql;
using Pulumi.AzureNative.Sql.Inputs;
using Pulumi.AzureNative.Storage;

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

        var enableADS = new ServerSecurityAlertPolicy("default", new ServerSecurityAlertPolicyArgs
        {
            ResourceGroupName = resourceGroup.Name,
            ServerName = server.Name,
            State = SecurityAlertsPolicyState.Enabled
        });

        var sa = new StorageAccount("sa", new StorageAccountArgs
        {
            ResourceGroupName = resourceGroup.Name,
            Sku = new Pulumi.AzureNative.Storage.Inputs.SkuArgs
            {
                Name = Pulumi.AzureNative.Storage.SkuName.Standard_LRS
            },
            Kind = Kind.StorageV2,
        });

        var container = new BlobContainer("blobs", new BlobContainerArgs
        {
            ResourceGroupName = resourceGroup.Name,
            AccountName = sa.Name
        });

        new ServerVulnerabilityAssessment("default", new ServerVulnerabilityAssessmentArgs
        {
            ResourceGroupName = resourceGroup.Name,
            ServerName = server.Name,
            RecurringScans = new VulnerabilityAssessmentRecurringScansPropertiesArgs
            {
                EmailSubscriptionAdmins = false,
                Emails =
                {
                    "hi@example.com",
                },
                IsEnabled = true
            },
            StorageContainerPath = Output.Format($"https://{sa.Name}.blob.core.windows.net/{container.Name}"),
            StorageAccountAccessKey = Output.Tuple(resourceGroup.Name, sa.Name).Apply(names =>
                GetStorageAccountPrimaryKey(names.Item1, names.Item2))
        }, new CustomResourceOptions { DependsOn = { enableADS } });
        
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

        var config = Output.Create(GetClientConfig.InvokeAsync());

        var serverAzureAdAdministrator = new ServerAzureADAdministrator("ad-admin",new ServerAzureADAdministratorArgs
        {
            ResourceGroupName = resourceGroup.Name,
            ServerName = server.Name,
            AdministratorName = "ActiveDirectory",
            AdministratorType = AdministratorType.ActiveDirectory,
            Login = "foo@example.com",
            Sid = "c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c",
            TenantId = config.Apply(c => c.TenantId)
        });

        var serverAzureAdOnlyAuth = new ServerAzureADOnlyAuthentication("ad-only-auth", new ServerAzureADOnlyAuthenticationArgs
        {
            ResourceGroupName = resourceGroup.Name,
            ServerName = server.Name,
            AuthenticationName = "Default",
            AzureADOnlyAuthentication = false
        }, new CustomResourceOptions {DependsOn = { serverAzureAdAdministrator }});
    }
    
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
