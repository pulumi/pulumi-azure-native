package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/elasticsan/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticsan.NewVolumeGroup(ctx, "volumeGroup", &elasticsan.VolumeGroupArgs{
			ElasticSanName:    pulumi.String("ti7q-k952-1qB3J_5"),
			ResourceGroupName: pulumi.String("rgelasticsan"),
			VolumeGroupName:   pulumi.String("u_5I_1j4t3"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
