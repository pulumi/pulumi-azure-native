package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewTable(ctx, "table", &storage.TableArgs{
			AccountName:       pulumi.String("sto328"),
			ResourceGroupName: pulumi.String("res3376"),
			TableName:         pulumi.String("table6185"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
