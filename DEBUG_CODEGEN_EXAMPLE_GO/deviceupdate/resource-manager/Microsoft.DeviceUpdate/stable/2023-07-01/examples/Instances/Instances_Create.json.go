package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/deviceupdate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := deviceupdate.NewInstance(ctx, "instance", &deviceupdate.InstanceArgs{
			AccountName: pulumi.String("contoso"),
			DiagnosticStorageProperties: &deviceupdate.DiagnosticStoragePropertiesArgs{
				AuthenticationType: pulumi.String("KeyBased"),
				ConnectionString:   pulumi.String("string"),
				ResourceId:         pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/adu-resource-group/providers/Microsoft.Storage/storageAccounts/testAccount"),
			},
			EnableDiagnostics: pulumi.Bool(false),
			InstanceName:      pulumi.String("blue"),
			IotHubs: deviceupdate.IotHubSettingsArray{
				&deviceupdate.IotHubSettingsArgs{
					ResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Devices/IotHubs/blue-contoso-hub"),
				},
			},
			Location:          pulumi.String("westus2"),
			ResourceGroupName: pulumi.String("test-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
