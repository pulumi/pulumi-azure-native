package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewAFDOriginGroup(ctx, "afdOriginGroup", &cdn.AFDOriginGroupArgs{
			HealthProbeSettings: &cdn.HealthProbeParametersArgs{
				ProbeIntervalInSeconds: pulumi.Int(10),
				ProbePath:              pulumi.String("/path2"),
				ProbeProtocol:          cdn.ProbeProtocolNotSet,
				ProbeRequestType:       cdn.HealthProbeRequestTypeNotSet,
			},
			LoadBalancingSettings: &cdn.LoadBalancingSettingsParametersArgs{
				AdditionalLatencyInMilliseconds: pulumi.Int(1000),
				SampleSize:                      pulumi.Int(3),
				SuccessfulSamplesRequired:       pulumi.Int(3),
			},
			OriginGroupName:   pulumi.String("origingroup1"),
			ProfileName:       pulumi.String("profile1"),
			ResourceGroupName: pulumi.String("RG"),
			TrafficRestorationTimeToHealedOrNewEndpointsInMinutes: pulumi.Int(5),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
