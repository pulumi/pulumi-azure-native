package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewGuestDiagnosticsSettingsAssociation(ctx, "guestDiagnosticsSettingsAssociation", &insights.GuestDiagnosticsSettingsAssociationArgs{
			AssociationName:             pulumi.String("healthSystemMachineConfigAssociation"),
			GuestDiagnosticSettingsName: pulumi.String("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/Default-Association-ResourceGroup/providers/providers/microsoft.insights/guestDiagnosticSettings/vmSettingForSecurity"),
			Location:                    pulumi.String("Global"),
			ResourceUri:                 pulumi.String("subscriptions/8498f01b-8064-4e37-856e-318f3c6c685f/resourceGroups/healthsystem/providers/Microsoft.Compute/virtualMachines/eastussojai"),
			Tags:                        nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
