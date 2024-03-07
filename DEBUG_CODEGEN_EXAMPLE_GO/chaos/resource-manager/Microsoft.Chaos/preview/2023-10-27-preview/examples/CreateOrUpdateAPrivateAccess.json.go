package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/chaos/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := chaos.NewPrivateAccess(ctx, "privateAccess", &chaos.PrivateAccessArgs{
			Location:          pulumi.String("centraluseuap"),
			PrivateAccessName: pulumi.String("myPrivateAccess"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
