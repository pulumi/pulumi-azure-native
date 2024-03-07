package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/standbypool/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := standbypool.NewStandbyContainerGroupPool(ctx, "standbyContainerGroupPool", &standbypool.StandbyContainerGroupPoolArgs{
			Location:                      pulumi.String("West US"),
			ResourceGroupName:             pulumi.String("rgstandbypool"),
			StandbyContainerGroupPoolName: pulumi.String("pool"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
