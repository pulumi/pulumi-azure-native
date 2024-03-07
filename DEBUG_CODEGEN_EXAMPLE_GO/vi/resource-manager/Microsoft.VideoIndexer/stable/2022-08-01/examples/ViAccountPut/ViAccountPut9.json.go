package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/videoindexer/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := videoindexer.NewAccount(ctx, "account", &videoindexer.AccountArgs{
			AccountName:       pulumi.String("contosto-videoanalyzer"),
			ResourceGroupName: pulumi.String("contosto-videoanalyzer-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
