package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/desktopvirtualization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := desktopvirtualization.NewScalingPlanPersonalSchedule(ctx, "scalingPlanPersonalSchedule", &desktopvirtualization.ScalingPlanPersonalScheduleArgs{
			DaysOfWeek: pulumi.StringArray{
				pulumi.String("Monday"),
				pulumi.String("Tuesday"),
				pulumi.String("Wednesday"),
				pulumi.String("Thursday"),
				pulumi.String("Friday"),
			},
			OffPeakActionOnDisconnect:        pulumi.String("None"),
			OffPeakActionOnLogoff:            pulumi.String("Deallocate"),
			OffPeakMinutesToWaitOnDisconnect: pulumi.Int(10),
			OffPeakMinutesToWaitOnLogoff:     pulumi.Int(10),
			OffPeakStartTime: &desktopvirtualization.TimeArgs{
				Hour:   pulumi.Int(20),
				Minute: pulumi.Int(0),
			},
			OffPeakStartVMOnConnect:       pulumi.String("Enable"),
			PeakActionOnDisconnect:        pulumi.String("None"),
			PeakActionOnLogoff:            pulumi.String("Deallocate"),
			PeakMinutesToWaitOnDisconnect: pulumi.Int(10),
			PeakMinutesToWaitOnLogoff:     pulumi.Int(10),
			PeakStartTime: &desktopvirtualization.TimeArgs{
				Hour:   pulumi.Int(8),
				Minute: pulumi.Int(0),
			},
			PeakStartVMOnConnect:              pulumi.String("Enable"),
			RampDownActionOnDisconnect:        pulumi.String("None"),
			RampDownActionOnLogoff:            pulumi.String("Deallocate"),
			RampDownMinutesToWaitOnDisconnect: pulumi.Int(10),
			RampDownMinutesToWaitOnLogoff:     pulumi.Int(10),
			RampDownStartTime: &desktopvirtualization.TimeArgs{
				Hour:   pulumi.Int(18),
				Minute: pulumi.Int(0),
			},
			RampDownStartVMOnConnect:        pulumi.String("Enable"),
			RampUpActionOnDisconnect:        pulumi.String("None"),
			RampUpActionOnLogoff:            pulumi.String("None"),
			RampUpAutoStartHosts:            pulumi.String("All"),
			RampUpMinutesToWaitOnDisconnect: pulumi.Int(10),
			RampUpMinutesToWaitOnLogoff:     pulumi.Int(10),
			RampUpStartTime: &desktopvirtualization.TimeArgs{
				Hour:   pulumi.Int(6),
				Minute: pulumi.Int(0),
			},
			RampUpStartVMOnConnect:  pulumi.String("Enable"),
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
