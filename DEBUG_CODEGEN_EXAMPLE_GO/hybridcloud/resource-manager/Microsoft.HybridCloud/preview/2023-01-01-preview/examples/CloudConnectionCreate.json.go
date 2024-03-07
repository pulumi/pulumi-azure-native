package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcloud.NewCloudConnection(ctx, "cloudConnection", &hybridcloud.CloudConnectionArgs{
			CloudConnectionName: pulumi.String("cloudconnection1"),
			CloudConnector: &hybridcloud.ResourceReferenceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/demo-rg/providers/Microsoft.HybridCloud/cloudConnectors/123456789012"),
			},
			Location:          pulumi.String("West US"),
			RemoteResourceId:  pulumi.String("arn:aws:ec2:us-east-1:123456789012:VPNGateway/vgw-043da592550819c8a"),
			ResourceGroupName: pulumi.String("demo-rg"),
			SharedKey:         pulumi.String("password123"),
			VirtualHub: &hybridcloud.ResourceReferenceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/demo-rg/providers/Microsoft.Network/VirtualHubs/testHub"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
