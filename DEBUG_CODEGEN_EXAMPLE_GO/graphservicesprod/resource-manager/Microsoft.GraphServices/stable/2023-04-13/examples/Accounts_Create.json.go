package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/graphservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := graphservices.NewAccount(ctx, "account", &graphservices.AccountArgs{
			Properties: &graphservices.AccountResourcePropertiesArgs{
				AppId: pulumi.String("11111111-aaaa-1111-bbbb-111111111111"),
			},
			ResourceGroupName: pulumi.String("testResourceGroupGRAM"),
			ResourceName:      pulumi.String("11111111-aaaa-1111-bbbb-1111111111111"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
