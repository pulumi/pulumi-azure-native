package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilenetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilenetwork.NewMobileNetwork(ctx, "mobileNetwork", &mobilenetwork.MobileNetworkArgs{
			Location:          pulumi.String("eastus"),
			MobileNetworkName: pulumi.String("testMobileNetwork"),
			PublicLandMobileNetworkIdentifier: &mobilenetwork.PlmnIdArgs{
				Mcc: pulumi.String("001"),
				Mnc: pulumi.String("01"),
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
