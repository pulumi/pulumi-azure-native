package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/aad/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aad.NewDomainService(ctx, "domainService", &aad.DomainServiceArgs{
			DomainName: pulumi.String("TestDomainService.com"),
			DomainSecuritySettings: &aad.DomainSecuritySettingsArgs{
				NtlmV1:            pulumi.String("Enabled"),
				SyncNtlmPasswords: pulumi.String("Enabled"),
				TlsV1:             pulumi.String("Disabled"),
			},
			DomainServiceName: pulumi.String("TestDomainService.com"),
			FilteredSync:      pulumi.String("Enabled"),
			LdapsSettings: &aad.LdapsSettingsArgs{
				ExternalAccess:         pulumi.String("Enabled"),
				Ldaps:                  pulumi.String("Enabled"),
				PfxCertificate:         pulumi.String("MIIDPDCCAiSgAwIBAgIQQUI9P6tq2p9OFIJa7DLNvTANBgkqhkiG9w0BAQsFADAgMR4w..."),
				PfxCertificatePassword: pulumi.String("<pfxCertificatePassword>"),
			},
			NotificationSettings: &aad.NotificationSettingsArgs{
				AdditionalRecipients: pulumi.StringArray{
					pulumi.String("jicha@microsoft.com"),
					pulumi.String("caalmont@microsoft.com"),
				},
				NotifyDcAdmins:     pulumi.String("Enabled"),
				NotifyGlobalAdmins: pulumi.String("Enabled"),
			},
			ReplicaSets: []aad.ReplicaSetArgs{
				{
					Location: pulumi.String("West US"),
					SubnetId: pulumi.String("/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/TestNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/TestVnetWUS/subnets/TestSubnetWUS"),
				},
			},
			ResourceGroupName: pulumi.String("TestResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
