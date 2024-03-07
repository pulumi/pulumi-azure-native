package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcloud.NewCloudConnector(ctx, "cloudConnector", &hybridcloud.CloudConnectorArgs{
			AccountId:          pulumi.String("123456789012"),
			CloudConnectorName: pulumi.String("123456789012"),
			CloudType:          pulumi.String("AWS"),
			Location:           pulumi.String("West US"),
			ResourceGroupName:  pulumi.String("demo-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
