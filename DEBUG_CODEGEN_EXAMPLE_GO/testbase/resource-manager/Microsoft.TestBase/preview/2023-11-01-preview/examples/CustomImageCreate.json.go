package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/testbase/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := testbase.NewCustomImage(ctx, "customImage", &testbase.CustomImageArgs{
			CustomImageName:     pulumi.String("image-2cfb2edc-13bc-4d54-8d6e-38b2a233b003"),
			DefinitionName:      pulumi.String("contoso-image-def"),
			ResourceGroupName:   pulumi.String("contoso-rg1"),
			Source:              pulumi.String("VHD"),
			TestBaseAccountName: pulumi.String("contoso-testBaseAccount1"),
			VersionName:         pulumi.String("1.0.0"),
			VhdId:               pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/VHDs/vhd-00ac3ccd-1503-4ee5-aa26-26569cfafe88"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
