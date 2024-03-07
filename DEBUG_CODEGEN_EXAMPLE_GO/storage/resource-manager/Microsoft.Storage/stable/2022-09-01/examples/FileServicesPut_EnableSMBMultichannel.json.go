package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewFileServiceProperties(ctx, "fileServiceProperties", &storage.FileServicePropertiesArgs{
			AccountName:      pulumi.String("sto8607"),
			FileServicesName: pulumi.String("default"),
			ProtocolSettings: &storage.ProtocolSettingsArgs{
				Smb: &storage.SmbSettingArgs{
					Multichannel: &storage.MultichannelArgs{
						Enabled: pulumi.Bool(true),
					},
				},
			},
			ResourceGroupName: pulumi.String("res4410"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
