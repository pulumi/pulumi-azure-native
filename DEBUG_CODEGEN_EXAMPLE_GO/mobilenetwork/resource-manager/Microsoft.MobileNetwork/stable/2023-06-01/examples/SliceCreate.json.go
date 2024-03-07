package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilenetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilenetwork.NewSlice(ctx, "slice", &mobilenetwork.SliceArgs{
			Description:       pulumi.String("myFavouriteSlice"),
			Location:          pulumi.String("eastus"),
			MobileNetworkName: pulumi.String("testMobileNetwork"),
			ResourceGroupName: pulumi.String("rg1"),
			SliceName:         pulumi.String("testSlice"),
			Snssai: &mobilenetwork.SnssaiArgs{
				Sd:  pulumi.String("1abcde"),
				Sst: pulumi.Int(1),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
