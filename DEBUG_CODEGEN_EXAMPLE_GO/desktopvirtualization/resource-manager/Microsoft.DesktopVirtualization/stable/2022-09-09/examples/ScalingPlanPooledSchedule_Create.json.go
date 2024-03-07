package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/desktopvirtualization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := desktopvirtualization.NewScalingPlanPooledSchedule(ctx, "scalingPlanPooledSchedule", &desktopvirtualization.ScalingPlanPooledScheduleArgs{
			DaysOfWeek: pulumi.StringArray{
				pulumi.String("Monday"),
				pulumi.String("Tuesday"),
				pulumi.String("Wednesday"),
				pulumi.String("Thursday"),
				pulumi.String("Friday"),
			},
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
			ResourceGroupName:       pulumi.String("resourceGroup1"),
			ScalingPlanName:         pulumi.String("scalingPlan1"),
			ScalingPlanScheduleName: pulumi.String("scalingPlanScheduleWeekdays1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
