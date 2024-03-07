package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/autonomousdevelopmentplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := autonomousdevelopmentplatform.NewAccount(ctx, "account", &autonomousdevelopmentplatform.AccountArgs{
			AccountName:       pulumi.String("sampleacct"),
			Location:          pulumi.String("Global"),
			ResourceGroupName: pulumi.String("adpClient"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
