package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/visualstudio/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := visualstudio.NewAccount(ctx, "account", &visualstudio.AccountArgs{
			AccountName:       pulumi.String("Example"),
			Location:          pulumi.String("Central US"),
			OperationType:     pulumi.String("create"),
			Properties:        nil,
			ResourceGroupName: pulumi.String("VS-Example-Group"),
			ResourceName:      pulumi.String("Example"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
