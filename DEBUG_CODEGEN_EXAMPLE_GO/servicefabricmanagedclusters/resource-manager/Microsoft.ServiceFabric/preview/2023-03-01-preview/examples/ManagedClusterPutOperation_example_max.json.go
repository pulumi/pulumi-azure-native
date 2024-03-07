package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabric.NewManagedCluster(ctx, "managedCluster", &servicefabric.ManagedClusterArgs{
			AddonFeatures: pulumi.StringArray{
				pulumi.String("DnsService"),
				pulumi.String("BackupRestoreService"),
				pulumi.String("ResourceMonitorService"),
			},
			AdminPassword:  pulumi.String("{vm-password}"),
			AdminUserName:  pulumi.String("vmadmin"),
			AllowRdpAccess: pulumi.Bool(true),
			ApplicationTypeVersionsCleanupPolicy: &servicefabric.ApplicationTypeVersionsCleanupPolicyArgs{
				MaxUnusedVersionsToKeep: pulumi.Int(3),
			},
			AuxiliarySubnets: []servicefabric.SubnetArgs{
				{
					EnableIpv6:                        pulumi.Bool(true),
					Name:                              pulumi.String("testSubnet1"),
					NetworkSecurityGroupId:            pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/networkSecurityGroups/sn1"),
					PrivateEndpointNetworkPolicies:    pulumi.String("enabled"),
					PrivateLinkServiceNetworkPolicies: pulumi.String("enabled"),
				},
			},
			ClientConnectionPort: pulumi.Int(19000),
			ClusterCodeVersion:   pulumi.String("7.1.168.9494"),
			ClusterName:          pulumi.String("myCluster"),
			ClusterUpgradeMode:   pulumi.String("Manual"),
			DnsName:              pulumi.String("myCluster"),
			EnableAutoOSUpgrade:  pulumi.Bool(true),
			EnableIpv6:           pulumi.Bool(true),
			FabricSettings: []servicefabric.SettingsSectionDescriptionArgs{
				{
					Name: pulumi.String("ManagedIdentityTokenService"),
					Parameters: servicefabric.SettingsParameterDescriptionArray{
						{
							Name:  pulumi.String("IsEnabled"),
							Value: pulumi.String("true"),
						},
					},
				},
			},
			HttpGatewayConnectionPort: pulumi.Int(19080),
			IpTags: []servicefabric.IPTagArgs{
				{
					IpTagType: pulumi.String("FirstPartyUsage"),
					Tag:       pulumi.String("SQL"),
				},
			},
			LoadBalancingRules: []servicefabric.LoadBalancingRuleArgs{
				{
					BackendPort:   pulumi.Int(80),
					FrontendPort:  pulumi.Int(80),
					ProbePort:     pulumi.Int(80),
					ProbeProtocol: pulumi.String("http"),
					Protocol:      pulumi.String("http"),
				},
				{
					BackendPort:   pulumi.Int(443),
					FrontendPort:  pulumi.Int(443),
					ProbePort:     pulumi.Int(443),
					ProbeProtocol: pulumi.String("http"),
					Protocol:      pulumi.String("http"),
				},
				{
					BackendPort:      pulumi.Int(10000),
					FrontendPort:     pulumi.Int(10000),
					LoadDistribution: pulumi.String("Default"),
					ProbePort:        pulumi.Int(10000),
					ProbeProtocol:    pulumi.String("http"),
					Protocol:         pulumi.String("tcp"),
				},
			},
			Location: pulumi.String("eastus"),
			NetworkSecurityRules: []servicefabric.NetworkSecurityRuleArgs{
				{
					Access:      pulumi.String("allow"),
					Description: pulumi.String("Test description"),
					DestinationAddressPrefixes: pulumi.StringArray{
						pulumi.String("*"),
					},
					DestinationPortRanges: pulumi.StringArray{
						pulumi.String("*"),
					},
					Direction: pulumi.String("inbound"),
					Name:      pulumi.String("TestName"),
					Priority:  pulumi.Int(1010),
					Protocol:  pulumi.String("tcp"),
					SourceAddressPrefixes: pulumi.StringArray{
						pulumi.String("*"),
					},
					SourcePortRanges: pulumi.StringArray{
						pulumi.String("*"),
					},
				},
				{
					Access:                   pulumi.String("allow"),
					DestinationAddressPrefix: pulumi.String("*"),
					DestinationPortRange:     pulumi.String("33500-33699"),
					Direction:                pulumi.String("inbound"),
					Name:                     pulumi.String("AllowARM"),
					Priority:                 pulumi.Int(2002),
					Protocol:                 pulumi.String("*"),
					SourceAddressPrefix:      pulumi.String("AzureResourceManager"),
					SourcePortRange:          pulumi.String("*"),
				},
			},
			PublicIPPrefixId:  pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.Network/publicIPPrefixes/myPublicIPPrefix"),
			ResourceGroupName: pulumi.String("resRg"),
			ServiceEndpoints: []servicefabric.ServiceEndpointArgs{
				{
					Locations: pulumi.StringArray{
						pulumi.String("eastus2"),
						pulumi.String("usnorth"),
					},
					Service: pulumi.String("Microsoft.Storage"),
				},
			},
			Sku: &servicefabric.SkuArgs{
				Name: pulumi.String("Basic"),
			},
			Tags:            nil,
			UseCustomVnet:   pulumi.Bool(true),
			ZonalResiliency: pulumi.Bool(true),
			ZonalUpdateMode: pulumi.String("Fast"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
