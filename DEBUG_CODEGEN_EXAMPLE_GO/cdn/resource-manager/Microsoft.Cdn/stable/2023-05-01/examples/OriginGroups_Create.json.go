package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewOriginGroup(ctx, "originGroup", &cdn.OriginGroupArgs{
			EndpointName: pulumi.String("endpoint1"),
			HealthProbeSettings: &cdn.HealthProbeParametersArgs{
				ProbeIntervalInSeconds: pulumi.Int(120),
				ProbePath:              pulumi.String("/health.aspx"),
				ProbeProtocol:          cdn.ProbeProtocolHttp,
				ProbeRequestType:       cdn.HealthProbeRequestTypeGET,
			},
			OriginGroupName: pulumi.String("origingroup1"),
			Origins: []cdn.ResourceReferenceArgs{
				{
					Id: pulumi.String("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin1"),
				},
			},
			ProfileName:       pulumi.String("profile1"),
			ResourceGroupName: pulumi.String("RG"),
			ResponseBasedOriginErrorDetectionSettings: &cdn.ResponseBasedOriginErrorDetectionParametersArgs{
				ResponseBasedDetectedErrorTypes:          cdn.ResponseBasedDetectedErrorTypesTcpErrorsOnly,
				ResponseBasedFailoverThresholdPercentage: pulumi.Int(10),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
