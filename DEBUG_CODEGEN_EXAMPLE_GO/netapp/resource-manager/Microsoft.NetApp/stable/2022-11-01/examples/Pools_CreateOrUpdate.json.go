package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/netapp/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := netapp.NewPool(ctx, "pool", &netapp.PoolArgs{
			AccountName:       pulumi.String("account1"),
			Location:          pulumi.String("eastus"),
			PoolName:          pulumi.String("pool1"),
			QosType:           pulumi.String("Auto"),
			ResourceGroupName: pulumi.String("myRG"),
			ServiceLevel:      pulumi.String("Premium"),
			Size:              pulumi.Float64(4398046511104),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
