package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApiManagementService(ctx, "apiManagementService", &apimanagement.ApiManagementServiceArgs{
			Location:          pulumi.String("East US 2 EUAP"),
			PublicIpAddressId: pulumi.String("/subscriptions/subid/resourceGroups/rgName/providers/Microsoft.Network/publicIPAddresses/apimazvnet"),
			PublisherEmail:    pulumi.String("apim@autorestsdk.com"),
			PublisherName:     pulumi.String("autorestsdk"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Sku: &apimanagement.ApiManagementServiceSkuPropertiesArgs{
				Capacity: pulumi.Int(2),
				Name:     pulumi.String("Premium"),
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
				"tag3": pulumi.String("value3"),
			},
			VirtualNetworkConfiguration: &apimanagement.VirtualNetworkConfigurationArgs{
				SubnetResourceId: pulumi.String("/subscriptions/subid/resourceGroups/rgName/providers/Microsoft.Network/virtualNetworks/apimcus/subnets/tenant"),
			},
			VirtualNetworkType: pulumi.String("External"),
			Zones: pulumi.StringArray{
				pulumi.String("1"),
				pulumi.String("2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
