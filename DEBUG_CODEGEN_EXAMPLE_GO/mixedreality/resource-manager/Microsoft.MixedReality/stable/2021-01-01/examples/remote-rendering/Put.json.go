package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mixedreality/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mixedreality.NewRemoteRenderingAccount(ctx, "remoteRenderingAccount", &mixedreality.RemoteRenderingAccountArgs{
			AccountName: pulumi.String("MyAccount"),
			Identity: &mixedreality.IdentityArgs{
				Type: mixedreality.ResourceIdentityTypeSystemAssigned,
			},
			Location:          pulumi.String("eastus2euap"),
			ResourceGroupName: pulumi.String("MyResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
