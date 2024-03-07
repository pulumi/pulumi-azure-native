package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/labservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := labservices.NewLabPlan(ctx, "labPlan", &labservices.LabPlanArgs{
			DefaultAutoShutdownProfile: &labservices.AutoShutdownProfileArgs{
				DisconnectDelay:          pulumi.String("PT5M"),
				IdleDelay:                pulumi.String("PT5M"),
				NoConnectDelay:           pulumi.String("PT5M"),
				ShutdownOnDisconnect:     labservices.EnableStateEnabled,
				ShutdownOnIdle:           labservices.ShutdownOnIdleModeUserAbsence,
				ShutdownWhenNotConnected: labservices.EnableStateEnabled,
			},
			DefaultConnectionProfile: &labservices.ConnectionProfileArgs{
				ClientRdpAccess: labservices.ConnectionTypePublic,
				ClientSshAccess: labservices.ConnectionTypePublic,
				WebRdpAccess:    labservices.ConnectionTypeNone,
				WebSshAccess:    labservices.ConnectionTypeNone,
			},
			DefaultNetworkProfile: &labservices.LabPlanNetworkProfileArgs{
				SubnetId: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
			},
			LabPlanName:       pulumi.String("testlabplan"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("testrg123"),
			SharedGalleryId:   pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Compute/galleries/testsig"),
			SupportInfo: &labservices.SupportInfoArgs{
				Email:        pulumi.String("help@contoso.com"),
				Instructions: pulumi.String("Contact support for help."),
				Phone:        pulumi.String("+1-202-555-0123"),
				Url:          pulumi.String("help.contoso.com"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
