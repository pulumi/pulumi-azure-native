package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/labservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := labservices.NewLab(ctx, "lab", &labservices.LabArgs{
			AutoShutdownProfile: &labservices.AutoShutdownProfileArgs{
				DisconnectDelay:          pulumi.String("PT5M"),
				IdleDelay:                pulumi.String("PT5M"),
				NoConnectDelay:           pulumi.String("PT5M"),
				ShutdownOnDisconnect:     labservices.EnableStateEnabled,
				ShutdownOnIdle:           labservices.ShutdownOnIdleModeUserAbsence,
				ShutdownWhenNotConnected: labservices.EnableStateEnabled,
			},
			ConnectionProfile: &labservices.ConnectionProfileArgs{
				ClientRdpAccess: labservices.ConnectionTypePublic,
				ClientSshAccess: labservices.ConnectionTypePublic,
				WebRdpAccess:    labservices.ConnectionTypeNone,
				WebSshAccess:    labservices.ConnectionTypeNone,
			},
			Description: pulumi.String("This is a test lab."),
			LabName:     pulumi.String("testlab"),
			LabPlanId:   pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.LabServices/labPlans/testlabplan"),
			Location:    pulumi.String("westus"),
			NetworkProfile: &labservices.LabNetworkProfileArgs{
				SubnetId: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
			},
			ResourceGroupName: pulumi.String("testrg123"),
			SecurityProfile: &labservices.SecurityProfileArgs{
				OpenAccess: labservices.EnableStateDisabled,
			},
			Title: pulumi.String("Test Lab"),
			VirtualMachineProfile: labservices.VirtualMachineProfileResponse{
				AdditionalCapabilities: &labservices.VirtualMachineAdditionalCapabilitiesArgs{
					InstallGpuDrivers: labservices.EnableStateDisabled,
				},
				AdminUser: &labservices.CredentialsArgs{
					Username: pulumi.String("test-user"),
				},
				CreateOption: labservices.CreateOptionTemplateVM,
				ImageReference: &labservices.ImageReferenceArgs{
					Offer:     pulumi.String("WindowsServer"),
					Publisher: pulumi.String("Microsoft"),
					Sku:       pulumi.String("2019-Datacenter"),
					Version:   pulumi.String("2019.0.20190410"),
				},
				Sku: &labservices.SkuArgs{
					Name: pulumi.String("Medium"),
				},
				UsageQuota:        pulumi.String("PT10H"),
				UseSharedPassword: labservices.EnableStateDisabled,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
