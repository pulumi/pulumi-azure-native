package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewFirewallPolicy(ctx, "firewallPolicy", &network.FirewallPolicyArgs{
			DnsSettings: &network.DnsSettingsArgs{
				EnableProxy:                 pulumi.Bool(true),
				RequireProxyForNetworkRules: pulumi.Bool(false),
				Servers: pulumi.StringArray{
					pulumi.String("30.3.4.5"),
				},
			},
			ExplicitProxy: &network.ExplicitProxyArgs{
				EnableExplicitProxy: pulumi.Bool(true),
				EnablePacFile:       pulumi.Bool(true),
				HttpPort:            pulumi.Int(8087),
				HttpsPort:           pulumi.Int(8087),
				PacFile:             pulumi.String("https://tinawstorage.file.core.windows.net/?sv=2020-02-10&ss=bfqt&srt=sco&sp=rwdlacuptfx&se=2021-06-04T07:01:12Z&st=2021-06-03T23:01:12Z&sip=68.65.171.11&spr=https&sig=Plsa0RRVpGbY0IETZZOT6znOHcSro71LLTTbzquYPgs%3D"),
				PacFilePort:         pulumi.Int(8087),
			},
			FirewallPolicyName: pulumi.String("firewallPolicy"),
			Insights: &network.FirewallPolicyInsightsArgs{
				IsEnabled: pulumi.Bool(true),
				LogAnalyticsResources: &network.FirewallPolicyLogAnalyticsResourcesArgs{
					DefaultWorkspaceId: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/defaultWorkspace"),
					},
					Workspaces: network.FirewallPolicyLogAnalyticsWorkspaceArray{
						&network.FirewallPolicyLogAnalyticsWorkspaceArgs{
							Region: pulumi.String("westus"),
							WorkspaceId: &network.SubResourceArgs{
								Id: pulumi.String("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace1"),
							},
						},
						&network.FirewallPolicyLogAnalyticsWorkspaceArgs{
							Region: pulumi.String("eastus"),
							WorkspaceId: &network.SubResourceArgs{
								Id: pulumi.String("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace2"),
							},
						},
					},
				},
				RetentionDays: pulumi.Int(100),
			},
			IntrusionDetection: &network.FirewallPolicyIntrusionDetectionArgs{
				Configuration: &network.FirewallPolicyIntrusionDetectionConfigurationArgs{
					BypassTrafficSettings: network.FirewallPolicyIntrusionDetectionBypassTrafficSpecificationsArray{
						&network.FirewallPolicyIntrusionDetectionBypassTrafficSpecificationsArgs{
							Description: pulumi.String("Rule 1"),
							DestinationAddresses: pulumi.StringArray{
								pulumi.String("5.6.7.8"),
							},
							DestinationPorts: pulumi.StringArray{
								pulumi.String("*"),
							},
							Name:     pulumi.String("bypassRule1"),
							Protocol: pulumi.String("TCP"),
							SourceAddresses: pulumi.StringArray{
								pulumi.String("1.2.3.4"),
							},
						},
					},
					SignatureOverrides: network.FirewallPolicyIntrusionDetectionSignatureSpecificationArray{
						&network.FirewallPolicyIntrusionDetectionSignatureSpecificationArgs{
							Id:   pulumi.String("2525004"),
							Mode: pulumi.String("Deny"),
						},
					},
				},
				Mode: pulumi.String("Alert"),
			},
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("rg1"),
			Sku: &network.FirewallPolicySkuArgs{
				Tier: pulumi.String("Premium"),
			},
			Snat: &network.FirewallPolicySNATArgs{
				PrivateRanges: pulumi.StringArray{
					pulumi.String("IANAPrivateRanges"),
				},
			},
			Sql: &network.FirewallPolicySQLArgs{
				AllowSqlRedirect: pulumi.Bool(true),
			},
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			ThreatIntelMode: pulumi.String("Alert"),
			ThreatIntelWhitelist: &network.FirewallPolicyThreatIntelWhitelistArgs{
				Fqdns: pulumi.StringArray{
					pulumi.String("*.microsoft.com"),
				},
				IpAddresses: pulumi.StringArray{
					pulumi.String("20.3.4.5"),
				},
			},
			TransportSecurity: &network.FirewallPolicyTransportSecurityArgs{
				CertificateAuthority: &network.FirewallPolicyCertificateAuthorityArgs{
					KeyVaultSecretId: pulumi.String("https://kv/secret"),
					Name:             pulumi.String("clientcert"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
