package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datashare/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datashare.NewShare(ctx, "share", &datashare.ShareArgs{
			AccountName:       pulumi.String("Account1"),
			Description:       pulumi.String("share description"),
			ResourceGroupName: pulumi.String("SampleResourceGroup"),
			ShareKind:         pulumi.String("CopyBased"),
			ShareName:         pulumi.String("Share1"),
			Terms:             pulumi.String("Confidential"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
