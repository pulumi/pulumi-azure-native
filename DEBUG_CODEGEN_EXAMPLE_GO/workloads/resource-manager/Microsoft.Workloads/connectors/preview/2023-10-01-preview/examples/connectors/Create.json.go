package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewConnector(ctx, "connector", &workloads.ConnectorArgs{
			ConnectorName:     pulumi.String("C1"),
			Location:          pulumi.String("westcentralus"),
			ResourceGroupName: pulumi.String("test-rg"),
			SourceResourceId:  pulumi.String("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
