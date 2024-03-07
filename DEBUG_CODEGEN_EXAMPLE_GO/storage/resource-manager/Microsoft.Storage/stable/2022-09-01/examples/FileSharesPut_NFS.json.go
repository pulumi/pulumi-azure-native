package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewFileShare(ctx, "fileShare", &storage.FileShareArgs{
			AccountName:       pulumi.String("sto666"),
			EnabledProtocols:  pulumi.String("NFS"),
			ResourceGroupName: pulumi.String("res346"),
			ShareName:         pulumi.String("share1235"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
