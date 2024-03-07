package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewCloudService(ctx, "cloudService", &compute.CloudServiceArgs{
			CloudServiceName: pulumi.String("{cs-name}"),
			Location:         pulumi.String("westus"),
			Properties: &compute.CloudServicePropertiesArgs{
				Configuration: pulumi.String("{ServiceConfiguration}"),
				ExtensionProfile: &compute.CloudServiceExtensionProfileArgs{
					Extensions: compute.ExtensionArray{
						&compute.ExtensionArgs{
							Name: pulumi.String("RDPExtension"),
							Properties: &compute.CloudServiceExtensionPropertiesArgs{
								AutoUpgradeMinorVersion: pulumi.Bool(false),
								ProtectedSettings:       pulumi.Any("<PrivateConfig><Password>{password}</Password></PrivateConfig>"),
								Publisher:               pulumi.String("Microsoft.Windows.Azure.Extensions"),
								Settings:                pulumi.Any("<PublicConfig><UserName>UserAzure</UserName><Expiration>10/22/2021 15:05:45</Expiration></PublicConfig>"),
								Type:                    pulumi.String("RDP"),
								TypeHandlerVersion:      pulumi.String("1.2"),
							},
						},
					},
				},
				NetworkProfile: &compute.CloudServiceNetworkProfileArgs{
					LoadBalancerConfigurations: compute.LoadBalancerConfigurationArray{
						&compute.LoadBalancerConfigurationArgs{
							Name: pulumi.String("contosolb"),
							Properties: &compute.LoadBalancerConfigurationPropertiesArgs{
								FrontendIpConfigurations: compute.LoadBalancerFrontendIpConfigurationArray{
									&compute.LoadBalancerFrontendIpConfigurationArgs{
										Name: pulumi.String("contosofe"),
										Properties: &compute.LoadBalancerFrontendIpConfigurationPropertiesArgs{
											PublicIPAddress: &compute.SubResourceArgs{
												Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Network/publicIPAddresses/contosopublicip"),
											},
										},
									},
								},
							},
						},
					},
				},
				PackageUrl: pulumi.String("{PackageUrl}"),
				RoleProfile: &compute.CloudServiceRoleProfileArgs{
					Roles: compute.CloudServiceRoleProfilePropertiesArray{
						&compute.CloudServiceRoleProfilePropertiesArgs{
							Name: pulumi.String("ContosoFrontend"),
							Sku: &compute.CloudServiceRoleSkuArgs{
								Capacity: pulumi.Float64(1),
								Name:     pulumi.String("Standard_D1_v2"),
								Tier:     pulumi.String("Standard"),
							},
						},
					},
				},
				UpgradeMode: pulumi.String("Auto"),
			},
			ResourceGroupName: pulumi.String("ConstosoRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
