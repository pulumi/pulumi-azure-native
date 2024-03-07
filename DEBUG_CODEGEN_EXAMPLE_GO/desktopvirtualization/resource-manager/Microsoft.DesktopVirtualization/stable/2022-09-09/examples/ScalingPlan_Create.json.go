package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/desktopvirtualization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := desktopvirtualization.NewScalingPlan(ctx, "scalingPlan", &desktopvirtualization.ScalingPlanArgs{
			Description:  pulumi.String("Description of Scaling Plan"),
			ExclusionTag: pulumi.String("value"),
			FriendlyName: pulumi.String("Scaling Plan 1"),
			HostPoolReferences: desktopvirtualization.ScalingHostPoolReferenceArray{
				&desktopvirtualization.ScalingHostPoolReferenceArgs{
					HostPoolArmPath:    pulumi.String("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1"),
					ScalingPlanEnabled: pulumi.Bool(true),
				},
			},
			HostPoolType:      pulumi.String("Pooled"),
			Location:          pulumi.String("centralus"),
			ResourceGroupName: pulumi.String("resourceGroup1"),
			ScalingPlanName:   pulumi.String("scalingPlan1"),
			Schedules: desktopvirtualization.ScalingScheduleArray{
				&desktopvirtualization.ScalingScheduleArgs{
					DaysOfWeek: pulumi.StringArray{
						pulumi.String("Monday"),
						pulumi.String("Tuesday"),
						pulumi.String("Wednesday"),
						pulumi.String("Thursday"),
						pulumi.String("Friday"),
					},
					Name:                          pulumi.String("schedule1"),
					OffPeakLoadBalancingAlgorithm: pulumi.String("DepthFirst"),
					OffPeakStartTime: &desktopvirtualization.TimeArgs{
						Hour:   pulumi.Int(20),
						Minute: pulumi.Int(0),
					},
					PeakLoadBalancingAlgorithm: pulumi.String("BreadthFirst"),
					PeakStartTime: &desktopvirtualization.TimeArgs{
						Hour:   pulumi.Int(8),
						Minute: pulumi.Int(0),
					},
					RampDownCapacityThresholdPct:   pulumi.Int(50),
					RampDownForceLogoffUsers:       pulumi.Bool(true),
					RampDownLoadBalancingAlgorithm: pulumi.String("DepthFirst"),
					RampDownMinimumHostsPct:        pulumi.Int(20),
					RampDownNotificationMessage:    pulumi.String("message"),
					RampDownStartTime: &desktopvirtualization.TimeArgs{
						Hour:   pulumi.Int(18),
						Minute: pulumi.Int(0),
					},
					RampDownWaitTimeMinutes:      pulumi.Int(30),
					RampUpCapacityThresholdPct:   pulumi.Int(80),
					RampUpLoadBalancingAlgorithm: pulumi.String("DepthFirst"),
					RampUpMinimumHostsPct:        pulumi.Int(20),
					RampUpStartTime: &desktopvirtualization.TimeArgs{
						Hour:   pulumi.Int(6),
						Minute: pulumi.Int(0),
					},
				},
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
			TimeZone: pulumi.String("Central Standard Time"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
