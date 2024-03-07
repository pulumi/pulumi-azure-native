package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewFleetUpdateStrategy(ctx, "fleetUpdateStrategy", &containerservice.FleetUpdateStrategyArgs{
			FleetName:         pulumi.String("fleet1"),
			ResourceGroupName: pulumi.String("rg1"),
			Strategy: &containerservice.UpdateRunStrategyArgs{
				Stages: containerservice.UpdateStageArray{
					&containerservice.UpdateStageArgs{
						AfterStageWaitInSeconds: pulumi.Int(3600),
						Groups: containerservice.UpdateGroupArray{
							&containerservice.UpdateGroupArgs{
								Name: pulumi.String("group-a"),
							},
						},
						Name: pulumi.String("stage1"),
					},
				},
			},
			UpdateStrategyName: pulumi.String("strartegy1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
