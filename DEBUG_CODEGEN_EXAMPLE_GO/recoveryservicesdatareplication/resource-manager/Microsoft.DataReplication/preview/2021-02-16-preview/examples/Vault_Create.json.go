package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datareplication/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datareplication.NewVault(ctx, "vault", &datareplication.VaultArgs{
			Location: pulumi.String("eck"),
			Properties: &datareplication.VaultModelPropertiesArgs{
				VaultType: pulumi.String("DisasterRecovery"),
			},
			ResourceGroupName: pulumi.String("rgrecoveryservicesdatareplication"),
			Tags: pulumi.StringMap{
				"key5359": pulumi.String("ljfilxolxzuxrauopwtyxghrp"),
			},
			VaultName: pulumi.String("4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
