package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewVirtualNetwork(ctx, "virtualNetwork", &devtestlab.VirtualNetworkArgs{
			LabName:           pulumi.String("{labName}"),
			Location:          pulumi.String("{location}"),
			Name:              pulumi.String("{virtualNetworkName}"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"tagName1": pulumi.String("tagValue1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
