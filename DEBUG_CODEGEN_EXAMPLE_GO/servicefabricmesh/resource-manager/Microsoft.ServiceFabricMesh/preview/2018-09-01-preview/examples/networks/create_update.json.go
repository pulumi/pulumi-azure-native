package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabricmesh/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabricmesh.NewNetwork(ctx, "network", &servicefabricmesh.NetworkArgs{
			Location:            pulumi.String("EastUS"),
			NetworkResourceName: pulumi.String("sampleNetwork"),
			Properties:          nil,
			ResourceGroupName:   pulumi.String("sbz_demo"),
			Tags:                nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
