package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewVendor(ctx, "vendor", &hybridnetwork.VendorArgs{
			VendorName: pulumi.String("TestVendor"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
