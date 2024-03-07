package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/batch/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := batch.NewPool(ctx, "pool", &batch.PoolArgs{
			AccountName: pulumi.String("sampleacct"),
			DeploymentConfiguration: batch.DeploymentConfigurationResponse{
				CloudServiceConfiguration: &batch.CloudServiceConfigurationArgs{
					OsFamily: pulumi.String("5"),
				},
			},
			PoolName:          pulumi.String("testpool"),
			ResourceGroupName: pulumi.String("default-azurebatch-japaneast"),
			ScaleSettings: batch.ScaleSettingsResponse{
				FixedScale: &batch.FixedScaleSettingsArgs{
					TargetDedicatedNodes: pulumi.Int(3),
				},
			},
			VmSize: pulumi.String("STANDARD_D4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
