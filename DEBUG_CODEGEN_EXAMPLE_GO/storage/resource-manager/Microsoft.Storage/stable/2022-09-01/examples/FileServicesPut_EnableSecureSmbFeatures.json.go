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
					AuthenticationMethods:    pulumi.String("NTLMv2;Kerberos"),
					ChannelEncryption:        pulumi.String("AES-128-CCM;AES-128-GCM;AES-256-GCM"),
					KerberosTicketEncryption: pulumi.String("RC4-HMAC;AES-256"),
					Versions:                 pulumi.String("SMB2.1;SMB3.0;SMB3.1.1"),
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
