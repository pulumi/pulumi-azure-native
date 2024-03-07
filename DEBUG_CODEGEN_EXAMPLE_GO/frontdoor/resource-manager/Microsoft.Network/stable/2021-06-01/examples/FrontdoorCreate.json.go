package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewFrontDoor(ctx, "frontDoor", &network.FrontDoorArgs{
			BackendPools: network.BackendPoolArray{
				&network.BackendPoolArgs{
					Backends: network.BackendArray{
						&network.BackendArgs{
							Address:   pulumi.String("w3.contoso.com"),
							HttpPort:  pulumi.Int(80),
							HttpsPort: pulumi.Int(443),
							Priority:  pulumi.Int(2),
							Weight:    pulumi.Int(1),
						},
						&network.BackendArgs{
							Address:                    pulumi.String("contoso.com.website-us-west-2.othercloud.net"),
							HttpPort:                   pulumi.Int(80),
							HttpsPort:                  pulumi.Int(443),
							Priority:                   pulumi.Int(1),
							PrivateLinkApprovalMessage: pulumi.String("Please approve the connection request for this Private Link"),
							PrivateLinkLocation:        pulumi.String("eastus"),
							PrivateLinkResourceId:      pulumi.String("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Network/privateLinkServices/pls1"),
							Weight:                     pulumi.Int(2),
						},
						&network.BackendArgs{
							Address:                    pulumi.String("10.0.1.5"),
							HttpPort:                   pulumi.Int(80),
							HttpsPort:                  pulumi.Int(443),
							Priority:                   pulumi.Int(1),
							PrivateLinkAlias:           pulumi.String("APPSERVER.d84e61f0-0870-4d24-9746-7438fa0019d1.westus2.azure.privatelinkservice"),
							PrivateLinkApprovalMessage: pulumi.String("Please approve this request to connect to the Private Link"),
							Weight:                     pulumi.Int(1),
						},
					},
					HealthProbeSettings: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/healthProbeSettings/healthProbeSettings1"),
					},
					LoadBalancingSettings: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/loadBalancingSettings/loadBalancingSettings1"),
					},
					Name: pulumi.String("backendPool1"),
				},
			},
			BackendPoolsSettings: &network.BackendPoolsSettingsArgs{
				EnforceCertificateNameCheck: pulumi.String("Enabled"),
				SendRecvTimeoutSeconds:      pulumi.Int(60),
			},
			EnabledState:  pulumi.String("Enabled"),
			FrontDoorName: pulumi.String("frontDoor1"),
			FrontendEndpoints: network.FrontendEndpointArray{
				&network.FrontendEndpointArgs{
					HostName:                    pulumi.String("www.contoso.com"),
					Name:                        pulumi.String("frontendEndpoint1"),
					SessionAffinityEnabledState: pulumi.String("Enabled"),
					SessionAffinityTtlSeconds:   pulumi.Int(60),
					WebApplicationFirewallPolicyLink: &network.FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/policy1"),
					},
				},
				&network.FrontendEndpointArgs{
					HostName: pulumi.String("frontDoor1.azurefd.net"),
					Name:     pulumi.String("default"),
				},
			},
			HealthProbeSettings: network.HealthProbeSettingsModelArray{
				&network.HealthProbeSettingsModelArgs{
					EnabledState:      pulumi.String("Enabled"),
					HealthProbeMethod: pulumi.String("HEAD"),
					IntervalInSeconds: pulumi.Int(120),
					Name:              pulumi.String("healthProbeSettings1"),
					Path:              pulumi.String("/"),
					Protocol:          pulumi.String("Http"),
				},
			},
			LoadBalancingSettings: network.LoadBalancingSettingsModelArray{
				&network.LoadBalancingSettingsModelArgs{
					Name:                      pulumi.String("loadBalancingSettings1"),
					SampleSize:                pulumi.Int(4),
					SuccessfulSamplesRequired: pulumi.Int(2),
				},
			},
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("rg1"),
			RoutingRules: network.RoutingRuleArray{
				&network.RoutingRuleArgs{
					AcceptedProtocols: pulumi.StringArray{
						pulumi.String("Http"),
					},
					EnabledState: pulumi.String("Enabled"),
					FrontendEndpoints: network.SubResourceArray{
						&network.SubResourceArgs{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/frontendEndpoints/frontendEndpoint1"),
						},
						&network.SubResourceArgs{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/frontendEndpoints/default"),
						},
					},
					Name: pulumi.String("routingRule1"),
					PatternsToMatch: pulumi.StringArray{
						pulumi.String("/*"),
					},
					RouteConfiguration: network.ForwardingConfiguration{
						BackendPool: network.SubResource{
							Id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/backendPools/backendPool1",
						},
						OdataType: "#Microsoft.Azure.FrontDoor.Models.FrontdoorForwardingConfiguration",
					},
					RulesEngine: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/rulesEngines/rulesEngine1"),
					},
					WebApplicationFirewallPolicyLink: &network.RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/policy1"),
					},
				},
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
