package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/desktopvirtualization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := desktopvirtualization.NewHostPool(ctx, "hostPool", &desktopvirtualization.HostPoolArgs{
			AgentUpdate: &desktopvirtualization.AgentUpdatePropertiesArgs{
				MaintenanceWindowTimeZone: pulumi.String("Alaskan Standard Time"),
				MaintenanceWindows: desktopvirtualization.MaintenanceWindowPropertiesArray{
					&desktopvirtualization.MaintenanceWindowPropertiesArgs{
						DayOfWeek: desktopvirtualization.DayOfWeekFriday,
						Hour:      pulumi.Int(7),
					},
					&desktopvirtualization.MaintenanceWindowPropertiesArgs{
						DayOfWeek: desktopvirtualization.DayOfWeekSaturday,
						Hour:      pulumi.Int(8),
					},
				},
				Type:                    pulumi.String("Scheduled"),
				UseSessionHostLocalTime: pulumi.Bool(false),
			},
			Description:                   pulumi.String("des1"),
			FriendlyName:                  pulumi.String("friendly"),
			HostPoolName:                  pulumi.String("hostPool1"),
			HostPoolType:                  pulumi.String("Pooled"),
			LoadBalancerType:              pulumi.String("BreadthFirst"),
			Location:                      pulumi.String("centralus"),
			MaxSessionLimit:               pulumi.Int(999999),
			PersonalDesktopAssignmentType: pulumi.String("Automatic"),
			PreferredAppGroupType:         pulumi.String("Desktop"),
			RegistrationInfo: &desktopvirtualization.RegistrationInfoArgs{
				ExpirationTime:             pulumi.String("2020-10-01T14:01:54.9571247Z"),
				RegistrationTokenOperation: pulumi.String("Update"),
			},
			ResourceGroupName:           pulumi.String("resourceGroup1"),
			SsoClientId:                 pulumi.String("client"),
			SsoClientSecretKeyVaultPath: pulumi.String("https://keyvault/secret"),
			SsoSecretType:               pulumi.String("SharedKey"),
			SsoadfsAuthority:            pulumi.String("https://adfs"),
			StartVMOnConnect:            pulumi.Bool(false),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
			VmTemplate: pulumi.String("{json:json}"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
