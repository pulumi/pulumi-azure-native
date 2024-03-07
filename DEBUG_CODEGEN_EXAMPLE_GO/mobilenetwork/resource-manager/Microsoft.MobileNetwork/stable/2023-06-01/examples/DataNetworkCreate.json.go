package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilenetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilenetwork.NewDataNetwork(ctx, "dataNetwork", &mobilenetwork.DataNetworkArgs{
			DataNetworkName:   pulumi.String("testDataNetwork"),
			Description:       pulumi.String("myFavouriteDataNetwork"),
			Location:          pulumi.String("eastus"),
			MobileNetworkName: pulumi.String("testMobileNetwork"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
